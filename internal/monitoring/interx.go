package monitoring

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/mrlutik/kira2.0/internal/types"
)

type (
	// ResponseValopers represents the information about Valopers returned by the Valopers API.
	ResponseValopers struct {
		Status struct {
			ActiveValidators  int `json:"active_validators"`
			TotalValidators   int `json:"total_validators"`
			WaitingValidators int `json:"waiting_validators"`
		} `json:"status"`

		Validators []struct {
			Top     string `json:"top"`
			Address string `json:"address"`
		} `json:"validators"`
	}

	// InterxInfo represents the information about Interx.
	InterxInfo struct {
		NodeID            string
		LatestBlockHeight int
		LatestBlockTime   time.Time
		CatchingUp        bool
	}

	// ValopersInfo represents the information about Valopers.
	ValopersInfo struct {
		ActiveValidators  int
		TotalValidators   int
		WaitingValidators int
	}

	// ResponseBlockStats represents the JSON needed response structure for the BlockStats API.
	ResponseBlockStats struct {
		AverageBlockTime float64 `json:"average_block_time"`
		ConsensusStopped bool    `json:"consensus_stopped"`
	}

	// ConsensusInfo represents the most needed consensus information.
	ConsensusInfo struct {
		BlockTime        float64
		ConsensusStopped bool
	}
)

// doGetValopersQuery performs the Valopers query using the provided HTTP client,
// interxPort, and timeout duration, and returns the parsed response or an error.
func doGetValopersQuery(ctx context.Context, httpClient *http.Client, interxPort string, timeout time.Duration) (*ResponseValopers, error) {
	var response ResponseValopers
	err := doHTTPGetQuery(ctx, httpClient, interxPort, timeout, "api/valopers?all=true", &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetValopersInfo retrieves the information about Valopers using the provided
// context and interxPort, and returns the ValopersInfo or an error.
func (m *MonitoringService) GetValopersInfo(ctx context.Context, interxPort string) (*ValopersInfo, error) {
	response, err := doGetValopersQuery(ctx, m.httpClient, interxPort, time.Second)
	if err != nil {
		log.Errorf("GET query error: %s", err)
		return nil, err
	}

	return &ValopersInfo{
		ActiveValidators:  response.Status.ActiveValidators,
		TotalValidators:   response.Status.TotalValidators,
		WaitingValidators: response.Status.WaitingValidators,
	}, nil
}

// GetTopForValidator retrieves the top value for the validator with the specified
// validatorAddress using the provided context and interxPort. It returns the top
// value or an error if the validator address is not found.
func (m *MonitoringService) GetTopForValidator(ctx context.Context, interxPort, validatorAddress string) (string, error) {
	response, err := doGetValopersQuery(ctx, m.httpClient, interxPort, getQueryTimeout)
	if err != nil {
		log.Errorf("GET query error: %s", err)
		return "", err
	}

	for _, validator := range response.Validators {
		if validator.Address == validatorAddress {
			return validator.Top, nil
		}
	}

	return "", &ValidatorAddressNotFoundError{Address: validatorAddress}
}

// doGetConsensusQuery performs the Consensus query using the provided HTTP client,
// interxPort, and timeout duration, and returns the parsed response or an error.
func doGetConsensusQuery(ctx context.Context, httpClient *http.Client, interxPort string, timeout time.Duration) (*ResponseBlockStats, error) {
	var response ResponseBlockStats
	err := doHTTPGetQuery(ctx, httpClient, interxPort, timeout, "api/consensus", &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetConsensusInfo retrieves the consensus information using the provided context
// and interxPort, and returns the ConsensusInfo or an error.
func (m *MonitoringService) GetConsensusInfo(ctx context.Context, interxPort string) (*ConsensusInfo, error) {
	response, err := doGetConsensusQuery(ctx, m.httpClient, interxPort, getQueryTimeout)
	if err != nil {
		log.Errorf("GET query error: %s", err)
		return nil, err
	}

	return &ConsensusInfo{
		BlockTime:        response.AverageBlockTime,
		ConsensusStopped: response.ConsensusStopped,
	}, nil
}

// doGetInterxStatusQuery performs the InterxStatus query using the provided HTTP client,
// interxPort, and timeout duration, and returns the parsed response or an error.
func doGetInterxStatusQuery(ctx context.Context, httpClient *http.Client, interxPort string, timeout time.Duration) (*types.ResponseInterxStatus, error) {
	var response *types.ResponseInterxStatus
	err := doHTTPGetQuery(ctx, httpClient, interxPort, timeout, "api/kira/status", &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetInterxInfo retrieves the information about Interx using the provided context
// and interxPort, and returns the InterxInfo or an error.
func (m *MonitoringService) GetInterxInfo(ctx context.Context, interxPort string) (*InterxInfo, error) {
	response, err := doGetInterxStatusQuery(ctx, m.httpClient, interxPort, getQueryTimeout)
	if err != nil {
		log.Errorf("GET query error: %s", err)
		return nil, err
	}

	latestBlockHeight, err := strconv.Atoi(response.SyncInfo.LatestBlockHeight)
	if err != nil {
		log.Errorf("Can't parse 'latest_block_height' value, got '%s': %s", response.SyncInfo.LatestBlockHeight, err)
		return nil, err
	}

	return &InterxInfo{
		NodeID:            response.NodeInfo.ID,
		LatestBlockHeight: latestBlockHeight,
		LatestBlockTime:   response.SyncInfo.LatestBlockTime,
		CatchingUp:        response.SyncInfo.CatchingUp,
	}, nil
}
