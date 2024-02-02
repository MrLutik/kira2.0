package manager

import (
	"context"
	"fmt"
)

// RunSekaid if exist starts sekaid container and running sekaid bin inside
func (s *SekaidManager) RunSekaid(ctx context.Context) error {
	err := s.containerManager.StartContainer(ctx, s.config.SekaidContainerName)
	if err != nil {
		return fmt.Errorf("cannot start '%s' container, error: %w", s.config.SekaidContainerName, err)
	}

	err = s.startSekaidBinInContainer(ctx)
	if err != nil {
		return fmt.Errorf("cannot start 'sekaid' bin in '%s' container, error: %w", s.config.SekaidContainerName, err)
	}
	return nil
}

// RunInterx if exist starts interx container and running interx bin inside
func (i *InterxManager) RunInterx(ctx context.Context) error {
	err := i.containerManager.StartContainer(ctx, i.config.InterxContainerName)
	if err != nil {
		return fmt.Errorf("cannot start '%s' container, error: %w", i.config.InterxContainerName, err)
	}
	err = i.startInterxBinInContainer(ctx)
	if err != nil {
		return fmt.Errorf("cannot start 'interx' bin in '%s' container, error: %w", i.config.InterxContainerName, err)
	}
	return nil
}
