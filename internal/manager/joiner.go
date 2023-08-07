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

	chainID, err := j.getChainID(ctx)
	if err != nil {
		log.Errorf("Can't get network name (chain-id) from genesis, error: %s", err)
		return nil, err
	}

	configs, err := j.getConfigsBasedOnSeed(ctx)
	if err != nil {
		log.Errorf("Can't get calculated config values: %s", err)
		return nil, err
	}

	// TODO How to provide this config from kira_launcher?
	cfg := &config.KiraConfig{
		NetworkName:         chainID,
		SekaidHome:          "/joiner_data/.sekai",
		InterxHome:          "/joiner_data/.interx",
		KeyringBackend:      "test",
		DockerImageName:     "ghcr.io/kiracore/docker/kira-base",
		DockerImageVersion:  "v0.13.11",
		DockerNetworkName:   "joiner_kira_network",
		SekaiVersion:        "latest", // or v0.3.16
		InterxVersion:       "latest", // or v0.4.33
		SekaidContainerName: "joiner_sekaid",
		InterxContainerName: "joiner_interx",
		VolumeName:          "joiner_kira_volume:/joiner_data",
		MnemonicDir:         "~/mnemonics",
		RpcPort:             "36657",
		P2PPort:             "36656",
		GrpcPort:            "9090",
		InterxPort:          "21000",
		Moniker:             "JOINER_VALIDATOR",
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

func (j *JoinerManager) getConfigsBasedOnSeed(ctx context.Context) ([]config.TomlValue, error) {
	log := logging.Log

	configValues := make([]config.TomlValue, 0)

	seeds, err := j.getSeeds(ctx)
	if err != nil {
		log.Errorf("Can't get seeds, error: %s", err)
		return nil, err
	}

	configValues = append(configValues, config.TomlValue{Tag: "p2p", Name: "seeds", Value: seeds})

	// TODO add other calculated fields

	// TODO: State sync is used for 3rd+ joiner
	// // [STATESYNC]
	// {Tag: "statesync", Name: "trust_period", Value: "168h0m0s"},
	// {Tag: "statesync", Name: "enable", Value: "true"},
	// {Tag: "statesync", Name: "temp_dir", Value: "/tmp"},

	return configValues, nil
}

// getChainID retrieves the Chain ID from a target Sekaid node's status endpoint.
func (j *JoinerManager) getChainID(ctx context.Context) (string, error) {
	log := logging.Log

	url := fmt.Sprintf("http://%s:%s/%s", j.targetConfig.IpAddress, j.targetConfig.SekaidRPCPort, "status")

	body, err := j.doGetHttpQuery(ctx, url)
	if err != nil {
		log.Errorf("Querying error: %s", err)
		return "", err
	}

	var response *types.ResponseSekaidStatus
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Errorf("Can't parse JSON response: %s", err)
		return "", err
	}

	return response.Result.NodeInfo.Network, nil
}

// getSeeds retrieves the Seed node information from a target Sekaid node's status endpoint.
func (j *JoinerManager) getSeeds(ctx context.Context) (string, error) {
	log := logging.Log

	url := fmt.Sprintf("http://%s:%s/%s", j.targetConfig.IpAddress, j.targetConfig.SekaidRPCPort, "status")

	body, err := j.doGetHttpQuery(ctx, url)
	if err != nil {
		log.Errorf("Querying error: %s", err)
		return "", err
	}

	var response *types.ResponseSekaidStatus
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Errorf("Can't parse JSON response: %s", err)
		return "", err
	}

	return fmt.Sprintf("tcp://%s@%s:%s", response.Result.NodeInfo.ID, j.targetConfig.IpAddress, j.targetConfig.SekaidP2PPort), nil
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
