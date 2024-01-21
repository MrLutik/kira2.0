package maintenance

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/client"
	"github.com/mrlutik/kira2.0/internal/config/controller"
	"github.com/mrlutik/kira2.0/internal/config/handler"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/maintenance"
	"github.com/mrlutik/kira2.0/internal/osutils"
	"github.com/mrlutik/kira2.0/internal/utils"
	"github.com/spf13/cobra"
)

const (
	// Command info
	use   = "maintenance"
	short = "command for maintenance mode"
	long  = "command for maintenance mode: pause for maintenance, unpause, and activate if validator was deactivated"

	// Flags
	pauseFlag    = "pause"
	unpauseFlag  = "unpause"
	activateFlag = "activate"
	statusFlag   = "status"
)

func Maintenance(log *logging.Logger) *cobra.Command {
	log.Info("Adding `maintenance` command...")
	maintenanceCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, _ []string) {
			if err := mainMaintenance(cmd, log); err != nil {
				log.Errorf("Error while executing maintenance command: %s", err)
				if err := cmd.Help(); err != nil {
					log.Fatalf("Error displaying help: %s", err)
				}
				return
			}
		},
	}

	maintenanceCmd.Flags().Bool(pauseFlag, false, "Set this flag to pause block validation by node")
	maintenanceCmd.Flags().Bool(unpauseFlag, false, "Set this flag to unpause block validation by node")
	maintenanceCmd.Flags().Bool(activateFlag, false, "Set this flag to reactivate block validation by node (if node being deactivated for long period of inaction)")
	maintenanceCmd.Flags().Bool(statusFlag, false, "Set this flag to get current node status")

	return maintenanceCmd
}

func mainMaintenance(cmd *cobra.Command, log *logging.Logger) error {
	pause, err := cmd.Flags().GetBool(pauseFlag)
	if err != nil {
		return fmt.Errorf("%w: '%s' flag", ErrGettingFlagError, pauseFlag)
	}
	unpause, err := cmd.Flags().GetBool(unpauseFlag)
	if err != nil {
		return fmt.Errorf("%w: '%s' flag", ErrGettingFlagError, unpauseFlag)
	}
	activate, err := cmd.Flags().GetBool(activateFlag)
	if err != nil {
		return fmt.Errorf("%w: '%s' flag", ErrGettingFlagError, activateFlag)
	}
	status, err := cmd.Flags().GetBool(statusFlag)
	if err != nil {
		return fmt.Errorf("%w: '%s' flag", ErrGettingFlagError, statusFlag)
	}

	err = validateBoolFlags(pause, unpause, activate, status)
	if err != nil {
		return err
	}

	log.Debugf("Flags provided: pause - %t, unpause - %t, activate - %t, status - %t", pause, unpause, activate, status)

	utilsOS := osutils.NewOSUtils(log)
	configHandler := handler.NewHandler(utilsOS, log)
	configController := controller.NewConfigController(configHandler, utilsOS, log)

	kiraCfg, err := configController.ReadOrCreateConfig()
	if err != nil {
		return fmt.Errorf("error while getting kira manager config: %w", err)
	}

	log.Debugf("kira manager cfg: %+v", kiraCfg)

	// TODO make flexible setting timeout
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancelFunc()

	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Can't initialize the Docker client: %s", err)
	}

	containerManager := docker.NewTestContainerManager(client, log)
	if err != nil {
		return fmt.Errorf("error while getting containerManager: %w", err)
	}

	helper := utils.NewHelperManager(containerManager, containerManager, utilsOS, kiraCfg, log)
	validatorManager := maintenance.NewValidatorManager(helper, containerManager, kiraCfg, log)

	switch {
	case pause:
		err = validatorManager.PauseValidator(ctx, kiraCfg)
		if err != nil {
			return err
		}
	case unpause:
		err = validatorManager.UnpauseValidator(ctx, kiraCfg)
		if err != nil {
			return err
		}
	case activate:
		err = validatorManager.ActivateValidator(ctx, kiraCfg)
		if err != nil {
			return err
		}
	case status:
		valStatus, err := validatorManager.GetValidatorStatus(ctx)
		if err != nil {
			return err
		}
		log.Infof("Validator Status\nStatus: %s\nStreak: %s\nRank: %s\n", valStatus.Status, valStatus.Streak, valStatus.Rank)
		log.Debugf("valStatus: %+v\n", valStatus)
	}
	return nil
}

func validateBoolFlags(flags ...bool) error {
	sum := 0
	for _, val := range flags {
		if val {
			sum++
		}
	}

	if sum > 1 {
		return ErrOnlyOneFlagAllowed
	} else if sum == 0 {
		return ErrNotSelectFlag
	}
	return nil
}
