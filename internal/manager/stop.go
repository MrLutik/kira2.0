package manager

import (
	"context"
	"fmt"
)

// StopSekaid is stopping sekaid process with StopProcessInsideContainer func (signal code - 15) and stopping sekaid container
func (s *SekaidManager) StopSekaid(ctx context.Context) error {
	err := s.containerManager.StopProcessInsideContainer(ctx, "sekaid", 15, s.config.SekaidContainerName)
	if err != nil {
		return fmt.Errorf("stopping sekaid bin in container '%s' error: %w", s.config.SekaidContainerName, err)
	}

	// TODO change method for dockerManager method instead of Cli
	s.log.Infof("Stopping '%s' container", s.config.SekaidContainerName)
	err = s.containerManager.StopContainer(ctx, s.config.SekaidContainerName)
	if err != nil {
		return fmt.Errorf("cannot stop '%s' container, error: %w", s.config.SekaidContainerName, err)
	}
	return nil
}

// StopInterx is stopping interx process with StopProcessInsideContainer func (signal code - 9) and stopping interx container
func (i *InterxManager) StopInterx(ctx context.Context) error {
	err := i.containerManager.StopProcessInsideContainer(ctx, interxProcessName, 9, i.config.InterxContainerName)
	if err != nil {
		return fmt.Errorf("stopping interx bin in container, error: %w", err)
	}

	// TODO change method for dockerManager method instead of Cli
	i.log.Infof("Stopping <%s> container\n", i.config.InterxContainerName)
	err = i.containerManager.StopContainer(ctx, i.config.InterxContainerName)
	if err != nil {
		return fmt.Errorf("cannot stop '%s' container, error: %w", i.config.InterxContainerName, err)
	}
	return nil
}
