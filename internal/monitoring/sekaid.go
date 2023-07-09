// Sekaid monitoring provides a monitoring service for gathering information
// and performing monitoring operations using various methods and APIs.
package monitoring

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// GetValidatorAddress retrieves the address of the validator using the specified
// sekaid container name, keyring backend, and home directory.
func (m *MonitoringService) GetValidatorAddress(ctx context.Context, sekaidContainerName, keyringBackend, homeDir string) (string, error) {
	cmd := fmt.Sprintf("sekaid keys show validator -a --keyring-backend=%s --home=%s", keyringBackend, homeDir)
	output, err := m.dockerManager.ExecCommandInContainer(ctx, sekaidContainerName, []string{"bash", "-c", cmd})
	if err != nil {
		log.Errorf("Can't execute command '%s', error: '%s'", cmd, err)
		return "", err
	}

	result := strings.ReplaceAll(string(output), "\n", "")
	return result, nil
}

// ResponseSekaidStatus represents the JSON response structure for the Sekaid status query.
type ResponseSekaidStatus struct {
	Result struct {
		NodeInfo struct {
			ID string `json:"id"`
		} `json:"node_info"`
		SyncInfo struct {
			LatestBlockHeight string    `json:"latest_block_height"`
			LatestBlockTime   time.Time `json:"latest_block_time"`
			CatchingUp        bool      `json:"catching_up"`
		} `json:"sync_info"`
	} `json:"result"`
}

// doGetSekaidStatusQuery performs the Sekaid status query using the provided HTTP client,
// sekaid port, and timeout duration, and returns the parsed response or an error.
func doGetSekaidStatusQuery(ctx context.Context, httpClient *http.Client, sekaidPort string, timeout time.Duration) (*ResponseSekaidStatus, error) {
	var response ResponseSekaidStatus
	err := doHTTPGetQuery(ctx, httpClient, sekaidPort, timeout, "status", &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// SekaidInfo represents the needed information about Sekaid.
type SekaidInfo struct {
	NodeID            string
	LatestBlockHeight int
	LatestBlockTime   time.Time
	CatchingUp        bool
}

// GetSekaidInfo retrieves the information about Sekaid using the provided
// context, sekaidPort, and MonitoringService's HTTP client, and returns
// the SekaidInfo or an error.
func (m *MonitoringService) GetSekaidInfo(ctx context.Context, sekaidPort string) (*SekaidInfo, error) {
	response, err := doGetSekaidStatusQuery(ctx, m.httpClient, sekaidPort, getQueryTimeout)
	if err != nil {
		log.Errorf("GET query error: %s", err)
		return nil, err
	}

	latestBlockHeight, err := strconv.Atoi(response.Result.SyncInfo.LatestBlockHeight)
	if err != nil {
		log.Errorf("Can't parse 'latest_block_height' value, got '%s': %s", response.Result.SyncInfo.LatestBlockHeight, err)
		return nil, err
	}

	return &SekaidInfo{
		NodeID:            response.Result.NodeInfo.ID,
		LatestBlockHeight: latestBlockHeight,
		LatestBlockTime:   response.Result.SyncInfo.LatestBlockTime,
		CatchingUp:        response.Result.SyncInfo.CatchingUp,
	}, nil
}
