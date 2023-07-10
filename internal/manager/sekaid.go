package manager

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"sigs.k8s.io/yaml"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/types"
)

// SekaidManager represents a manager for Sekaid container and its associated configurations.
type SekaidManager struct {
	ContainerConfig        *container.Config
	SekaiHostConfig        *container.HostConfig
	SekaidNetworkingConfig *network.NetworkingConfig
	dockerManager          *docker.DockerManager
	config                 *config.KiraConfig
}

const (
	timeWaitBetweenBlocks = time.Second * 10
	validatorAccountName  = "validator"
)

// Returns configured SekaidManager.
//
//	*docker.DockerManager // The pointer for docker.DockerManager instance.
//	*config	// Pointer to config struct, can create new instance by calling NewConfig() function
func NewSekaidManager(dockerManager *docker.DockerManager, config *config.KiraConfig) (*SekaidManager, error) {
	log := logging.Log
	log.Infof("Creating sekaid manager with ports: %s, %s, image: '%s', volume: '%s' in '%s' network\n",
		config.GrpcPort, config.RpcPort, config.DockerImageName, config.VolumeName, config.DockerNetworkName)

	natGrpcPort, err := nat.NewPort("tcp", config.GrpcPort)
	if err != nil {
		log.Errorf("Creating NAT GRPC port error: %s", err)
		return nil, err
	}

	natRpcPort, err := nat.NewPort("tcp", config.RpcPort)
	if err != nil {
		log.Errorf("Creating NAT RPC port error: %s", err)
		return nil, err
	}

	sekaiContainerConfig := &container.Config{
		Image:       fmt.Sprintf("%s:%s", config.DockerImageName, config.DockerImageVersion),
		Cmd:         []string{"/bin/bash"},
		Tty:         true,
		AttachStdin: true,
		OpenStdin:   true,
		StdinOnce:   true,
		Hostname:    fmt.Sprintf("%s.local", config.SekaidContainerName),
		ExposedPorts: nat.PortSet{
			natGrpcPort: struct{}{},
			natRpcPort:  struct{}{},
		},
	}

	sekaiHostConfig := &container.HostConfig{
		Binds: []string{
			config.VolumeName,
		},
		PortBindings: nat.PortMap{
			natGrpcPort: []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: config.GrpcPort}},
			natRpcPort:  []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: config.RpcPort}},
		},
		Privileged: true,
	}

	sekaidNetworkingConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			config.DockerNetworkName: {},
		},
	}

	return &SekaidManager{sekaiContainerConfig, sekaiHostConfig, sekaidNetworkingConfig, dockerManager, config}, err
}

// InitSekaidBinInContainer sets up the 'sekaid' container with the specified configurations.
// ctx: The context for the operation.
// Moniker: The Moniker for the 'sekaid' container.
// SekaidContainerName: The name of the 'sekaid' container.
// sekaidNetworkName: The name of the network associated with the 'sekaid' container.
// SekaidHome: The home directory for 'sekaid'.
// KeyringBackend: The keyring backend to use.
// RpcPort: The RPC port for 'sekaid'.
// MnemonicDir: The directory to store the generated mnemonics.
// Returns an error if any issue occurs during the init process.
func (s *SekaidManager) InitSekaidBinInContainer(ctx context.Context) error {
	log := logging.Log
	log.Infof("Setting up '%s' (sekaid) container", s.config.SekaidContainerName)

	commands := []string{
		fmt.Sprintf(`sekaid init  --overwrite --chain-id=%s --home=%s "%s"`,
			s.config.NetworkName, s.config.SekaidHome, s.config.Moniker),
		fmt.Sprintf(`mkdir %s`, s.config.MnemonicDir),
		fmt.Sprintf(`sekaid keys add "%s" --keyring-backend=%s --home=%s --output=json | jq .mnemonic > %s/sekai.mnemonic`,
			validatorAccountName, s.config.KeyringBackend, s.config.SekaidHome, s.config.MnemonicDir),
		fmt.Sprintf(`sekaid keys add "faucet" --keyring-backend=%s --home=%s --output=json | jq .mnemonic > %s/faucet.mnemonic`,
			s.config.KeyringBackend, s.config.SekaidHome, s.config.MnemonicDir),
		fmt.Sprintf(`sekaid add-genesis-account %s 150000000000000ukex,300000000000000test,2000000000000000000000000000samolean,1000000lol --keyring-backend=%v --home=%v`,
			validatorAccountName, s.config.KeyringBackend, s.config.SekaidHome),
		fmt.Sprintf(`sekaid gentx-claim %s --keyring-backend=%s --moniker="%s" --home=%s`,
			validatorAccountName, s.config.KeyringBackend, s.config.Moniker, s.config.SekaidHome),
	}

	for _, command := range commands {
		_, err := s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
		if err != nil {
			log.Errorf("Command '%s' execution error: %s", command, err)
			return err
		}
	}

	log.Infoln("'sekaid' container started")
	return nil
}

// StartSekaidBinInContainer starts sekaid binary inside sekaid container name
// Returns an error if any issue occurs during the start process.
func (s *SekaidManager) StartSekaidBinInContainer(ctx context.Context) error {
	log := logging.Log
	log.Infoln("Starting 'sekaid' container")
	command := fmt.Sprintf(`sekaid start --rpc.laddr "tcp://0.0.0.0:%s" --home=%s`, s.config.RpcPort, s.config.SekaidHome)
	_, err := s.dockerManager.ExecCommandInContainerInDetachMode(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s", command, err)
	}

	return nil
}

// Combine SetupSekaidBinInContainer and StartSekaidBinInContainer together.
// First trying to run sekaid bin from previous state if exist.
// Then checking if sekaid bin running inside container.
// If not initialized new one, then starting again.
// If no sekaid bin running inside container second time - return error.
// Then starting propagating transactions for permissions as in sekai-env.sh
// Returns an error if any issue occurs during the run process.
func (s *SekaidManager) RunSekaidContainer(ctx context.Context) error {
	log := logging.Log

	if err := s.StartSekaidBinInContainer(ctx); err != nil {
		log.Errorf("Cannot start 'sekaid' bin in '%s' container, error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("cannot start 'sekaid' bin in '%s' container, error: %w", s.config.SekaidContainerName, err)
	}

	const delay = time.Second * 1
	log.Warningf("Waiting to start 'sekaid' for %0.0f seconds", delay.Seconds())
	time.Sleep(delay)

	check, _, err := s.dockerManager.CheckIfProcessIsRunningInContainer(ctx, "sekaid", s.config.SekaidContainerName)
	if err != nil {
		log.Errorf("Setup '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("setup '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	if !check {
		if err := s.initializeSekaid(ctx); err != nil {
			return err
		}
	}

	log.Printf("SEKAID CONTAINER STARTED")
	return nil
}

// initializeSekaid initializes the sekaid node by performing several setup steps.
// It starts the sekaid binary for the first time in the specified container,
// initializes a new instance, and waits for it to start.
// If any errors occur during the initialization process, an error is returned.
// The method also propagates genesis proposals and updates the identity registrar from the validator account.
func (s *SekaidManager) initializeSekaid(ctx context.Context) error {
	log := logging.Log

	log.Warningf("Starting sekaid binary first time in '%s' container, initializing new instance", s.config.SekaidContainerName)

	if err := s.InitSekaidBinInContainer(ctx); err != nil {
		log.Errorf("Setup '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("setup '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	if err := s.StartSekaidBinInContainer(ctx); err != nil {
		log.Errorf("Starting 'sekaid' bin in '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("starting 'sekaid' bin in '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	const delay = time.Second * 1
	log.Warningf("Waiting to start 'sekaid' for %0.0f seconds", delay.Seconds())
	time.Sleep(delay)

	check, _, err := s.dockerManager.CheckIfProcessIsRunningInContainer(ctx, "sekaid", s.config.SekaidContainerName)
	if err != nil || !check {
		log.Errorf("Starting 'sekaid' bin second time in '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("starting 'sekaid' bin second time in '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	err = s.PostGenesisProposals(ctx)
	if err != nil {
		log.Errorf("propagating transaction error: %s", err)
		return fmt.Errorf("propagating transaction error: %w", err)
	}

	err = s.updateIdentityRegistrarFromValidator(ctx, validatorAccountName)
	if err != nil {
		log.Errorf("updating identity registrar error: %s", err)
		return fmt.Errorf("updating identity registrar error: %w", err)
	}

	return nil
}

// Post genesis proposals after launching new network from KM1 await-validator-init.sh file.
// Adding required permissions for validator.
// First getting validator address with GetAddressByName.
// Then in loop calling GivePermissionsToAddress func with delay between calls 10 sec because tx can be propagated once per 10 sec
func (s *SekaidManager) PostGenesisProposals(ctx context.Context) error {
	log := logging.Log

	address, err := s.getAddressByName(ctx, validatorAccountName)
	if err != nil {
		log.Errorf("Getting address in '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("getting address in '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	permissions := []int{
		types.PermWhitelistAccountPermissionProposal,
		types.PermRemoveWhitelistedAccountPermissionProposal,
		types.PermCreateUpsertTokenAliasProposal,
		types.PermCreateSoftwareUpgradeProposal,
		types.PermVoteWhitelistAccountPermissionProposal,
		types.PermVoteRemoveWhitelistedAccountPermissionProposal,
		types.PermVoteUpsertTokenAliasProposal,
		types.PermVoteSoftwareUpgradeProposal,
	}
	log.Printf("Permissions to add: '%d' for: '%s'", permissions, address)

	// waiting 10 sec to first block to be propagated
	// TODO await block propagated?
	log.Infof("Waiting for %0.0f seconds before first block be propagated", timeWaitBetweenBlocks.Seconds())
	time.Sleep(timeWaitBetweenBlocks)

	for _, perm := range permissions {
		log.Printf("Adding permission: '%d'", perm)

		err = s.givePermissionsToAddress(ctx, perm, address)
		if err != nil {
			log.Errorf("Giving permission '%d' error: %s", perm, err)
			return fmt.Errorf("giving permission '%d' error: %w", perm, err)
		}

		log.Printf("Checking if '%s' address has '%d' permission", address, perm)
		check, err := s.checkAccountPermission(ctx, perm, address)
		if err != nil {
			log.Errorf("Checking account permission error: %s", err)

			// TODO skip error?
		}
		if !check {
			log.Errorf("Could not find '%d' permission for '%s'", perm, address)

			// TODO skip error?
		}

	}
	return nil
}

// Getting TX by parsing json output of `sekaid query tx <TXhash>`
func (s *SekaidManager) getTxQuery(ctx context.Context, transactionHash string) (types.TxData, error) {
	log := logging.Log
	var data types.TxData

	command := fmt.Sprintf(`sekaid query tx %s -output=json`, transactionHash)
	out, err := s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Couldn't checking tx: '%s'. Command: '%s'. Error: %s", transactionHash, command, err)
		return types.TxData{}, err
	}

	err = json.Unmarshal(out, &data)
	if err != nil {
		log.Errorf("Cannot unmarshaling tx: '%s'. Error: %s", transactionHash, err)
		log.Errorf("Data to unmarshal: %s", string(out))
		return types.TxData{}, fmt.Errorf("unmarshaling '%s' tx error: %w", transactionHash, err)
	}

	log.Debugf("Checking '%s' transaction status: %d. Height: %s", data.Txhash, data.Code, data.Height)
	return data, nil
}

// awaitNextBlock waits for the next block to be propagated and reached by the sekaid node.
// It continuously checks the current block height and compares it with the initial height.
// The method waits for a specified timeout duration for the next block to be reached,
// and returns an error if the timeout is exceeded.
// If the next block is reached within the timeout, the method returns nil.
func (s *SekaidManager) awaitNextBlock(ctx context.Context, timeout time.Duration) error {
	log := logging.Log

	log.Infof("Checking current block height")
	currentBlockHeight, err := s.getBlockHeight(ctx)
	if err != nil {
		return fmt.Errorf("getting current block height error: %w", err)
	}

	log.Infof("Current block height: %s", currentBlockHeight)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	startTime := time.Now()
	for {
		select {
		case <-ticker.C:
			elapsed := time.Since(startTime)
			if elapsed > timeout {
				log.Errorf("Awaiting next block reached timeout: %0.0f seconds", timeout.Seconds())
				return fmt.Errorf("timeout, failed to await next block within %0.2f s limit", timeout.Seconds())
			}

			blockHeight, err := s.getBlockHeight(ctx)
			if err != nil {
				return fmt.Errorf("getting next block height error: %w", err)
			}

			if blockHeight == currentBlockHeight {
				log.Warningf("WAITING: Block is NOT propagated yet: elapsed %0.0f / %0.0f seconds", elapsed.Seconds(), timeout.Seconds())
				continue
			}

			// exit awaiting block
			log.Infof("Next block '%s' reached...", blockHeight)
			return nil

		case <-ctx.Done():
			return fmt.Errorf("awaiting context timeout error: %w", ctx.Err())
		}
	}
}

// nodeStatus is a structure which represents the partial response from `sekaid status`
type nodeStatus struct {
	SyncInfo struct {
		LatestBlockHeight string `json:"latest_block_height"`
	} `json:"SyncInfo"`
}

// getBlockHeight retrieves the latest block height from the sekaid node.
// It executes the "sekaid status" command in the specified container
// and parses the JSON output to extract the block height.
// If successful, it returns the latest block height as a string.
// Otherwise, it returns an error.
func (s *SekaidManager) getBlockHeight(ctx context.Context) (string, error) {
	log := logging.Log

	cmd := "sekaid status"
	statusOutput, err := s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{"bash", "-c", cmd})
	if err != nil {
		return "", fmt.Errorf("getting '%s' error: %s", cmd, err)
	}

	var status nodeStatus
	err = json.Unmarshal(statusOutput, &status)
	if err != nil {
		log.Errorf("Parsing JSON output of '%s' error: %s", cmd, err)
		return "", fmt.Errorf("parsing '%s' error: %w", cmd, err)
	}

	return status.SyncInfo.LatestBlockHeight, nil
}

// Giving permission for chosen address.
// Permissions are ints thats have 0-65 range
//
// Using command: sekaid tx customgov permission whitelist --from "$KM_ACC" --keyring-backend=test --permission="$PERM" --addr="$ADDR" --chain-id=$NETWORK_NAME --home=$SEKAID_HOME --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json | txAwait $TIMEOUT
//
// Then unmarshaling json output and checking sekaid hex of tx
// Then waiting timeWaitBetweenBlocks for tx to propagate in blockchain and checking status code of Tx with GetTxQuery
func (s *SekaidManager) givePermissionsToAddress(ctx context.Context, permissionToAdd int, address string) error {
	log := logging.Log
	command := fmt.Sprintf(`sekaid tx customgov permission whitelist --from %s --keyring-backend=test --permission=%v --addr=%s --chain-id=%s --home=%s --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json`, address, permissionToAdd, address, s.config.NetworkName, s.config.SekaidHome)
	out, err := s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Giving '%d' permission error. Command: '%s'. Error: %s", permissionToAdd, command, err)
		return err
	}
	log.Printf("Permission '%d' is pushed to network for address '%s'", permissionToAdd, address)

	var data types.TxData
	err = json.Unmarshal(out, &data)
	if err != nil {
		log.Errorf("Unmarshaling [%s]Error: %s", string(out), err)
		return err
	}
	log.Debugf("Give permission to address output: Hash: '%s'.Code: %d", data.Txhash, data.Code)

	err = s.awaitNextBlock(ctx, timeWaitBetweenBlocks)
	if err != nil {
		log.Errorf("Awaiting error: %s", err)
		return fmt.Errorf("awaiting error: %s", err)
	}

	txData, err := s.getTxQuery(ctx, data.Txhash)
	if err != nil {
		log.Errorf("Getting transaction query error: %s", err)
		return fmt.Errorf("getting tx query error: %w", err)
	}

	if txData.Code != 0 {
		log.Errorf("Propagating transaction '%s' error. Transaction status: %d", data.Txhash, txData.Code)
		return fmt.Errorf("adding '%d' permission to '%s' address error.\nTransaction hash: '%s'.\nCode: '%d'", permissionToAdd, address, data.Txhash, txData.Code)
	}

	return nil
}

// Checking if account has a specific permission
//
// https://github.com/KiraCore/sekai/blob/master/scripts/sekai-env.sh
//
// sekaid query customgov permissions kira12tptcuw0cp9fccng80vkmqen96npyyrvh2nw5q --output=json --home=/data/.sekai
//
//	permissionToCheck  is a int with 0-65 range
//
// address has to be kira address(not name) : kira12tptcuw0cp9fccng80vkmqen96npyyrvh2nw5q for example, you can get it from local keyring by func GetAddressByName()
func (s *SekaidManager) checkAccountPermission(ctx context.Context, permissionToCheck int, address string) (bool, error) {
	log := logging.Log

	command := fmt.Sprintf("sekaid query customgov permissions %s --output=json", address)
	out, err := s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Executing '%s' command in '%s' container error: %s", command, s.config.SekaidContainerName, err)
		return false, err
	}

	var perms types.AddressPermissions
	err = json.Unmarshal(out, &perms)
	if err != nil {
		log.Errorf("Unmarshaling data error: %s", err)
		log.Errorf("Output: %s", string(out))
		return false, err
	}

	log.Debugf("Checking account permission: %+v", perms)
	for _, perm := range perms.WhiteList {
		if permissionToCheck == perm {
			log.Printf("Permission '%d' was found with '%s' address", permissionToCheck, address)

			return true, nil
		}
	}

	// TODO Warning or Error?
	log.Errorf("Permission '%d' wasn't found with '%s' address", permissionToCheck, address)
	return false, nil
}

// Getting address from keyring.
//
// sekaid keys show validator --keyring-backend=test --home=test
func (s *SekaidManager) getAddressByName(ctx context.Context, addressName string) (string, error) {
	log := logging.Log
	command := fmt.Sprintf("sekaid keys show %s --keyring-backend=%s --home=%s", addressName, s.config.KeyringBackend, s.config.SekaidHome)
	out, err := s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Can't get address by '%s' name. Command: '%s'. Error: %s", addressName, command, err)
		return "", err
	}
	log.Debugf("'keys show %s' command's output:\n%s", addressName, string(out))

	var key []types.SekaidKey
	err = yaml.Unmarshal([]byte(out), &key)
	if err != nil {
		log.Errorf("Cannot unmarshal output to yaml, error: %s", err)
		log.Errorf("Output: %s", string(out))
		return "", err
	}

	log.Printf("Validator address: '%s'", key[0].Address)
	return key[0].Address, nil
}

// Updating identity registrar from KM1 await-validator-init.sh file.
func (s *SekaidManager) updateIdentityRegistrarFromValidator(ctx context.Context, accountName string) error {
	log := logging.Log

	nodeStruct, err := s.getSekaidStatus()
	if err != nil {
		log.Errorf("Getting sekaid status error: %s", err)
		return err
	}

	address, err := s.getAddressByName(ctx, accountName)
	if err != nil {
		log.Errorf("Getting kira address from keyring error: %s", err)
		return err
	}

	records := []struct {
		key   string
		value string
	}{
		{"description", "This is genesis validator account of the KIRA Team"},
		{"social", "https://tg.kira.network,twitter.kira.network"},
		{"contact", "https://support.kira.network"},
		{"website", "https://kira.network"},
		{"username", "KIRA"},
		{"logo", "https://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/kex.svg"},
		{"avatar", "https://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/kex.svg"},
		{"pentest1", "<iframe src=javascript:alert(1)>"},
		{"pentest2", "<img/src=x a='' onerror=alert(2)>"},
		{"pentest3", "<img src=1 onerror=alert(3)>"},
		{"validator_node_id", nodeStruct.Result.NodeInfo.ID},
	}

	for _, record := range records {
		err = s.upsertIdentityRecord(ctx, address, accountName, record.key, record.value)
		if err != nil {
			log.Errorf("Upserting identity record '%+v' error: %s", record, err)
			return err
		}

		log.Infof("Record identity: '%+v' from '%s' is successfully registered", record, accountName)
	}

	log.Infoln("Upserting identity records finished successfully")
	return nil
}

// upsertIdentityRecord  from sekai-utils.sh
func (s *SekaidManager) upsertIdentityRecord(ctx context.Context, address, account, key, value string) error {
	var (
		log = logging.Log
		err error
		out []byte
	)

	if value != "" {
		log.Infof("Registering identity record from address '%s': {'%s': '%s'}", address, key, value)
		command := fmt.Sprintf(`sekaid tx customgov register-identity-records --infos-json="{\"%s\":\"%s\"}" --from=%s --keyring-backend=%s --home=%s --chain-id=%s --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json`, key, value, address, s.config.KeyringBackend, s.config.SekaidHome, s.config.NetworkName)
		out, err = s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{"bash", "-c", command})
		if err != nil {
			log.Errorf("Executing command '%s' in '%s' container error: %s", command, s.config.SekaidContainerName, err)
			return err
		}

	} else {
		log.Infof("Deleting identity record from address '%s': key %s", address, key)
		command := fmt.Sprintf(`sekaid tx customgov delete-identity-records --keys="%s" --from=%s --keyring-backend=%s --home=%s --chain-id=%s --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json`, key, address, s.config.KeyringBackend, s.config.SekaidHome, s.config.NetworkName)
		out, err = s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{"bash", "-c", command})
		if err != nil {
			log.Errorf("Executing command '%s' in '%s' container error: %s", command, s.config.SekaidContainerName, err)
			return err
		}
	}

	var data types.TxData
	err = json.Unmarshal(out, &data)
	if err != nil {
		log.Errorf("Unmarshaling data error: %s", err)
		log.Errorf("Data: %s", string(out))
		return err
	}

	log.Debugf("Register identity record output: Hash: '%s'. Code: %d", data.Txhash, data.Code)

	err = s.awaitNextBlock(ctx, timeWaitBetweenBlocks)
	if err != nil {
		log.Errorf("Awaiting error: %s", err)
		return fmt.Errorf("awaiting error: %s", err)
	}

	txData, err := s.getTxQuery(ctx, data.Txhash)
	if err != nil {
		log.Errorf("Getting transaction query error: %s", err)
		return fmt.Errorf("getting tx query error: %w", err)
	}

	if txData.Code != 0 {
		log.Errorf("The '%s' transaction was executed with error. Code: %d", data.Txhash, txData.Code)
		return fmt.Errorf("the '%s' transaction was executed with error. Code: %d", data.Txhash, txData.Code)
	}

	return nil
}

// func to get status of sekaid node
// same as curl localhost:26657/status (port for sekaid's rpc endpoint)
func (s *SekaidManager) getSekaidStatus() (*types.Status, error) {
	log := logging.Log

	url := fmt.Sprintf("http://localhost:%s/status", s.config.RpcPort)
	log.Println(url)
	response, err := http.Get(url)
	if err != nil {
		log.Errorf("Failed to send GET request: %s", err)
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Errorf("Failed to read response body: %s", err)
		return nil, err
	}

	var statusData *types.Status
	err = json.Unmarshal(body, &statusData)
	if err != nil {
		log.Errorf("Failed to parse JSON: %s", err)
		return nil, err
	}

	return statusData, nil
}
