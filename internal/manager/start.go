package manager

import (
	"context"
	"fmt"

	"github.com/mrlutik/kira2.0/internal/errors"
)

// MustRunSekaid if exist starts sekaid container and running sekaid bin inside
func (s *SekaidManager) MustRunSekaid(ctx context.Context) {
	err := s.containerManager.StartContainer(ctx, s.config.SekaidContainerName)
	errors.HandleFatalErr(fmt.Sprintf("Cannot start '%s' container", s.config.SekaidContainerName), err)
	err = s.startSekaidBinInContainer(ctx)
	errors.HandleFatalErr(fmt.Sprintf("Cannot start 'sekaid' bin in '%s' container", s.config.SekaidContainerName), err)
}

// MustRunInterx if exist starts interx container and running interx bin inside
func (i *InterxManager) MustRunInterx(ctx context.Context) {
	err := i.containerManager.StartContainer(ctx, i.config.InterxContainerName)
	errors.HandleFatalErr(fmt.Sprintf("Cannot start '%s' container", i.config.InterxContainerName), err)
	err = i.startInterxBinInContainer(ctx)
	errors.HandleFatalErr(fmt.Sprintf("Cannot start 'interx' bin in '%s' container", i.config.InterxContainerName), err)

}
