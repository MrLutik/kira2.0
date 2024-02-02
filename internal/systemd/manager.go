package systemd

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/go-systemd/v22/dbus"

	"github.com/mrlutik/kira2.0/internal/logging"
)

type ServiceManager struct {
	connection  *dbus.Conn
	serviceName string
	modeAction  string

	log *logging.Logger
}

// NewServiceManager returns new instance of systemd manager
// Callers should call Close() when done with the connection.
func NewServiceManager(ctx context.Context, logger *logging.Logger, serviceName, mode string) (*ServiceManager, error) {
	conn, err := dbus.NewWithContext(ctx)
	if err != nil {
		logger.Errorf("Failed to connect to D-Bus: %s", err)
		return nil, err
	}

	return &ServiceManager{
		connection:  conn,
		serviceName: serviceName,
		modeAction:  mode,
		log:         logger,
	}, nil
}

func (s *ServiceManager) Close() {
	s.connection.Close()
}

func (s *ServiceManager) CheckServiceExists(ctx context.Context) (bool, error) {
	units, err := s.connection.ListUnitsByNamesContext(ctx, []string{s.serviceName})
	if err != nil {
		return false, fmt.Errorf("failed to get list of services: %w", err)
	}

	for _, unit := range units {
		if unit.Name == s.serviceName {
			s.log.Infof("'%s' exists: %t", s.serviceName, true)
			return true, nil
		}
	}

	s.log.Infof("'%s' exists: %t", s.serviceName, false)
	return false, nil
}

// GetServiceStatus retrieves the status of the specified service.
func (s *ServiceManager) GetServiceStatus(ctx context.Context) (string, error) {
	unitStates, err := s.connection.ListUnitsByNamesContext(ctx, []string{s.serviceName})
	if err != nil {
		return "", fmt.Errorf("failed to get list of units: %w", err)
	}

	unitStatus, err := s.connection.GetUnitPropertiesContext(ctx, s.serviceName)
	if err != nil {
		return "", fmt.Errorf("failed to get service status: %w", err)
	}

	active, ok := unitStatus["ActiveState"].(string)
	if !ok {
		return "", fmt.Errorf("failed to determine unit file state: %w", err)
	}

	if active != unitStates[0].ActiveState {
		return "", &DifferentPropertyValuesError{
			PropertyName:  "yourPropertyName",
			ExpectedValue: unitStates[0].ActiveState,
			ActualValue:   active,
		}
	}

	s.log.Infof("'%s' service status is '%s'", s.serviceName, active)
	return unitStates[0].ActiveState, nil
}

// RestartService restarts the specified service.
func (s *ServiceManager) RestartService(ctx context.Context) error {
	ch := make(chan string)

	job, err := s.connection.RestartUnitContext(ctx, s.serviceName, s.modeAction, ch)
	if err != nil {
		return fmt.Errorf("failed to restart service: %w", err)
	}

	s.log.Infof("Job path: %d", job)

	done := make(chan struct{})

	// Monitor the D-Bus channel for a job completion message
	go func() {
		defer close(done)
		defer close(ch)

		for res := range ch {
			s.log.Infof("Restart service operation: %s", res)
			if res != "done" {
				s.log.Infof("Failed to restart service: %s", s.serviceName)
			} else {
				s.log.Infof("Successfully restarted service: %s", s.serviceName)
				return
			}
		}
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ErrServiceRestartCancellationTimeout
	}
}

// StopService stops the specified service.
func (s *ServiceManager) StopService(ctx context.Context) error {
	ch := make(chan string)

	job, err := s.connection.StopUnitContext(ctx, s.serviceName, s.modeAction, ch)
	if err != nil {
		return fmt.Errorf("failed to stop service: %w", err)
	}

	s.log.Infof("Job path: %d", job)

	done := make(chan struct{})

	// Monitor the D-Bus channel for a job completion message
	go func() {
		defer close(done)
		defer close(ch)

		for res := range ch {
			s.log.Infof("Stop service operation: %s", res)
			if res != "done" {
				s.log.Infof("Failed to stop service: %s", s.serviceName)
			} else {
				s.log.Infof("Successfully stopped service: %s", s.serviceName)
				return
			}
		}
	}()

	select {
	case <-done:
		// Completed successfully.
		return nil
	case <-ctx.Done():
		// Timeout or cancellation from the upstream context.
		return fmt.Errorf("stopping service was cancelled or timed out: %w", ErrServiceTimeout)
	}
}

// EnableService enables the specified service to start on boot.
func (s *ServiceManager) EnableService(ctx context.Context, runtime, force bool) error {
	_, changes, err := s.connection.EnableUnitFilesContext(ctx, []string{s.serviceName}, runtime, force)
	if err != nil {
		return fmt.Errorf("failed to enable service: %w", err)
	}

	if len(changes) == 0 {
		s.log.Infof("'%s' is already enabled", s.serviceName)
	}

	for i, change := range changes {
		s.log.Infof("Change [%d]: %+v", i, change)
	}

	return nil
}

// WaitForServiceStatus waits until the service reaches the specified status or times out.
func (s *ServiceManager) WaitForServiceStatus(ctx context.Context, targetStatus string, timeout time.Duration) error {
	s.log.Infof("Waiting for '%s' service status '%s'", s.serviceName, targetStatus)
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("waiting service cancelled: %w", ErrServiceTimeout)
		default:
			status, err := s.GetServiceStatus(ctx)
			if err != nil {
				return fmt.Errorf("failed to get service status: %w", err)
			}

			if status == targetStatus {
				s.log.Infof("'%s': target status '%s' has reached", s.serviceName, targetStatus)
				return nil
			}

			time.Sleep(1 * time.Second)
		}
	}
}
