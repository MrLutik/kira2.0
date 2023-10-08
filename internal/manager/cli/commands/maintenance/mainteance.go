package maintenance

import (
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/spf13/cobra"
)

var log = logging.Log

const (
	use   = "maintenance"
	short = "command for maintenance mode"
	long  = "command for maintenance mode: pause for maintenance, unpause, and activate if validator was deactivated"
)

func Maintenance() *cobra.Command {
	log.Info("Adding `maintenance` command...")
	maintenanceCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, _ []string) {
			if err := validateFlags(cmd); err != nil {
				log.Errorf("Some flag are not valid: %s", err)
				cmd.Help()
				return
			}
			mainMaitenance(cmd)
		},
	}

	// maintenanceCmd.AddCommand(openport.OpenPort())

	maintenanceCmd.Flags().Bool("pause", false, "Set this flag to pause block validation by node")
	maintenanceCmd.Flags().Bool("unpause", false, "Set this flag to unpause block validation by node")
	maintenanceCmd.Flags().Bool("activate", false, "Set this flag to reactivate block validation by node (if node bein deactivated for long period of inaction)")
	maintenanceCmd.Flags().Bool("status", false, "Set this flag to get curent node status")

	return maintenanceCmd
}
func validateFlags(cmd *cobra.Command) error {
	return nil
}

func mainMaitenance(cmd *cobra.Command) {
}
