package join

import (
	"context"
	"fmt"
	"regexp"

	"github.com/spf13/cobra"

	"github.com/mrlutik/kira2.0/internal/adapters"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/errors"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/manager"
	"github.com/mrlutik/kira2.0/internal/systemd"
)

const (
	use   = "join"
	short = "Join to sekaid network"
	long  = "Joining a running sekaid network"
)

// Regular expression to match IPv4 and IPv6 addresses.
const ipRegex = `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$|^(?:[0-9A-Fa-f]{1,4}:){7}[0-9A-Fa-f]{1,4}$`

// Regular expression to match valid port numbers (from 1 to 65535).
const portRegex = `^([1-9]\d{0,4}|[1-5]\d{4}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])$`

var log = logging.Log

func Join() *cobra.Command {
	log.Info("Adding `join` command...")
	joinCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, _ []string) {
			if err := validateFlags(cmd); err != nil {
				log.Errorf("Some flag is not valid: %s", err)
				cmd.Help()
				return
			}
			mainJoin(cmd)
		},
	}

	joinCmd.Flags().String("ip", "", "IP address of the validator to join")
	joinCmd.MarkFlagRequired("ip")

	joinCmd.Flags().String("interx-port", "11000", "Interx port of the validator")
	joinCmd.Flags().String("rpc-port", "26657", "Sekaid RPC port of the validator")
	joinCmd.Flags().String("p2p-port", "26656", "Sekaid P2P port of the validator")

	return joinCmd
}

func validateFlags(cmd *cobra.Command) error {
	ip, err := cmd.Flags().GetString("ip")
	if err != nil {
		return fmt.Errorf("error retrieving 'ip' flag: %s", err)
	}
	if !isValidIP(ip) {
		return fmt.Errorf("'%s' is not a valid IP address", ip)
	}

	interxPort, err := cmd.Flags().GetString("interx-port")
	if err != nil {
		return fmt.Errorf("error retrieving 'interx-port' flag: %s", err)
	}
	if !isValidPort(interxPort) {
		return fmt.Errorf("'%s' is not a valid Interx port", interxPort)
	}

	rpcPort, err := cmd.Flags().GetString("rpc-port")
	if err != nil {
		return fmt.Errorf("error retrieving 'rpc-port' flag: %s", err)
	}
	if !isValidPort(rpcPort) {
		return fmt.Errorf("'%s' is not a valid Sekaid RPC port", rpcPort)
	}

	p2pPort, err := cmd.Flags().GetString("p2p-port")
	if err != nil {
		return fmt.Errorf("error retrieving 'p2p-port' flag: %s", err)
	}
	if !isValidPort(p2pPort) {
		return fmt.Errorf("'%s' is not a valid Sekaid P2P port", p2pPort)
	}

	return nil
}

func isValidIP(ip string) bool {
	match, err := regexp.MatchString(ipRegex, ip)
	if err != nil {
		log.Errorf("Can't match string, error: %s", err)
		return false
	}

	return match
}

func isValidPort(port string) bool {
	match, err := regexp.MatchString(portRegex, port)
	if err != nil {
		log.Errorf("Can't match string, error: %s", err)
		return false
	}

	return match
}

func mainJoin(cmd *cobra.Command) {
	systemd.DockerServiceManagement()

	dockerManager, err := docker.NewTestDockerManager()
	errors.HandleFatalErr("Can't create instance of docker manager", err)
	defer dockerManager.Cli.Close()

	containerManager, err := docker.NewTestContainerManager()
	errors.HandleFatalErr("Can't create instance of container docker manager", err)
	defer containerManager.Cli.Close()

	ctx := context.Background()

	// Skip errors here due to validateFlags method
	ip, _ := cmd.Flags().GetString("ip")
	interxPort, _ := cmd.Flags().GetString("interx-port")
	sekaidRPCPort, _ := cmd.Flags().GetString("rpc-port")
	sekaidP2PPort, _ := cmd.Flags().GetString("p2p-port")

	// Information about validator we need to join
	joinerCfg := &manager.SeedKiraConfig{
		IpAddress:     ip,
		InterxPort:    interxPort,
		SekaidRPCPort: sekaidRPCPort,
		SekaidP2PPort: sekaidP2PPort,
	}
	joinerManager := manager.NewJoinerManager(joinerCfg)

	cfg, err := joinerManager.GenerateKiraConfig(ctx)
	errors.HandleFatalErr("Can't get kira config", err)

	// TODO method called twice
	genesis, err := joinerManager.GetVerifiedGenesisFile(ctx)
	errors.HandleFatalErr("Can't get genesis", err)

	log.Infof("%+v", cfg)

	docker.VerifyingDockerEnvironment(ctx, dockerManager, cfg)

	// TODO Do we need to safe deb packages in temporary directory?
	// Right now the files are downloaded in current directory, where the program starts
	adapters.MustDownloadBinaries(ctx, cfg)

	sekaiManager, err := manager.NewSekaidManager(containerManager, cfg)
	errors.HandleFatalErr("Can't create new 'sekai' manager instance", err)
	sekaiManager.MustInitAndRunJoiner(ctx, genesis)

	interxManager, err := manager.NewInterxManager(containerManager, cfg)
	errors.HandleFatalErr("Can't create new 'interx' manager instance", err)
	interxManager.MustInitAndRunInterx(ctx)
}
