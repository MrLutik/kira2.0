package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/types"
	"sigs.k8s.io/yaml"
)

type HelperManager struct {
	config           *config.KiraConfig
	containerManager *docker.ContainerManager
}

func NewHelperManager(containerManager *docker.ContainerManager, config *config.KiraConfig) *HelperManager {
	return &HelperManager{containerManager: containerManager, config: config}
}

// Getting TX by parsing json output of `sekaid query tx <TXhash>`
func (h *HelperManager) GetTxQuery(ctx context.Context, transactionHash string) (types.TxData, error) {
	log := logging.Log
	var data types.TxData

	command := fmt.Sprintf(`sekaid query tx %s -output=json`, transactionHash)
	out, err := h.containerManager.ExecCommandInContainer(ctx, h.config.SekaidContainerName, []string{`bash`, `-c`, command})
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

// AwaitNextBlock waits for the next block to be propagated and reached by the sekaid node.
// It continuously checks the current block height and compares it with the initial height.
// The method waits for a specified timeout duration for the next block to be reached,
// and returns an error if the timeout is exceeded.
// If the next block is reached within the timeout, the method returns nil.
func (h *HelperManager) AwaitNextBlock(ctx context.Context, timeout time.Duration) error {
	// Adding 1 more second because in real case scenario block propagation takes a few seconds\milliseconds longer
	timeout += time.Second * 5
	log := logging.Log

	log.Infof("Checking current block height")
	currentBlockHeight, err := h.getBlockHeight(ctx)
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
			if elapsed > timeout-1 {
				log.Errorf("Awaiting next block reached timeout: %0.0f seconds", timeout.Seconds())
				return &TimeoutError{TimeoutSeconds: timeout.Seconds()}
			}

			blockHeight, err := h.getBlockHeight(ctx)
			if err != nil {
				return fmt.Errorf("getting next block height error: %w", err)
			}

			if blockHeight == currentBlockHeight {
				log.Warningf("WAITING: Block is NOT propagated yet: elapsed %0.0f / %0.0f seconds", elapsed.Seconds(), timeout.Seconds())
				continue
			}

			// Exit awaiting block
			log.Infof("Next block '%s' reached...", blockHeight)
			return nil

		case <-ctx.Done():
			return fmt.Errorf("awaiting context timeout error: %w", ctx.Err())
		}
	}
}

// getBlockHeight retrieves the latest block height from the sekaid node.
// It executes the "sekaid status" command in the specified container
// and parses the JSON output to extract the block height.
// If successful, it returns the latest block height as a string.
// Otherwise, it returns an error.
func (h *HelperManager) getBlockHeight(ctx context.Context) (string, error) {
	log := logging.Log

	cmd := "sekaid status"
	statusOutput, err := h.containerManager.ExecCommandInContainer(ctx, h.config.SekaidContainerName, []string{"bash", "-c", cmd})
	if err != nil {
		return "", fmt.Errorf("getting '%s' error: %w", cmd, err)
	}

	var status struct { // Anonymous structure which represents the partial response from `sekaid status`
		SyncInfo struct {
			LatestBlockHeight string `json:"latest_block_height"`
		} `json:"SyncInfo"`
	}
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
func (h *HelperManager) GivePermissionToAddress(ctx context.Context, permissionToAdd int, address string) error {
	log := logging.Log
	command := fmt.Sprintf(`sekaid tx customgov permission whitelist --from %s --keyring-backend=test --permission=%v --addr=%s --chain-id=%s --home=%s --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json`, address, permissionToAdd, address, h.config.NetworkName, h.config.SekaidHome)
	out, err := h.containerManager.ExecCommandInContainer(ctx, h.config.SekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Giving '%d' permission error. Command: '%s'. Error: %s", permissionToAdd, command, err)
		return err
	}
	log.Infof("Permission '%d' is pushed to network for address '%s'", permissionToAdd, address)

	var data types.TxData
	err = json.Unmarshal(out, &data)
	if err != nil {
		log.Errorf("Unmarshaling [%s]Error: %s", string(out), err)
		return err
	}
	log.Debugf("Give permission to address output: Hash: '%s'.Code: %d", data.Txhash, data.Code)

	err = h.AwaitNextBlock(ctx, h.config.TimeBetweenBlocks)
	if err != nil {
		log.Errorf("Awaiting error: %s", err)
		return fmt.Errorf("awaiting error: %w", err)
	}

	txData, err := h.GetTxQuery(ctx, data.Txhash)
	if err != nil {
		log.Errorf("Getting transaction query error: %s", err)
		return fmt.Errorf("getting tx query error: %w", err)
	}

	if txData.Code != 0 {
		log.Errorf("Propagating transaction '%s' error. Transaction status: %d", data.Txhash, txData.Code)
		return &PermissionAddingError{
			PermissionToAdd: permissionToAdd,
			Address:         address,
			TxHash:          data.Txhash,
			Code:            txData.Code,
		}
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
func (h *HelperManager) CheckAccountPermission(ctx context.Context, permissionToCheck int, address string) (bool, error) {
	log := logging.Log

	command := fmt.Sprintf("sekaid query customgov permissions %s --output=json", address)
	out, err := h.containerManager.ExecCommandInContainer(ctx, h.config.SekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Executing '%s' command in '%s' container error: %s", command, h.config.SekaidContainerName, err)
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
			log.Infof("Permission '%d' was found with '%s' address", permissionToCheck, address)

			return true, nil
		}
	}

	// TODO Warning or Error?
	log.Errorf("Permission '%d' wasn't found with '%s' address", permissionToCheck, address)
	return false, nil
}

// Getting address from keyring.
// Command: sekaid keys show validator --keyring-backend=test --home=test
func (h *HelperManager) GetAddressByName(ctx context.Context, addressName string) (string, error) {
	log := logging.Log
	command := fmt.Sprintf("sekaid keys show %s --keyring-backend=%s --home=%s", addressName, h.config.KeyringBackend, h.config.SekaidHome)
	out, err := h.containerManager.ExecCommandInContainer(ctx, h.config.SekaidContainerName, []string{`bash`, `-c`, command})
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

	log.Infof("Validator address: '%s'", key[0].Address)
	return key[0].Address, nil
}

// Updating identity registrar from KM1 await-validator-init.sh file.
func (h *HelperManager) UpdateIdentityRegistrarFromValidator(ctx context.Context, accountName string) error {
	log := logging.Log

	nodeStruct, err := h.getSekaidStatus()
	if err != nil {
		log.Errorf("Getting sekaid status error: %s", err)
		return err
	}

	address, err := h.GetAddressByName(ctx, accountName)
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
		err = h.UpsertIdentityRecord(ctx, address, accountName, record.key, record.value)
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
func (h *HelperManager) UpsertIdentityRecord(ctx context.Context, address, account, key, value string) error {
	var (
		log = logging.Log
		err error
		out []byte
	)

	if value != "" {
		log.Infof("Registering identity record from address '%s': {'%s': '%s'}", address, key, value)
		command := fmt.Sprintf(`sekaid tx customgov register-identity-records --infos-json="{\"%s\":\"%s\"}" --from=%s --keyring-backend=%s --home=%s --chain-id=%s --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json`, key, value, address, h.config.KeyringBackend, h.config.SekaidHome, h.config.NetworkName)
		out, err = h.containerManager.ExecCommandInContainer(ctx, h.config.SekaidContainerName, []string{"bash", "-c", command})
		if err != nil {
			log.Errorf("Executing command '%s' in '%s' container error: %s", command, h.config.SekaidContainerName, err)
			return err
		}

	} else {
		log.Infof("Deleting identity record from address '%s': key %s", address, key)
		command := fmt.Sprintf(`sekaid tx customgov delete-identity-records --keys="%s" --from=%s --keyring-backend=%s --home=%s --chain-id=%s --fees=100ukex --yes --broadcast-mode=async --log_format=json --output=json`, key, address, h.config.KeyringBackend, h.config.SekaidHome, h.config.NetworkName)
		out, err = h.containerManager.ExecCommandInContainer(ctx, h.config.SekaidContainerName, []string{"bash", "-c", command})
		if err != nil {
			log.Errorf("Executing command '%s' in '%s' container error: %s", command, h.config.SekaidContainerName, err)
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

	err = h.AwaitNextBlock(ctx, h.config.TimeBetweenBlocks)
	if err != nil {
		log.Errorf("Awaiting error: %s", err)
		return fmt.Errorf("awaiting error: %w", err)
	}

	txData, err := h.GetTxQuery(ctx, data.Txhash)
	if err != nil {
		log.Errorf("Getting transaction query error: %s", err)
		return fmt.Errorf("getting tx query error: %w", err)
	}

	if txData.Code != 0 {
		log.Errorf("The '%s' transaction was executed with error. Code: %d", data.Txhash, txData.Code)
		return &TransactionExecutionError{
			TxHash: data.Txhash,
			Code:   txData.Code,
		}
	}

	return nil
}

// func to get status of sekaid node
// same as curl localhost:26657/status (port for sekaid's rpc endpoint)
func (h *HelperManager) getSekaidStatus() (*types.Status, error) {
	log := logging.Log

	url := fmt.Sprintf("http://localhost:%s/status", h.config.RpcPort)
	log.Infof("URL: %s", url)
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
