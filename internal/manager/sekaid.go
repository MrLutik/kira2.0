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
}

const (
	timeWaitBetweenBlocks = time.Second * 10
	validatorAccountName  = "validator"
)

// Returns configured SekaidManager.
// *docker.DockerManager: The pointer for docker.DockerManager instance.
// grpcPort: The grpc port for sekaid.
// rpcPort: The rpc port for sekaid.
// imageName: The name of a image that sekaid container will be using.
// volumeName: The name of a volume that sekaid container will be using.
// dockerNetworkName: The name of a docker network that sekaid container will be using.
// containerName: The name of a container that sekaid will have.
func NewSekaidManager(dockerManager *docker.DockerManager, grpcPort, rpcPort, imageName, volumeName, dockerNetworkName, containerName string) (*SekaidManager, error) {
	log := logging.Log
	log.Infof("Creating sekaid manager with ports: %s, %s, image: '%s', volume: '%s' in '%s' network\n", grpcPort, rpcPort, imageName, volumeName, dockerNetworkName)

	natGrcpPort, err := nat.NewPort("tcp", grpcPort)
	if err != nil {
		log.Errorf("Creating NAT GRPC port error: %s\n", err)
		return nil, err
	}

	natRpcPort, err := nat.NewPort("tcp", rpcPort)
	if err != nil {
		log.Errorf("Creating NAT RPC port error: %s\n", err)
		return nil, err
	}

	sekaiContainerConfig := &container.Config{
		Image:       imageName,
		Cmd:         []string{"/bin/bash"},
		Tty:         true,
		AttachStdin: true,
		OpenStdin:   true,
		StdinOnce:   true,
		Hostname:    fmt.Sprintf("%s.local", containerName),
		ExposedPorts: nat.PortSet{
			natGrcpPort: struct{}{},
			natRpcPort:  struct{}{},
		},
	}

	sekaiHostConfig := &container.HostConfig{
		Binds: []string{
			volumeName,
		},
		PortBindings: nat.PortMap{
			natGrcpPort: []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: grpcPort}},
			natRpcPort:  []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: rpcPort}},
		},
		Privileged: true,
	}

	sekaidNetworkingConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			dockerNetworkName: {},
		},
	}

	return &SekaidManager{sekaiContainerConfig, sekaiHostConfig, sekaidNetworkingConfig, dockerManager}, err
}

// InitSekaidBinInContainer sets up the 'sekaid' container with the specified configurations.
// ctx: The context for the operation.
// moniker: The moniker for the 'sekaid' container.
// sekaidContainerName: The name of the 'sekaid' container.
// sekaidNetworkName: The name of the network associated with the 'sekaid' container.
// sekaidHome: The home directory for 'sekaid'.
// keyringBackend: The keyring backend to use.
// rcpPort: The RPC port for 'sekaid'.
// mnemonicDir: The directory to store the generated mnemonics.
// Returns an error if any issue occurs during the init process.
func (s *SekaidManager) InitSekaidBinInContainer(ctx context.Context, moniker, sekaidContainerName, sekaidNetworkName, sekaidHome, keyringBackend, rcpPort, mnemonicDir string) error {
	log := logging.Log
	log.Infoln("Setting up 'sekaid' container")

	command := fmt.Sprintf(`sekaid init  --overwrite --chain-id=%s --home=%s "%s"`, sekaidNetworkName, sekaidHome, moniker)
	_, err := s.dockerManager.ExecCommandInContainer(ctx, sekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}

	command = fmt.Sprintf(`mkdir %s`, mnemonicDir)
	_, err = s.dockerManager.ExecCommandInContainer(ctx, sekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}

	command = fmt.Sprintf(`sekaid keys add "%s" --keyring-backend=%s --home=%s --output=json | jq .mnemonic > %s/sekai.mnemonic`, validatorAccountName, keyringBackend, sekaidHome, mnemonicDir)
	_, err = s.dockerManager.ExecCommandInContainer(ctx, sekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}

	command = fmt.Sprintf(`sekaid keys add "faucet" --keyring-backend=%s --home=%s --output=json | jq .mnemonic > %s/faucet.mnemonic`, keyringBackend, sekaidHome, mnemonicDir)
	_, err = s.dockerManager.ExecCommandInContainer(ctx, sekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}

	command = fmt.Sprintf(`sekaid add-genesis-account %s 150000000000000ukex,300000000000000test,2000000000000000000000000000samolean,1000000lol --keyring-backend=%s --home=%s`, validatorAccountName, keyringBackend, sekaidHome)
	_, err = s.dockerManager.ExecCommandInContainer(ctx, sekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}

	command = fmt.Sprintf(`sekaid gentx-claim %s --keyring-backend=%s --moniker="%s" --home=%s`, validatorAccountName, keyringBackend, moniker, sekaidHome)
	_, err = s.dockerManager.ExecCommandInContainer(ctx, sekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}

	log.Infoln("'sekaid' container started")
	return nil
}

// StartSekaidBinInContainer starts sekaid binary inside sekaidContainerName var
// ctx: The context for the operation.
// moniker: The moniker for the 'sekaid' container.
// sekaidContainerName: The name of the 'sekaid' container.
// sekaidNetworkName: The name of the network associated with the 'sekaid' container.
// sekaidHome: The home directory for 'sekaid'.
// keyringBackend: The keyring backend to use.
// rcpPort: The RPC port for 'sekaid'.
// mnemonicDir: The directory to store the generated mnemonics.
// Returns an error if any issue occurs during the start process.
func (s *SekaidManager) StartSekaidBinInContainer(ctx context.Context, moniker, sekaidContainerName, sekaidNetworkName, sekaidHome, keyringBackend, rcpPort, mnemonicDir string) error {
	log := logging.Log
	log.Infoln("Starting 'sekaid' container")
	command := fmt.Sprintf(`sekaid start --rpc.laddr "tcp://0.0.0.0:%s" --home=%s`, rcpPort, sekaidHome)
	_, err := s.dockerManager.ExecCommandInContainerInDetachMode(ctx, sekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
	}

	return nil
}

// Combine SetupSekaidBinInContainer and StartSekaidBinInContainer together.
// First trying to run sekaid bin from previous state if exist.
// Then checking if sekaid bin running inside container.
// If no initing new one.
// Then starting again.
// If no sekaid bin running inside container second time - return error.
// Then starting propagating transactions for permissions as in sekai-env.sh
// ctx: The context for the operation.
// moniker: The moniker for the 'sekaid' container.
// sekaidContainerName: The name of the 'sekaid' container.
// sekaidNetworkName: The name of the network associated with the 'sekaid' container.
// sekaidHome: The home directory for 'sekaid'.
// keyringBackend: The keyring backend to use.
// rcpPort: The RPC port for 'sekaid'.
// mnemonicDir: The directory to store the generated mnemonics.
// Returns an error if any issue occurs during the run process.
func (s *SekaidManager) RunSekaidContainer(ctx context.Context, moniker, sekaidContainerName, sekaidNetworkName, sekaidHome, keyringBackend, rcpPort, mnemonicDir string) error {
	log := logging.Log
	err := s.StartSekaidBinInContainer(ctx, moniker, sekaidContainerName, sekaidNetworkName, sekaidHome, keyringBackend, rcpPort, mnemonicDir)
	if err != nil {
		log.Errorf("Cannot start sekaid bin in %s container", sekaidContainerName)
	}
	time.Sleep(time.Second * 1)
	check, _, err := s.dockerManager.CheckIfProcessIsRunningInContainer(ctx, "sekaid", sekaidContainerName)
	if err != nil {
		log.Infof("Error while setup '%s' container: %s\n", sekaidContainerName, err)
		return err
	}
	if !check {
		log.Infof("Error starting sekaid binary first time in '%s' container, initing new instance\n", sekaidContainerName)
		err = s.InitSekaidBinInContainer(ctx, moniker, sekaidContainerName, sekaidNetworkName, sekaidHome, keyringBackend, rcpPort, mnemonicDir)
		if err != nil {
			log.Errorf("Error while setup '%s' container: %s\n", sekaidContainerName, err)
			return err
		}
		err := s.StartSekaidBinInContainer(ctx, moniker, sekaidContainerName, sekaidNetworkName, sekaidHome, keyringBackend, rcpPort, mnemonicDir)
		if err != nil {
			log.Errorf("Cannot start sekaid bin in %s container", sekaidContainerName)
		}

		time.Sleep(time.Second * 1)

		check, _, err = s.dockerManager.CheckIfProcessIsRunningInContainer(ctx, "sekaid", sekaidContainerName)
		if err != nil {
			log.Errorf("Error while setup '%s' container: %s\n", sekaidContainerName, err)
			return err
		}
		if !check {
			log.Errorf("Error starting sekaid bin second time in '%s' container\n", sekaidContainerName)
			return fmt.Errorf("couldn't start sekaid bin second time")
		}
		err = s.PostGenesisProposals(ctx, sekaidContainerName, sekaidHome, sekaidNetworkName, keyringBackend)
		if err != nil {
			log.Errorf("Error while propagating transaction: %s \n", err)
			return err
		}
		err = s.UpdateIdentityRegistrar(ctx, validatorAccountName, sekaidContainerName, sekaidHome, keyringBackend, sekaidNetworkName, rcpPort)
		if err != nil {
			log.Errorf("Error while updating identity registrar %s \n", err)
			return err
		}

	}
	log.Printf("SEKAID CONTAINER STARTED\n\n\n")
	return nil
}

// Post genesis proposals after launching new network from KM1 await-validator-init.sh file.
// Adding required permissions for validator.
// First getting validator address with GetAddressByName.
// Then in loop calling GivePermissionsToAddress func with delay between calls 10 sec because tx can be propagated once per 10 sec
func (s *SekaidManager) PostGenesisProposals(ctx context.Context, sekaidContainerName, sekaidHome, networkName, keyringBackend string) error {
	log := logging.Log
	address, err := s.GetAddressByName(ctx, validatorAccountName, sekaidContainerName, sekaidHome, keyringBackend)
	if err != nil {
		log.Errorf("Error while getting address in '%s' container: %s", sekaidContainerName, err)
		return fmt.Errorf("error while getting address in '%s' container: %s", sekaidContainerName, err)
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

	time.Sleep(timeWaitBetweenBlocks)
	for _, perm := range permissions {
		log.Printf("Adding permission: '%d'", perm)
		err = s.GivePermissionsToAddress(ctx, perm, address, sekaidContainerName, sekaidHome, networkName)
		if err != nil {
			log.Errorf("Giving permission error: %s", err)
		}

		log.Printf("Checking if '%s' address has '%d' permission", address, perm)
		check, err := s.CheckAccountPermission(ctx, perm, address, sekaidContainerName, sekaidHome)
		if err != nil {
			log.Errorf("Checking account permission error: %s", err)

			// TODO skip error?
		}
		if !check {
			log.Errorf("Could not find '%d' permission for '%s'", perm, address)

			// TODO skip error?
		}
		// time.Sleep(timeWaitBetweenBlocks)

	}
	return nil
}

// Getting TX by parsing json output of `sekaid query tx <TXhash>`
func (s *SekaidManager) GetTxQuery(ctx context.Context, transactionHash, sekaidContainerName string) (types.TxData, error) {
	log := logging.Log
	var data types.TxData

	command := fmt.Sprintf(`sekaid query tx %s -output=json`, transactionHash)
	out, err := s.dockerManager.ExecCommandInContainer(ctx, sekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Couldn't checking tx: '%s'. Command: '%s'. Error:%s\n", transactionHash, command, err)
		return types.TxData{}, err
	}

	err = json.Unmarshal(out, &data)
	if err != nil {
		log.Errorf("Cannot unmarshaling tx: '%s'\nData to unmarshal: %s\nError: %s", transactionHash, string(out), err)
		return types.TxData{}, err
	}

	log.Debugf("Checking transaction status: %+v", data)
	return data, nil
}

func (s *SekaidManager) awaitNextBlock(ctx context.Context, sekaidContainerName string, timeout time.Duration) error {
	log := logging.Log

	currentBlockHeight, err := s.getBlockHeight(ctx, sekaidContainerName)
	if err != nil {
		return fmt.Errorf("getting current block height error: %s", err)
	}

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

			blockHeight, err := s.getBlockHeight(ctx, sekaidContainerName)
			if err != nil {
				return fmt.Errorf("getting next block height error: %s", err)
			}

			if blockHeight == currentBlockHeight {
				log.Warningf("WAITING: Block is NOT propagated yet: elapsed %0.0f / %0.0f seconds\n", elapsed.Seconds(), timeout.Seconds())
				continue
			}

			// exit awaiting block
			return nil

		case <-ctx.Done():
			return fmt.Errorf("awaiting timeout error: %s", ctx.Err())
		}
	}
}

type NodeStatus struct {
	SyncInfo struct {
		LatestBlockHeight string `json:"latest_block_height"`
	} `json:"SyncInfo"`
}

func (s *SekaidManager) getBlockHeight(ctx context.Context, sekaidContainerName string) (string, error) {
	log := logging.Log

	cmd := "sekaid status"
	statusOutput, err := s.dockerManager.ExecCommandInContainer(ctx, sekaidContainerName, []string{"bash", "-c", cmd})
	if err != nil {
		return "", fmt.Errorf("getting '%s' error: %s", cmd, err)
	}

	var status NodeStatus
	err = json.Unmarshal(statusOutput, &status)
	if err != nil {
		log.Errorf("Parsing JSON output of '%s' error: %s", cmd, err)
		return "", fmt.Errorf("parsing '%s' error: %s", cmd, err)
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
func (s *SekaidManager) GivePermissionsToAddress(ctx context.Context, permissionToAdd int, address, sekaidContainerName, sekaidHome, networkName string) error {
	log := logging.Log
	command := fmt.Sprintf(`sekaid tx customgov permission whitelist --from %s --keyring-backend=test --permission=%d --addr=%s --chain-id=%s --home=%s --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json`, address, permissionToAdd, address, networkName, sekaidHome)
	out, err := s.dockerManager.ExecCommandInContainer(ctx, sekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("error when giving '%d' permission. Command: '%s'", permissionToAdd, command)
	}
	log.Printf("Permission '%d' is pushed to network for address '%s'", permissionToAdd, address)

	var data types.TxData
	err = json.Unmarshal(out, &data)
	if err != nil {
		log.Errorf("Unmarshaling [%s]\nError: %s", string(out), err)
		return err
	}
	log.Debugf("Give permission to address output: %+v", data)

	err = s.awaitNextBlock(ctx, sekaidContainerName, timeWaitBetweenBlocks)
	if err != nil {
		log.Errorf("Awaiting error: %s", err)
		return fmt.Errorf("awaiting error: %s", err)
	}

	txData, err := s.GetTxQuery(ctx, data.Txhash, sekaidContainerName)
	if err != nil {
		log.Errorf("Getting transaction query error: %s", err)
		return fmt.Errorf("getting tx query error: %s", err)
	}

	if txData.Code != 0 {
		log.Errorf("Propagating transaction. Transaction status: %d", txData.Code)
		return fmt.Errorf("error in adding '%d' permission to '%s' address.\nTransaction hash: '%s'\nCode: '%d'", permissionToAdd, address, data.Txhash, txData.Code)
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
func (s *SekaidManager) CheckAccountPermission(ctx context.Context, permissionToCheck int, address, sekaidContainerName, sekaidHome string) (bool, error) {
	log := logging.Log

	command := fmt.Sprintf("sekaid query customgov permissions %s --output=json --home=%s", address, sekaidHome)
	out, err := s.dockerManager.ExecCommandInContainer(ctx, sekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("")
		return false, err
	}

	var perms types.AddressPermissions
	err = json.Unmarshal(out, &perms)
	if err != nil {
		log.Errorf("Error unmarshaling:%s \n%s", string(out), err)
		return false, err
	}

	log.Debugf("Checking account permission: %+v", perms)
	for _, perm := range perms.WhiteList {
		if permissionToCheck == perm {
			log.Printf("Permission '%d' was found with '%s' address", permissionToCheck, address)

			return true, nil
		}
	}

	log.Printf("No '%d' permission were found with %s address", permissionToCheck, address)
	return false, nil
}

// Getting address from keyring.
//
// sekaid keys show validator --keyring-backend=test --home=test
func (s *SekaidManager) GetAddressByName(ctx context.Context, addressName, sekaidContainerName, sekaidHome, keyringBackend string) (string, error) {
	log := logging.Log
	command := fmt.Sprintf("sekaid keys show %s --keyring-backend=%s --home=%s", addressName, keyringBackend, sekaidHome)
	out, err := s.dockerManager.ExecCommandInContainer(ctx, sekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Can't get address by '%s' name. Command: '%s'. Error: %s", addressName, command, err)
		return "", err
	}
	log.Debugf("'keys show %s' command's output:\n%s", addressName, string(out))

	var key []types.SekaidKey
	err = yaml.Unmarshal([]byte(out), &key)
	if err != nil {
		log.Fatalf("Cannot unmarshal output to yaml, error: %s", err)
		return "", err
	}
	if len(key) <= 0 {
		log.Errorf(`no keys were found in keyring with "%s" name`, addressName)
		return "", fmt.Errorf("no keys were found in keyring with %s name", addressName)
	}
	log.Printf("val address: %s\n", key[0].Address)
	return key[0].Address, nil
}

// Updating identity registrar from KM1 await-validator-init.sh file.
func (s *SekaidManager) UpdateIdentityRegistrar(ctx context.Context, account, sekaidContainerName, sekaidHome, keyringBackend, networkName, rpcPorrt string) error {
	log := logging.Log
	nodeStruct, err := s.GetSekaidStatus(sekaidContainerName, rpcPorrt)
	if err != nil {
		log.Errorf("")
		return err
	}

	s.UpsertIdentityRecord(ctx, account, "description", "This is genesis validator account of the KIRA Team", sekaidContainerName, sekaidHome, keyringBackend, networkName)
	s.UpsertIdentityRecord(ctx, account, "social", "https://tg.kira.network,twitter.kira.network", sekaidContainerName, sekaidHome, keyringBackend, networkName)
	s.UpsertIdentityRecord(ctx, account, "contact", "https://support.kira.network", sekaidContainerName, sekaidHome, keyringBackend, networkName)
	s.UpsertIdentityRecord(ctx, account, "website", "https://kira.network", sekaidContainerName, sekaidHome, keyringBackend, networkName)
	s.UpsertIdentityRecord(ctx, account, "username", "KIRA", sekaidContainerName, sekaidHome, keyringBackend, networkName)
	s.UpsertIdentityRecord(ctx, account, "logo", "https://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/kex.svg", sekaidContainerName, sekaidHome, keyringBackend, networkName)
	s.UpsertIdentityRecord(ctx, account, "avatar", "https://kira-network.s3-eu-west-1.amazonaws.com/assets/img/tokens/kex.svg", sekaidContainerName, sekaidHome, keyringBackend, networkName)
	s.UpsertIdentityRecord(ctx, account, "pentest1", "<iframe src=javascript:alert(1)>", sekaidContainerName, sekaidHome, keyringBackend, networkName)
	s.UpsertIdentityRecord(ctx, account, "pentest2", "<img/src=x a='' onerror=alert(2)>", sekaidContainerName, sekaidHome, keyringBackend, networkName)
	s.UpsertIdentityRecord(ctx, account, "pentest3", "<img src=1 onerror=alert(3)>", sekaidContainerName, sekaidHome, keyringBackend, networkName)
	s.UpsertIdentityRecord(ctx, account, "validator_node_id", nodeStruct.Result.NodeInfo.ID, sekaidContainerName, sekaidHome, keyringBackend, networkName)

	// s.UpsertIdentityRecord(ctx, "test", "username", "test", sekaidContainerName, sekaidHome, keyringBackend, networkName)
	// s.UpsertIdentityRecord(ctx, "signer", "username", "faucet", sekaidContainerName, sekaidHome, keyringBackend, networkName)
	return nil
}

// upsertIdentityRecord  from sekai-utils.sh
func (s *SekaidManager) UpsertIdentityRecord(ctx context.Context, account, key, value, sekaidContainerName, sekaidHome, keyringBackend, networkName string) error {
	log := logging.Log
	address, err := s.GetAddressByName(ctx, account, sekaidContainerName, sekaidHome, keyringBackend)
	if err != nil {
		log.Errorf("")
		return err
	}
	var out []byte
	if value != "" {
		command := fmt.Sprintf(`sekaid tx customgov register-identity-records --infos-json="{\"%s\":\"%s\"}" --from=%s --keyring-backend=%s --home=%s --chain-id=%s --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json`, key, value, address, keyringBackend, sekaidHome, networkName)
		out, err = s.dockerManager.ExecCommandInContainer(ctx, sekaidContainerName, []string{`bash`, `-c`, command})
		if err != nil {
			log.Errorf("")
			return err
		}
		log.Printf("%s\n", string(out))

	} else {
		command := fmt.Sprintf(`sekaid tx customgov delete-identity-records --keys="%s" --from=%s --keyring-backend=%s --home=%s --chain-id=%s --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json`, key, address, keyringBackend, sekaidHome, networkName)
		out, err = s.dockerManager.ExecCommandInContainer(ctx, sekaidContainerName, []string{`bash`, `-c`, command})
		if err != nil {
			log.Errorf("")
			return err
		}
		log.Printf("%s\n", string(out))
	}
	var data types.TxData
	err = json.Unmarshal(out, &data)
	if err != nil {
		log.Errorf("Error unmarshaling:%s \n%s", string(out), err)
		return err
	}
	time.Sleep(timeWaitBetweenBlocks)

	txData, err := s.GetTxQuery(ctx, data.Txhash, sekaidContainerName)
	if err != nil {
		log.Errorf("")
		return err
	}
	if txData.Code != 0 {
		log.Errorf("")
		return fmt.Errorf("the %s transaction was executed with error. Code %v", data.Txhash, txData.Code)
	}
	return nil
}

// func to get status of sekaid node
// same as curl localhost:26657/status (port for sekaid's rcp endpoint)
func (s *SekaidManager) GetSekaidStatus(sekaidContainerName, rcpPort string) (*types.Status, error) {
	log := logging.Log

	url := fmt.Sprintf("http://localhost:%s/status", rcpPort)
	log.Println(url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to send GET request:", err)
		return &types.Status{}, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return &types.Status{}, err
	}

	var statusData *types.Status
	if err := json.Unmarshal(body, &statusData); err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return &types.Status{}, err
	}
	return statusData, nil
}
