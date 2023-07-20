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

	"github.com/mrlutik/kira2.0/internal/logging"
)

type JoinerKiraConfig struct {
	IpAddress     string
	InterxPort    string
	SekaidRPCPort string
}

type JoinerManager struct {
	client *http.Client
	config *JoinerKiraConfig
}

func NewJoinerManager(config *JoinerKiraConfig) *JoinerManager {
	return &JoinerManager{
		client: &http.Client{},
		config: config,
	}
}

func (j *JoinerManager) GetGenesis(ctx context.Context) ([]byte, error) {
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

	if err := j.compareGenesisFiles(genesisInterx, genesisSekaid); err != nil {
		return nil, err
	}

	if err := j.checkGenSum(ctx, genesisSekaid); err != nil {
		return nil, err
	}

	log.Info("Genesis file is fine")
	return genesisSekaid, nil
}

func (j *JoinerManager) compareGenesisFiles(genesis1, genesis2 []byte) error {
	log := logging.Log

	if string(genesis1) != string(genesis2) {
		log.Error("Not identical genesis files")
		return fmt.Errorf("genesis files are not identical")
	}

	return nil
}

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

func (j *JoinerManager) doQuery(ctx context.Context, url string) ([]byte, error) {
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

func (j *JoinerManager) getInterxGenesis(ctx context.Context) ([]byte, error) {
	log := logging.Log

	url := fmt.Sprintf("http://%s:%s/%s", j.config.IpAddress, j.config.InterxPort, "api/genesis")

	body, err := j.doQuery(ctx, url)
	if err != nil {
		log.Errorf("Querying error: %s", err)
		return nil, err
	}

	return body, nil
}

type chunkedGenesisResponse struct {
	Result struct {
		Chunk json.Number `json:"chunk"`
		Total json.Number `json:"total"`
		Data  string      `json:"data"`
	} `json:"result"`
}

func (j *JoinerManager) getSekaidGenesis(ctx context.Context) ([]byte, error) {
	log := logging.Log

	var completeGenesis []byte
	var chunkCount int64

	for {
		url := fmt.Sprintf("http://%s:%s/%s", j.config.IpAddress, j.config.SekaidRPCPort, fmt.Sprintf("genesis_chunked?chunk=%d", chunkCount))

		chunkedGenesisResponseBody, err := j.doQuery(ctx, url)
		if err != nil {
			log.Errorf("Querying error: %s", err)
			return nil, err
		}

		var response chunkedGenesisResponse
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

type checkSumResponse struct {
	Checksum string `json:"checksum"`
}

func (j *JoinerManager) getGenSum(ctx context.Context) (string, error) {
	log := logging.Log

	const genSumPrefix = "0x"
	url := fmt.Sprintf("http://%s:%s/%s", j.config.IpAddress, j.config.InterxPort, "api/gensum")

	body, err := j.doQuery(ctx, url)
	if err != nil {
		log.Errorf("Querying error: %s", err)
		return "", err
	}

	var result checkSumResponse
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
