package config

import (
	"encoding/json"
	"fmt"

	"github.com/mrlutik/kira2.0/internal/config/configFileController"
	"github.com/mrlutik/kira2.0/internal/errors"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/spf13/cobra"
)

var log = logging.Log

const (
	Use   = "config"
	short = "Reading cfg file of km2"
	long  = "Printing config file of km2 with --print flag in json format"

	PrintFlag = "print"
)

func Config() *cobra.Command {
	log.Info("Adding `firewall` command...")
	configCmd := &cobra.Command{
		Use:   Use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, _ []string) {
			// if err := validateFlags(cmd); err != nil {
			// 	log.Errorf("Some flag are not valid: %s", err)
			// 	cmd.Help()
			// 	return
			// }
			mainConfig(cmd)

		},
	}
	configCmd.Flags().Bool(PrintFlag, false, "Set this flag to print KM2 current config if exist")
	return configCmd
}

func mainConfig(cmd *cobra.Command) {
	log.Println("validating flags")

	printConfig, err := cmd.Flags().GetBool(PrintFlag)
	errors.HandleFatalErr("cannot parse flag", err)
	if printConfig {
		cfg, err := configFileController.ReadOrCreateConfig()
		if err != nil {
			errors.HandleFatalErr("cannot read cfg", err)
			return
		}
		jsonData, err := json.MarshalIndent(cfg, "", "    ")
		if err != nil {
			log.Fatalf("Error marshalling struct to JSON: %s", err)
		}
		fmt.Println(string(jsonData))
		// fmt.Printf("%v", *cfg)
	}

}
