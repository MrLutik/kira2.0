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
	"github.com/mrlutik/kira2.0/internal/errors"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/types"
)

const (
	endpointStatus     = "status"
	endpointPubP2PList = "api/pub_p2p_list?peers_only=true"
)

type SeedKiraConfig struct {
	IpAddress     string
	InterxPort    string
	SekaidRPCPort string
	SekaidP2PPort string
}

type JoinerManager struct {
	client       *http.Client
	targetConfig *SeedKiraConfig
}

func NewJoinerManager(config *SeedKiraConfig) *JoinerManager {
	return &JoinerManager{
		client:       &http.Client{},
		targetConfig: config,
	}
}

// GenerateKiraConfig generates KiraConfig with target information for future connection
func (j *JoinerManager) GenerateKiraConfig(ctx context.Context) (*config.KiraConfig, error) {
	log := logging.Log

	networkInfo, err := j.retrieveNetworkInformation(ctx)
	if err != nil {
		return nil, errors.LogAndReturnErr("Retrieving network information", err)
	}

	configs, err := j.getConfigsBasedOnSeed(ctx, networkInfo)
	if err != nil {
		log.Errorf("Can't get calculated config values: %s", err)
		return nil, err
	}

	cfg := &config.KiraConfig{
		NetworkName:         networkInfo.NetworkName,
		SekaidHome:          "/data/.sekai",
		InterxHome:          "/data/.interx",
		KeyringBackend:      "test",
		DockerImageName:     "ghcr.io/kiracore/docker/kira-base",
		DockerImageVersion:  "v0.13.11",
		DockerNetworkName:   "kira_network",
		SekaiVersion:        "latest", // or v0.3.16
		InterxVersion:       "latest", // or v0.4.33
		SekaidContainerName: "sekaid",
		InterxContainerName: "interx",
		VolumeName:          "kira_volume:/data",
		MnemonicDir:         "~/mnemonics",
		RpcPort:             "26657",
		P2PPort:             "26656",
		GrpcPort:            "9090",
		InterxPort:          "11000",
		Moniker:             "VALIDATOR",
		SekaiDebFileName:    "sekai-linux-amd64.deb",
		InterxDebFileName:   "interx-linux-amd64.deb",
		TimeBetweenBlocks:   time.Second * 10,
		ConfigTomlValues:    configs,
	}

	return cfg, nil
}

// GetVerifiedGenesisFile fetches and verifies the Genesis file from both the 'sekaid' and 'interx' target sources.
func (j *JoinerManager) GetVerifiedGenesisFile(ctx context.Context) ([]byte, error) {
	log := logging.Log

	genesisSekaid, err := j.getSekaidGenesis(ctx)
	if err != nil {
		log.Error("Can't get 'sekaid' genesis")
		return nil, err
	}

	genesisInterx, err := j.getInterxGenesis(ctx)
	if err != nil {
		log.Error("Can't get 'interx' genesis")
		return nil, err
	}

	if err := j.checkFileContentGenesisFiles(genesisInterx, genesisSekaid); err != nil {
		log.Errorf("Comparing genesis files error: %s", err)
		return nil, err
	}

	if err := j.checkGenSum(ctx, genesisSekaid); err != nil {
		return nil, err
	}

	log.Info("Genesis file is valid")
	return genesisSekaid, nil
}

func (j *JoinerManager) getConfigsBasedOnSeed(ctx context.Context, netInfo *NetworkInfo) ([]config.TomlValue, error) {
	configValues := make([]config.TomlValue, 0)

	configValues = append(configValues, config.TomlValue{Tag: "p2p", Name: "seeds", Value: strings.Join(netInfo.Seeds, ",")})

	listOfRPC, err := j.parseRPCfromSeedsList(netInfo.Seeds)
	if err != nil {
		return nil, errors.LogAndReturnErr("Parsing RPCs from seeds list", err)
	}

	syncInfo, err := j.getSyncInfo(ctx, listOfRPC, netInfo.BlockHeight)
	if err != nil {
		return nil, errors.LogAndReturnErr("Getting sync information", err)
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

type NetworkInfo struct {
	NetworkName string
	NodeID      string
	BlockHeight string
	Seeds       []string
}

func (j *JoinerManager) retrieveNetworkInformation(ctx context.Context) (*NetworkInfo, error) {
	log := logging.Log
	statusResponse, err := j.getSekaidStatus(ctx, j.targetConfig.IpAddress, j.targetConfig.SekaidRPCPort)
	if err != nil {
		return nil, errors.LogAndReturnErr("Getting sekaid status", err)
	}

	pupP2PListResponse, err := j.getPubP2PList(ctx, j.targetConfig.IpAddress, j.targetConfig.InterxPort)
	if err != nil {
		return nil, errors.LogAndReturnErr("Getting sekaid public P2P list", err)
	}

	listOfSeeds, err := j.parsePubP2PListResponse(ctx, pupP2PListResponse)
	if err != nil {
		return nil, errors.LogAndReturnErr("Parsing sekaid public P2P list", err)
	}
	if len(listOfSeeds) == 0 {
		log.Warn("List of seeds is empty, the trusted seed will be used")
		listOfSeeds = []string{fmt.Sprintf("tcp://%s@%s:%s", statusResponse.Result.NodeInfo.ID, j.targetConfig.IpAddress, j.targetConfig.SekaidP2PPort)}
	}

	return &NetworkInfo{
		NetworkName: statusResponse.Result.NodeInfo.Network,
		NodeID:      statusResponse.Result.NodeInfo.ID,
		BlockHeight: statusResponse.Result.SyncInfo.LatestBlockHeight,
		Seeds:       listOfSeeds,
	}, nil
}

func (j *JoinerManager) getSekaidStatus(ctx context.Context, ipAddress, rpcPort string) (*types.ResponseSekaidStatus, error) {
	log := logging.Log

	url := fmt.Sprintf("http://%s:%s/%s", ipAddress, rpcPort, endpointStatus)

	body, err := j.doGetHttpQuery(ctx, url)
	if err != nil {
		log.Errorf("Querying error: %s", err)
		return nil, err
	}

	var response *types.ResponseSekaidStatus
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Errorf("Can't parse JSON response: %s", err)
		return nil, err
	}

	return response, nil
}

func (j *JoinerManager) getPubP2PList(ctx context.Context, ipAddress, rpcPort string) ([]byte, error) {
	log := logging.Log

	url := fmt.Sprintf("http://%s:%s/%s", ipAddress, rpcPort, endpointPubP2PList)

	body, err := j.doGetHttpQuery(ctx, url)
	if err != nil {
		log.Errorf("Querying error: %s", err)
		return nil, err
	}

	return body, nil
}

func (j *JoinerManager) parsePubP2PListResponse(ctx context.Context, response []byte) ([]string, error) {
	log := logging.Log

	if len(response) == 0 {
		log.Warning("The list of public seeds is not available")
		return nil, nil
	}

	linesOfPeers := strings.Split(string(response), "\n")
	listOfSeeds := make([]string, 0)

	for _, line := range linesOfPeers {
		formattedSeed := fmt.Sprintf("tcp://%s", line)
		log.Debugf("Got seed: %s", formattedSeed)
		listOfSeeds = append(listOfSeeds, formattedSeed)
	}

	return listOfSeeds, nil
}

type SyncInfo struct {
	rpcServers       []string
	trustHeightBlock string
	trustHashBlock   string
}

func (j *JoinerManager) getSyncInfo(ctx context.Context, listOfRPC []string, minHeight string) (*SyncInfo, error) {
	log := logging.Log

	resultSyncInfo := &SyncInfo{
		rpcServers:       []string{},
		trustHeightBlock: "",
		trustHashBlock:   "",
	}

	for _, rpcServer := range listOfRPC {
		responseBlock, err := j.getBlockInfo(ctx, rpcServer, minHeight)
		if err != nil {
			log.Infof("Can't get block information from RPC '%s'", rpcServer)
			continue
		}

		if responseBlock.Result.Block.Header.Height != minHeight {
			log.Infof("RPC (%s) height is '%s', but expected '%s'", rpcServer, responseBlock.Result.Block.Header.Height, minHeight)
			continue
		}

		if responseBlock.Result.BlockID.Hash != resultSyncInfo.trustHashBlock && resultSyncInfo.trustHashBlock != "" {
			log.Infof("RPC (%s) hash is '%s', but expected '%s'", rpcServer, responseBlock.Result.BlockID.Hash, resultSyncInfo.trustHashBlock)
			continue
		}

		resultSyncInfo.trustHashBlock = responseBlock.Result.BlockID.Hash
		resultSyncInfo.trustHeightBlock = minHeight

		log.Infof("Adding RPC (%s) to RPC connection list", rpcServer)
		resultSyncInfo.rpcServers = append(resultSyncInfo.rpcServers, rpcServer)
	}

	if len(resultSyncInfo.rpcServers) < 2 {
		log.Info("Sync is NOT possible (not enough RPC servers)")
		return nil, nil
	}

	log.Debug(resultSyncInfo)
	return resultSyncInfo, nil
}

func (j *JoinerManager) parseRPCfromSeedsList(seeds []string) ([]string, error) {
	log := logging.Log

	listOfRPCs := make([]string, 0)

	for _, seed := range seeds {
		// tcp://23ca3770ae3874ac8f5a6f84a5cfaa1b39e49fc9@128.140.86.241:26656 -> 128.140.86.241:26657
		parts := strings.Split(seed, "@")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid seed format: '%s'", seed)
		}

		ipAndPort := strings.Split(parts[1], ":")
		if len(ipAndPort) != 2 {
			return nil, fmt.Errorf("invalid IP and Port format in seed: '%s'", seed)
		}

		rpc := fmt.Sprintf("%s:%s", ipAndPort[0], j.targetConfig.SekaidRPCPort)
		log.Infof("Adding rpc to list: %s", rpc)
		listOfRPCs = append(listOfRPCs, rpc)
	}

	return listOfRPCs, nil
}

func (j *JoinerManager) getBlockInfo(ctx context.Context, rpcServer, minHeight string) (*types.ResponseBlock, error) {
	endpointBlock := fmt.Sprintf("block?height=%s", minHeight)

	url := fmt.Sprintf("http://%s/%s", rpcServer, endpointBlock)
	body, err := j.doGetHttpQuery(ctx, url)
	if err != nil {
		return nil, errors.LogAndReturnErr("Can't reach block response", err)
	}

	var response *types.ResponseBlock
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.LogAndReturnErr("Can't parse JSON response", err)
	}

	return response, nil
}

// checkFileContentGenesisFiles checks if the content of two Genesis files is identical.
func (JoinerManager) checkFileContentGenesisFiles(genesis1, genesis2 []byte) error {
	if string(genesis1) != string(genesis2) {
		return fmt.Errorf("genesis files are not identical")
	}

	return nil
}

// checkGenSum checks the integrity of a Genesis file using its SHA256 checksum.
func (j *JoinerManager) checkGenSum(ctx context.Context, genesis []byte) error {
	log := logging.Log

	genesisSum, err := j.getGenSum(ctx)
	if err != nil {
		return fmt.Errorf("can't get genesis check sum: %w", err)
	}

	genSumGenesisHash := sha256.Sum256(genesis)
	hashString := hex.EncodeToString(genSumGenesisHash[:])

	if genesisSum != hashString {
		log.Error("sha256 check sum is not the same")
		return fmt.Errorf("sha256 check sum is not the same")
	}

	return nil
}

// doGetHttpQuery performs an HTTP GET request to the specified URL and returns the response body as a byte slice.
func (j *JoinerManager) doGetHttpQuery(ctx context.Context, url string) ([]byte, error) {
	log := logging.Log

	const timeoutQuery = time.Second * 3

	ctx, cancel := context.WithTimeout(ctx, timeoutQuery)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Errorf("Failed to create request: %s", err)
		return nil, err
	}

	log.Infof("Querying to '%s'", url)
	resp, err := j.client.Do(req)
	if err != nil {
		log.Errorf("Failed to send request: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Failed to read response body: %s", err)
		return nil, err
	}

	log.Debug(string(body[:123]))

	return body, nil
}

// getInterxGenesis retrieves the Interx Genesis data from a target Interx server.
func (j *JoinerManager) getInterxGenesis(ctx context.Context) ([]byte, error) {
	log := logging.Log

	url := fmt.Sprintf("http://%s:%s/%s", j.targetConfig.IpAddress, j.targetConfig.InterxPort, "api/genesis")

	body, err := j.doGetHttpQuery(ctx, url)
	if err != nil {
		log.Errorf("Querying error: %s", err)
		return nil, err
	}

	return body, nil
}

// getSekaidGenesis retrieves the complete Sekaid Genesis data from a target Sekaid node
// by fetching the data in chunks using the Sekaid RPC API.
func (j *JoinerManager) getSekaidGenesis(ctx context.Context) ([]byte, error) {
	log := logging.Log

	var completeGenesis []byte
	var chunkCount int64

	for {
		url := fmt.Sprintf("http://%s:%s/%s", j.targetConfig.IpAddress, j.targetConfig.SekaidRPCPort, fmt.Sprintf("genesis_chunked?chunk=%d", chunkCount))

		chunkedGenesisResponseBody, err := j.doGetHttpQuery(ctx, url)
		if err != nil {
			log.Errorf("Querying error: %s", err)
			return nil, err
		}

		var response *types.ResponseChunkedGenesis
		err = json.Unmarshal([]byte(chunkedGenesisResponseBody), &response)
		if err != nil {
			log.Errorf("Error parsing JSON: %s", err)
			return nil, err
		}

		totalValue, err := response.Result.Total.Int64()
		if err != nil {
			log.Error("Cannot convert `total` field to Integer")
			return nil, err
		}

		decodedData, err := base64.StdEncoding.DecodeString(response.Result.Data)
		if err != nil {
			log.Errorf("Decoding Base64 error: %s", err)
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
func (j *JoinerManager) getGenSum(ctx context.Context) (string, error) {
	log := logging.Log

	const genSumPrefix = "0x"
	url := fmt.Sprintf("http://%s:%s/%s", j.targetConfig.IpAddress, j.targetConfig.InterxPort, "api/gensum")

	body, err := j.doGetHttpQuery(ctx, url)
	if err != nil {
		log.Errorf("Querying error: %s", err)
		return "", err
	}

	var result *types.ResponseCheckSum
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Errorf("Error parsing JSON: %s", err)
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
		return "", fmt.Errorf("input does not have prefix '%s'", prefix)
	}

	return s[len(prefix):], nil
}
