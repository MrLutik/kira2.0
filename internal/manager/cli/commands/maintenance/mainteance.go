package maintenance

import (
	"context"
	"fmt"

	"github.com/mrlutik/kira2.0/internal/config/configFileController"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/maintenance"
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

var log = logging.Log

func Maintenance() *cobra.Command {
	log.Info("Adding `maintenance` command...")
	maintenanceCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, _ []string) {
			if err := validateFlags(cmd); err != nil {
				log.Errorf("Some flag are not valid: %s", err)
				if err := cmd.Help(); err != nil {
					log.Fatalf("Error displaying help: %s", err)
				}
				return
			}
			if err := mainMaitenance(cmd); err != nil {
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
	maintenanceCmd.Flags().Bool(activateFlag, false, "Set this flag to reactivate block validation by node (if node bein deactivated for long period of inaction)")
	maintenanceCmd.Flags().Bool(statusFlag, false, "Set this flag to get curent node status")

	return maintenanceCmd
}

func validateFlags(cmd *cobra.Command) error {
	return nil
}

func mainMaitenance(cmd *cobra.Command) error {
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
	kiraCfg, err := configFileController.ReadOrCreateConfig()
	if err != nil {
		return fmt.Errorf("error while getting kira manager config: %w", err)
	}
	log.Debugf("kira manager cfg: %+v", kiraCfg)
	ctx := context.Background()
	cm, err := docker.NewTestContainerManager()
	if err != nil {
		return fmt.Errorf("error while getting containerManager: %w", err)
	}

	switch {
	case pause:
		err = maintenance.PauseValidator(ctx, kiraCfg, cm)
		if err != nil {
			return err
		}
	case unpause:
		err = maintenance.UnpauseValidator(ctx, kiraCfg, cm)
		if err != nil {
			return err
		}
	case activate:
		err = maintenance.ActivateValidator(ctx, kiraCfg, cm)
		if err != nil {
			return err
		}
	case status:
		valStatus, err := maintenance.GetValidatorStatus(ctx, kiraCfg, cm)
		if err != nil {
			return err
		}
		log.Infof("***Validator Status***\nStatus: %s\nStreak: %s\nRank: %s\n", valStatus.Status, valStatus.Streak, valStatus.Rank)
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
