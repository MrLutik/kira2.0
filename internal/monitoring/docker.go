package monitoring

import (
	"context"
	"fmt"
	"time"
)

type (
	// DockerNetworkInfo contains needed information about a Docker network.
	DockerNetworkInfo struct {
		Name      string
		IpAddress string
		Subnet    string
	}

	// PortBinding contains the port binding information for a Docker container.
	PortBinding struct {
		HostPort      string
		ContainerPort string
	}

	// ContainerInfo contains information about a Docker container.
	ContainerInfo struct {
		ID          string
		State       string
		Health      string
		Paused      bool
		Restarting  bool
		StartingAt  time.Time
		FinishedAt  time.Time
		Hostname    string
		IP          string
		PortBinding []PortBinding
	}
)

// GetDockerNetwork retrieves information about a Docker network.
// It queries the Docker daemon to obtain a list of network resources and searches for the specified network by name.
func (m *MonitoringService) GetDockerNetwork(ctx context.Context, networkName string) (*DockerNetworkInfo, error) {
	log.Infof("Getting network '%s' resource information", networkName)

	resources, err := m.dockerManager.GetNetworksInfo(ctx)
	if err != nil {
		log.Errorf("Getting network info error: %s", err)
		return nil, err
	}

	for _, network := range resources {
		if network.Name == networkName {
			return &DockerNetworkInfo{
				Name:      network.Name,
				IpAddress: network.IPAM.Config[0].Gateway,
				Subnet:    network.IPAM.Config[0].Subnet,
			}, nil
		}
	}
	return nil, &NetworkNotAvailableError{NetworkName: networkName}
}

// GetContainerInfo retrieves information about a Docker container and its network bindings.
// It returns a ContainerInfo struct containing the container's details, or an error if the operation fails.
//
// The method retrieves the container's information using the `GetInspectOfContainer` method of the `dockerManager`,
// and then parses the container's start and finish times using time.Parse with the RFC3339Nano format.
// The container's port bindings are extracted from the inspect result and stored in a slice of PortBinding structs.
// Finally, a ContainerInfo struct is constructed and returned with the retrieved information.
// If the network specified by `dockerNetworkName` does not exist in the container's network settings,
// an error is returned indicating the absence of the network.
func (m *MonitoringService) GetContainerInfo(ctx context.Context, containerName, dockerNetworkName string) (*ContainerInfo, error) {
	resultInspect, err := m.containerManager.GetInspectOfContainer(ctx, containerName)
	if err != nil {
		log.Errorf("Getting containers inspection error: %s", err)
		return nil, err
	}

	startingAt, err := time.Parse(time.RFC3339Nano, resultInspect.State.StartedAt)
	if err != nil {
		log.Errorf("Can't parse date: %s", resultInspect.State.StartedAt)
		return nil, fmt.Errorf("can't parse date '%s', error: %w", resultInspect.State.StartedAt, err)
	}

	finishedAt, err := time.Parse(time.RFC3339Nano, resultInspect.State.FinishedAt)
	if err != nil {
		log.Errorf("Can't parse date: %s", resultInspect.State.FinishedAt)
		return nil, fmt.Errorf("can't parse date '%s', error: %w", resultInspect.State.FinishedAt, err)
	}

	portBinding := make([]PortBinding, 0, len(resultInspect.NetworkSettings.Ports))
	for containerPort, hostPortBinding := range resultInspect.NetworkSettings.Ports {
		portBinding = append(portBinding, PortBinding{
			HostPort:      hostPortBinding[0].HostPort,
			ContainerPort: containerPort.Port(),
		})
	}

	network, exists := resultInspect.NetworkSettings.Networks[dockerNetworkName]
	if !exists {
		return nil, ErrNetworkDoesNotExist
	}

	r := &ContainerInfo{
		ID:    resultInspect.ID,
		State: resultInspect.State.Status,
		// Health:      resultInspect.State.Health.Status,
		Paused:      resultInspect.State.Paused,
		Restarting:  resultInspect.State.Restarting,
		StartingAt:  startingAt,
		FinishedAt:  finishedAt,
		Hostname:    resultInspect.Config.Hostname,
		IP:          network.IPAddress,
		PortBinding: portBinding,
	}
	log.Debugf("%+v", *r)
	return r, nil
}
