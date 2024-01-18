package manager

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/types"
)

type (
	TargetSeedKiraConfig struct {
		IpAddress     string
		InterxPort    string
		SekaidRPCPort string
		SekaidP2PPort string
	}
	JoinerManager struct {
		client           *httpKiraClient
		targetConfig     *TargetSeedKiraConfig
		configController ConfigController

		log *logging.Logger
	}

	ConfigController interface {
		ReadOrCreateConfig() (cfg *config.KiraConfig, err error)
		GetConfigFilePath() (filePath, folderPath string)
		WriteConfigFile(filePath string, cfg *config.KiraConfig) error
	}

	// Structure httpKiraClient is a struct for making HTTP calls to the nodes for retrieving information about network
	httpKiraClient struct {
		client *http.Client
		log    *logging.Logger
	}

	// Structure networkInfo is a struct which represents the general information
	// about network which will be used for future connection
	networkInfo struct {
		NetworkName string
		NodeID      string
		BlockHeight string
		Seeds       []string
	}

	// Structure syncInfo is a struct which represents synchronization info from target network
	// This struct is used for generating config values for starting sekaid instance
	syncInfo struct {
		rpcServers       []string
		trustHeightBlock string
		trustHashBlock   string
	}
)

const (
	endpointStatus     = "status"
	endpointPubP2PList = "api/pub_p2p_list?peers_only=true"
)

func NewJoinerManager(config *TargetSeedKiraConfig) *JoinerManager {
	return &JoinerManager{
		client:       &httpKiraClient{client: &http.Client{}},
		targetConfig: config,
	}
}

// GenerateKiraConfig generates KiraConfig with target information for future connection
func (j *JoinerManager) GenerateKiraConfig(ctx context.Context, recover bool) (*config.KiraConfig, error) {
	networkInfo, err := j.retrieveNetworkInformation(ctx)
	if err != nil {
		return nil, fmt.Errorf("retrieving network information, error: %w", err)
	}

	configs, err := j.getConfigsBasedOnSeed(ctx, networkInfo)
	if err != nil {
		j.log.Errorf("Can't get calculated config values: %s", err)
		return nil, err
	}

	cfg, err := j.configController.ReadOrCreateConfig()
	if err != nil {
		j.log.Errorf("Reading/Creating config error: %s", err)
		return nil, err
	}

	cfg.NetworkName = networkInfo.NetworkName
	cfg.ConfigTomlValues = configs
	cfg.Recover = recover
	filePath, _ := j.configController.GetConfigFilePath()

	err = j.configController.WriteConfigFile(filePath, cfg)
	if err != nil {
		j.log.Errorf("writing config file error: %s", err)
		return nil, err
	}

	return cfg, nil
}

// GetVerifiedGenesisFile fetches and verifies the Genesis file from both the 'sekaid' and 'interx' target sources.
func (j *JoinerManager) GetVerifiedGenesisFile(ctx context.Context) ([]byte, error) {
	genesisSekaid, err := j.client.getSekaidGenesis(ctx, j.targetConfig.IpAddress, j.targetConfig.SekaidRPCPort)
	if err != nil {
		j.log.Error("Can't get 'sekaid' genesis")
		return nil, err
	}

	genesisInterx, err := j.client.getInterxGenesis(ctx, j.targetConfig.IpAddress, j.targetConfig.InterxPort)
	if err != nil {
		j.log.Error("Can't get 'interx' genesis")
		return nil, err
	}

	if err := j.checkFileContentGenesisFiles(genesisInterx, genesisSekaid); err != nil {
		j.log.Errorf("Comparing genesis files error: %s", err)
		return nil, err
	}

	if err := j.checkGenSum(ctx, genesisSekaid); err != nil {
		return nil, err
	}

	j.log.Info("Genesis file is valid")
	return genesisSekaid, nil
}

// getConfigsBasedOnSeed generates a slice of configuration values based on the provided network information
// and joins the seeds, RPC servers, and other relevant parameters into the configuration values.
func (j *JoinerManager) getConfigsBasedOnSeed(ctx context.Context, netInfo *networkInfo) ([]config.TomlValue, error) {
	configValues := make([]config.TomlValue, 0)

	configValues = append(configValues, config.TomlValue{Tag: "p2p", Name: "seeds", Value: strings.Join(netInfo.Seeds, ",")})

	listOfRPC, err := j.parseRPCfromSeedsList(netInfo.Seeds)
	if err != nil {
		return nil, fmt.Errorf("parsing RPCs from seeds list, error: %w", err)
	}

	syncInfo, err := j.getSyncInfo(ctx, listOfRPC, netInfo.BlockHeight)
	if err != nil {
		return nil, fmt.Errorf("getting sync information, error: %w", err)
	}

	if syncInfo != nil {
		configValues = append(configValues, config.TomlValue{Tag: "statesync", Name: "trust_hash", Value: syncInfo.trustHashBlock})
		configValues = append(configValues, config.TomlValue{Tag: "statesync", Name: "trust_height", Value: syncInfo.trustHeightBlock})
		configValues = append(configValues, config.TomlValue{Tag: "statesync", Name: "rpc_servers", Value: strings.Join(syncInfo.rpcServers, ",")})
		configValues = append(configValues, config.TomlValue{Tag: "statesync", Name: "trust_period", Value: "168h0m0s"})
		configValues = append(configValues, config.TomlValue{Tag: "statesync", Name: "enable", Value: "true"})
		configValues = append(configValues, config.TomlValue{Tag: "statesync", Name: "temp_dir", Value: "/tmp"})
	}

	return configValues, nil
}

// retrieveNetworkInformation fetches network-related information from the Sekaid node.
// It queries the node's status and public P2P list to construct a networkInfo struct containing essential network details.
func (j *JoinerManager) retrieveNetworkInformation(ctx context.Context) (*networkInfo, error) {
	statusResponse, err := j.client.getSekaidStatus(ctx, j.targetConfig.IpAddress, j.targetConfig.SekaidRPCPort)
	if err != nil {
		return nil, fmt.Errorf("getting sekaid status, error: %w", err)
	}

	pupP2PListResponse, err := j.client.getPubP2PList(ctx, j.targetConfig.IpAddress, j.targetConfig.InterxPort)
	if err != nil {
		return nil, fmt.Errorf("getting sekaid public P2P list, error: %w", err)
	}

	listOfSeeds, err := j.parsePubP2PListResponse(ctx, pupP2PListResponse)
	if err != nil {
		return nil, fmt.Errorf("parsing sekaid public P2P list, error: %w", err)
	}
	if len(listOfSeeds) == 0 {
		j.log.Warn("List of seeds is empty, the trusted seed will be used")
		listOfSeeds = []string{fmt.Sprintf("tcp://%s@%s:%s", statusResponse.Result.NodeInfo.ID, j.targetConfig.IpAddress, j.targetConfig.SekaidP2PPort)}
	}

	return &networkInfo{
		NetworkName: statusResponse.Result.NodeInfo.Network,
		NodeID:      statusResponse.Result.NodeInfo.ID,
		BlockHeight: statusResponse.Result.SyncInfo.LatestBlockHeight,
		Seeds:       listOfSeeds,
	}, nil
}

// parsePubP2PListResponse parses the raw response data containing a list of public P2P peers and formats them as seed URLs.
// It splits the input response into lines, converts each line into a seed URL format, and returns a slice of formatted seed URLs.
func (j *JoinerManager) parsePubP2PListResponse(ctx context.Context, seedsResponse []byte) ([]string, error) {
	if len(seedsResponse) == 0 {
		j.log.Warning("The list of public seeds is not available")
		return nil, nil
	}

	linesOfPeers := strings.Split(string(seedsResponse), "\n")
	listOfSeeds := make([]string, 0)

	for _, line := range linesOfPeers {
		formattedSeed := fmt.Sprintf("tcp://%s", line)
		j.log.Debugf("Got seed: %s", formattedSeed)
		listOfSeeds = append(listOfSeeds, formattedSeed)
	}

	return listOfSeeds, nil
}

// getSyncInfo retrieves synchronization information based on a list of RPC servers and a minimum block height.
// It queries each RPC server for block information at the specified height and checks if the retrieved data is consistent.
func (j *JoinerManager) getSyncInfo(ctx context.Context, listOfRPC []string, minHeight string) (*syncInfo, error) {
	resultSyncInfo := &syncInfo{
		rpcServers:       []string{},
		trustHeightBlock: "",
		trustHashBlock:   "",
	}

	for _, rpcServer := range listOfRPC {
		responseBlock, err := j.client.getBlockInfo(ctx, rpcServer, minHeight)
		if err != nil {
			j.log.Infof("Can't get block information from RPC '%s'", rpcServer)
			continue
		}

		if responseBlock.Result.Block.Header.Height != minHeight {
			j.log.Infof("RPC (%s) height is '%s', but expected '%s'", rpcServer, responseBlock.Result.Block.Header.Height, minHeight)
			continue
		}

		if responseBlock.Result.BlockID.Hash != resultSyncInfo.trustHashBlock && resultSyncInfo.trustHashBlock != "" {
			j.log.Infof("RPC (%s) hash is '%s', but expected '%s'", rpcServer, responseBlock.Result.BlockID.Hash, resultSyncInfo.trustHashBlock)
			continue
		}

		resultSyncInfo.trustHashBlock = responseBlock.Result.BlockID.Hash
		resultSyncInfo.trustHeightBlock = minHeight

		j.log.Infof("Adding RPC (%s) to RPC connection list", rpcServer)
		resultSyncInfo.rpcServers = append(resultSyncInfo.rpcServers, rpcServer)
	}

	if len(resultSyncInfo.rpcServers) < 2 {
		j.log.Info("Sync is NOT possible (not enough RPC servers)")
		return nil, nil
	}

	j.log.Debug(resultSyncInfo)
	return resultSyncInfo, nil
}

// parseRPCfromSeedsList extracts and formats RPC server URLs from a list of seed nodes.
// It takes a slice of seed nodes and transforms them into RPC server URLs by extracting the IP address
// and combining it with the specified Sekaid RPC port.
func (j *JoinerManager) parseRPCfromSeedsList(seeds []string) ([]string, error) {
	listOfRPCs := make([]string, 0)

	for _, seed := range seeds {
		// tcp://23ca3770ae3874ac8f5a6f84a5cfaa1b39e49fc9@128.140.86.241:26656 -> 128.140.86.241:26657
		parts := strings.Split(seed, "@")
		if len(parts) != 2 {
			return nil, fmt.Errorf("%w: '%s'", ErrInvalidSeedFormat, seed)
		}

		ipAndPort := strings.Split(parts[1], ":")
		if len(ipAndPort) != 2 {
			return nil, fmt.Errorf("%w: '%s'", ErrInvalidIPPortFormat, seed)
		}

		rpc := fmt.Sprintf("%s:%s", ipAndPort[0], j.targetConfig.SekaidRPCPort)
		j.log.Infof("Adding rpc to list: %s", rpc)
		listOfRPCs = append(listOfRPCs, rpc)
	}

	return listOfRPCs, nil
}

// checkFileContentGenesisFiles checks if the content of two Genesis files is identical.
func (JoinerManager) checkFileContentGenesisFiles(genesis1, genesis2 []byte) error {
	if string(genesis1) != string(genesis2) {
		return ErrFilesContentNotIdentical
	}

	return nil
}

// checkGenSum checks the integrity of a Genesis file using its SHA256 checksum.
func (j *JoinerManager) checkGenSum(ctx context.Context, genesis []byte) error {
	genesisSum, err := j.client.getGenSum(ctx, j.targetConfig.IpAddress, j.targetConfig.InterxPort)
	if err != nil {
		return fmt.Errorf("can't get genesis check sum: %w", err)
	}

	genSumGenesisHash := sha256.Sum256(genesis)
	hashString := hex.EncodeToString(genSumGenesisHash[:])

	if genesisSum != hashString {
		j.log.Error("sha256 check sum is not the same")
		return ErrSHA256ChecksumMismatch
	}

	return nil
}

// doGetHttpQuery performs an HTTP GET request to the specified URL and returns the response body as a byte slice.
func (h *httpKiraClient) doGetHttpQuery(ctx context.Context, url string) ([]byte, error) {
	const timeoutQuery = time.Second * 60

	ctx, cancel := context.WithTimeout(ctx, timeoutQuery)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		h.log.Errorf("Failed to create request: %s", err)
		return nil, err
	}

	h.log.Infof("Querying to '%s'", url)
	resp, err := h.client.Do(req)
	if err != nil {
		h.log.Errorf("Failed to send request: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		h.log.Errorf("Failed to read response body: %s", err)
		return nil, err
	}

	h.log.Debug(string(body))

	return body, nil
}

// getSekaidGenesis retrieves the complete Sekaid Genesis data from a target Sekaid node
// by fetching the data in chunks using the Sekaid RPC API.
func (h *httpKiraClient) getSekaidGenesis(ctx context.Context, ipAddress, sekaidRPCport string) ([]byte, error) {
	var completeGenesis []byte
	var chunkCount int64

	for {
		url := fmt.Sprintf("http://%s:%s/%s", ipAddress, sekaidRPCport, fmt.Sprintf("genesis_chunked?chunk=%d", chunkCount))

		chunkedGenesisResponseBody, err := h.doGetHttpQuery(ctx, url)
		if err != nil {
			h.log.Errorf("Querying error: %s", err)
			return nil, err
		}

		var response *types.ResponseChunkedGenesis
		err = json.Unmarshal([]byte(chunkedGenesisResponseBody), &response)
		if err != nil {
			h.log.Errorf("Error parsing JSON: %s", err)
			return nil, err
		}

		totalValue, err := response.Result.Total.Int64()
		if err != nil {
			h.log.Error("Cannot convert `total` field to Integer")
			return nil, err
		}

		decodedData, err := base64.StdEncoding.DecodeString(response.Result.Data)
		if err != nil {
			h.log.Errorf("Decoding Base64 error: %s", err)
			return nil, err
		}

		completeGenesis = append(completeGenesis, decodedData...)

		chunkCount++
		if chunkCount >= totalValue {
			break
		}
	}

	return completeGenesis, nil
}

// getGenSum retrieves the Genesis Sum from a target Interx server
// and returns it as a string after trimming the prefix "0x".
func (h *httpKiraClient) getGenSum(ctx context.Context, ipAddress, interxPort string) (string, error) {
	const genSumPrefix = "0x"
	url := fmt.Sprintf("http://%s:%s/%s", ipAddress, interxPort, "api/gensum")

	body, err := h.doGetHttpQuery(ctx, url)
	if err != nil {
		h.log.Errorf("Querying error: %s", err)
		return "", err
	}

	var result *types.ResponseCheckSum
	err = json.Unmarshal(body, &result)
	if err != nil {
		h.log.Errorf("Error parsing JSON: %s", err)
		return "", err
	}

	trimmedChecksum, err := trimPrefix(result.Checksum, genSumPrefix)
	if err != nil {
		return "", err
	}

	return trimmedChecksum, nil
}

// trimPrefix trims the specified prefix from the given string.
func trimPrefix(s, prefix string) (string, error) {
	if !strings.HasPrefix(s, prefix) {
		return "", &StringPrefixError{
			StringValue: s,
			Prefix:      prefix,
		}
	}

	return s[len(prefix):], nil
}

// getSekaidStatus queries the status of the Sekaid node using an HTTP GET request.
// It constructs the URL based on the provided IP address, RPC port, and endpointStatus, and then makes a GET request
// to retrieve the Sekaid status information. The response is then unmarshaled into a ResponseSekaidStatus struct.
func (h *httpKiraClient) getSekaidStatus(ctx context.Context, ipAddress, rpcPort string) (*types.ResponseSekaidStatus, error) {
	url := fmt.Sprintf("http://%s:%s/%s", ipAddress, rpcPort, endpointStatus)

	body, err := h.doGetHttpQuery(ctx, url)
	if err != nil {
		h.log.Errorf("Querying error: %s", err)
		return nil, err
	}

	var response *types.ResponseSekaidStatus
	err = json.Unmarshal(body, &response)
	if err != nil {
		h.log.Errorf("Can't parse JSON response: %s", err)
		return nil, err
	}

	return response, nil
}

// getPubP2PList queries the public P2P peer list from the Sekaid node using an HTTP GET request.
// It constructs the URL based on the provided IP address, RPC port, and endpointPubP2PList, and then makes a GET request
// to retrieve the public P2P peer list as raw byte data.
func (h *httpKiraClient) getPubP2PList(ctx context.Context, ipAddress, rpcPort string) ([]byte, error) {
	url := fmt.Sprintf("http://%s:%s/%s", ipAddress, rpcPort, endpointPubP2PList)

	body, err := h.doGetHttpQuery(ctx, url)
	if err != nil {
		h.log.Errorf("Querying error: %s", err)
		return nil, err
	}

	return body, nil
}

// getBlockInfo queries block information from a specified RPC server at a given block height using an HTTP GET request.
// It constructs the URL based on the provided RPC server URL and the endpointBlock with the specified minHeight parameter.
// The function then makes an HTTP GET request to retrieve the block information as a ResponseBlock struct.
func (h *httpKiraClient) getBlockInfo(ctx context.Context, rpcServer, blockHeight string) (*types.ResponseBlock, error) {
	endpointBlock := fmt.Sprintf("block?height=%s", blockHeight)

	url := fmt.Sprintf("http://%s/%s", rpcServer, endpointBlock)
	body, err := h.doGetHttpQuery(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("can't reach block response, error: %w", err)
	}

	var response *types.ResponseBlock
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("can't parse JSON response, error: %w", err)
	}

	return response, nil
}

// getInterxGenesis retrieves the Interx Genesis data from a target Interx server.
func (h *httpKiraClient) getInterxGenesis(ctx context.Context, ipAddress, interxPort string) ([]byte, error) {
	url := fmt.Sprintf("http://%s:%s/%s", ipAddress, interxPort, "api/genesis")

	body, err := h.doGetHttpQuery(ctx, url)
	if err != nil {
		h.log.Errorf("Querying error: %s", err)
		return nil, err
	}

	return body, nil
}
