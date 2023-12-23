package manager

import (
	"context"
	"fmt"

	"github.com/mrlutik/kira2.0/internal/errors"
)

func (s *SekaidManager) MustRunSekaid(ctx context.Context) {
	s.containerManager.StartContainer(ctx, s.config.SekaidContainerName)
	err := s.startSekaidBinInContainer(ctx)
	errors.HandleFatalErr(fmt.Sprintf("Cannot start 'sekaid' bin in '%s' container", s.config.SekaidContainerName), err)

	// log.Errorf("Cannot start 'sekaid' bin in '%s' container, error: %s", s.config.SekaidContainerName, err)
}

func (i *InterxManager) MustRunInterx(ctx context.Context) {
	i.containerManager.StartContainer(ctx, i.config.InterxContainerName)
	err := i.startInterxBinInContainer(ctx)
	errors.HandleFatalErr(fmt.Sprintf("Cannot start 'interx' bin in '%s' container", i.config.InterxContainerName), err)

}
