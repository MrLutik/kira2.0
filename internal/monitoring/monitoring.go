// Package monitoring provides a monitoring service for gathering information
// and performing monitoring operations using the Docker and HTTP clients.
package monitoring

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/mrlutik/kira2.0/internal/logging"
)

type (
	// Service represents a monitoring service that interacts with the Docker and HTTP clients.
	Service struct {
		networkProvider  NetworkInfoProvider
		containerManager ContainerManager
		httpClient       *http.Client

		log *logging.Logger
	}

	NetworkInfoProvider interface {
		GetNetworksInfo(ctx context.Context) ([]types.NetworkResource, error)
	}

	ContainerManager interface {
		GetInspectOfContainer(ctx context.Context, containerIdentification string) (*types.ContainerJSON, error)
		ExecCommandInContainer(ctx context.Context, containerID string, command []string) ([]byte, error)
	}
)

const (
	gigabyte        = 1024 * 1024 * 1024
	getQueryTimeout = time.Second
)

func NewMonitoringService(np NetworkInfoProvider, cm ContainerManager, logger *logging.Logger) *Service {
	return &Service{
		networkProvider:  np,
		containerManager: cm,
		httpClient:       &http.Client{},
		log:              logger,
	}
}

// doHTTPGetQuery performs an HTTP GET request to the specified URL using the provided HTTP client,
// and populates the response object with the JSON response.
func (m *Service) doHTTPGetQuery(ctx context.Context, port string, timeout time.Duration, urlPath string, response any) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	fullLocalhostURL := fmt.Sprintf("http://localhost:%s/%s", port, urlPath)
	m.log.Infof("Querying '%s'", fullLocalhostURL)

	req, err := http.NewRequestWithContext(ctx, "GET", fullLocalhostURL, nil)
	if err != nil {
		m.log.Errorf("Can't generate GET query: %s", err)
		return err
	}

	resp, err := m.httpClient.Do(req)
	if err != nil {
		m.log.Errorf("Can't proceed GET query: %s", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		m.log.Errorf("HTTP request failed with status: %d", resp.StatusCode)
		return &HTTPRequestFailedError{StatusCode: resp.StatusCode}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		m.log.Errorf("Can't read the response body")
		return err
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		m.log.Errorf("Can't parse JSON response: %s", err)
		return err
	}

	return nil
}
