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
	use   = "maintenance"
	short = "command for maintenance mode"
	long  = "command for maintenance mode: pause for maintenance, unpause, and activate if validator was deactivated"
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

	maintenanceCmd.Flags().Bool("pause", false, "Set this flag to pause block validation by node")
	maintenanceCmd.Flags().Bool("unpause", false, "Set this flag to unpause block validation by node")
	maintenanceCmd.Flags().Bool("activate", false, "Set this flag to reactivate block validation by node (if node bein deactivated for long period of inaction)")
	maintenanceCmd.Flags().Bool("status", false, "Set this flag to get curent node status")

	return maintenanceCmd
}

func validateFlags(cmd *cobra.Command) error {
	return nil
}

func mainMaitenance(cmd *cobra.Command) error {
	pause, err := cmd.Flags().GetBool("pause")
	if err != nil {
		return fmt.Errorf("error while geting <pause> flag")
	}
	unpause, err := cmd.Flags().GetBool("unpause")
	if err != nil {
		return fmt.Errorf("error while geting <unpause> flag")
	}
	activate, err := cmd.Flags().GetBool("activate")
	if err != nil {
		return fmt.Errorf("error while geting <activate> flag")
	}
	status, err := cmd.Flags().GetBool("status")
	if err != nil {
		return fmt.Errorf("error while geting <status> flag")
	}

	err = validateBoolFlags(cmd, pause, unpause, activate, status)
	if err != nil {
		return err
	}
	log.Debugln(pause, unpause, activate, status)
	kiraCfg, err := configFileController.ReadOrCreateConfig()
	if err != nil {
		return fmt.Errorf("error while geting kira manager config: %s", err)
	}
	log.Debugf("kira manager cfg: %+v", kiraCfg)
	ctx := context.Background()
	cm, err := docker.NewTestContainerManager()
	if err != nil {
		return fmt.Errorf("error while geting containerManager: %s", err)
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
		// log.Infof("Validator Status:\nStatus: %s\nStreak: %s\nRank: %s\n", valStatus.Status, valStatus.Streak, valStatus.Rank)
		fmt.Printf("***Validator Status***\nStatus: %s\nStreak: %s\nRank: %s\n", valStatus.Status, valStatus.Streak, valStatus.Rank)
		log.Debugf("valStatus: %+v\n", valStatus)
	}
	return nil
}

func validateBoolFlags(cmd *cobra.Command, flags ...bool) error {
	sum := 0
	for _, val := range flags {
		if val {
			sum++
		}
	}

	if sum > 1 {
		return fmt.Errorf("only one flag at a time is allowed")
	} else if sum == 0 {
		return fmt.Errorf("select flag")
	}
	return nil
}
