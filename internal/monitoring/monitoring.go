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

	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/logging"
)

// MonitoringService represents a monitoring service that interacts with the Docker and HTTP clients.
type MonitoringService struct {
	dockerManager    *docker.DockerManager
	containerManager *docker.ContainerManager
	httpClient       *http.Client
}

const (
	gigabyte        = 1024 * 1024 * 1024
	getQueryTimeout = time.Second
)

var log = logging.Log

func NewMonitoringService(dm *docker.DockerManager, cm *docker.ContainerManager) *MonitoringService {
	return &MonitoringService{
		dockerManager:    dm,
		containerManager: cm,
		httpClient:       &http.Client{},
	}
}

// doHTTPGetQuery performs an HTTP GET request to the specified URL using the provided HTTP client,
// and populates the response object with the JSON response.
func doHTTPGetQuery(ctx context.Context, httpClient *http.Client, interxPort string, timeout time.Duration, url string, response any) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	fullURL := fmt.Sprintf("http://localhost:%s/%s", interxPort, url)
	log.Infof("Querying '%s'", fullURL)

	req, err := http.NewRequestWithContext(ctx, "GET", fullURL, nil)
	if err != nil {
		log.Errorf("Can't generate GET query: %s", err)
		return err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Errorf("Can't proceed GET query: %s", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("HTTP request failed with status: %d", resp.StatusCode)
		return fmt.Errorf("HTTP request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Can't read the response body")
		return err
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		log.Errorf("Can't parse JSON response: %s", err)
		return err
	}

	return nil
}
