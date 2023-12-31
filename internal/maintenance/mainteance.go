package maintenance

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/manager/utils"
	"github.com/mrlutik/kira2.0/internal/types"
)

var log = logging.Log

// GetValidatorStatus is geting status of validator from `sekaid query customstaking validator“, returns *ValidatorStatus
func GetValidatorStatus(ctx context.Context, cfg *config.KiraConfig, cm *docker.ContainerManager) (*types.ValidatorStatus, error) {
	h := utils.NewHelperManager(cm, cfg)
	kiraAddr, err := h.GetAddressByName(ctx, "validator")
	if err != nil {
		return &types.ValidatorStatus{}, err
	}
	command := fmt.Sprintf("sekaid query customstaking validator --addr %s --output=json", kiraAddr)
	out, err := cm.ExecCommandInContainer(ctx, cfg.SekaidContainerName, []string{"bash", "-c", command})
	if err != nil {
		return &types.ValidatorStatus{}, err
	}
	var data *types.ValidatorStatus
	err = json.Unmarshal(out, &data)
	if err != nil {
		return &types.ValidatorStatus{}, err
	}
	log.Debugf("Validator status:\n%+v\n", data)
	return data, nil
}

// PauseValidator is geting validator status, if status IS ACTIVE running `sekaid tx customslashing pause` inside sekaid container.
// Then checking if transaction of validator pausing was executed inside blockchain, if does -  again checking for validator status
func PauseValidator(ctx context.Context, cfg *config.KiraConfig, cm *docker.ContainerManager) error {
	log.Debugf("Pausing validator\n")
	log.Debugf("Geting validator status\n")
	nodeStatus, err := GetValidatorStatus(ctx, cfg, cm)
	if err != nil {
		return err
	}
	log.Debugf("VALIDATOR STATUS %s\n", strings.ToLower(nodeStatus.Status))
	if strings.ToLower(nodeStatus.Status) != types.Active {
		log.Errorf("Validator status is:  %s\n ", strings.ToLower(nodeStatus.Status))
		return fmt.Errorf("cannot pause validator, node status is not <%s>, curent status <%s>", types.Active, nodeStatus.Status)
	}
	// os.Exit(1)
	//sekaid tx customslashing unpause --from validator --chain-id testnet-1 --keyring-backend=test --home  /data/.sekai --fees 100ukex --gas=1000000 --broadcast-mode=async --yes
	command := fmt.Sprintf("sekaid tx customslashing pause --from %s --chain-id %s --keyring-backend=test --home  %s --fees 100ukex --gas=1000000 --broadcast-mode=async --yes --output json", types.ValidatorAccountName, cfg.NetworkName, cfg.SekaidHome)
	log.Debugf("Running command\n %s\n", command)
	out, err := cm.ExecCommandInContainer(ctx, cfg.SekaidContainerName, []string{"bash", "-c", command})
	if err != nil {
		log.Errorf("Command executing error: %s\n %s\n", command, err)
		return err
	}
	var data types.TxData
	log.Debugf("Unmarshaling out: \n%s\n", out)
	err = json.Unmarshal(out, &data)
	if err != nil {
		log.Errorf("Unmarshaling [%s]Error: %s", string(out), err)
		return err
	}
	hm := utils.NewHelperManager(cm, cfg)
	err = hm.AwaitNextBlock(ctx, cfg.TimeBetweenBlocks)
	if err != nil {
		return err
	}
	txData, err := hm.GetTxQuery(ctx, data.Txhash)
	if err != nil {
		return err
	}
	if txData.Code != 0 {
		log.Errorf("Propagating transaction '%s' error. Transaction status: %d\n%+v\n", data.Txhash, txData.Code, txData)
		return fmt.Errorf("pausing validator error\nTransaction hash: '%s'.\nCode: '%d'", data.Txhash, txData.Code)
	}

	log.Debugf("Geting validator status second time\n")
	nodeStatus, err = GetValidatorStatus(ctx, cfg, cm)
	if err != nil {
		return err
	}
	log.Debugf("VALIDATOR STATUS %s\n", strings.ToLower(nodeStatus.Status))
	if strings.ToLower(nodeStatus.Status) != types.Paused {
		return fmt.Errorf("cannot pause validator, node status is not <%s>, curent status <%s>", types.Active, nodeStatus.Status)
	}
	return nil
}

// UnpauseValidator is geting validator status, if status IS PAUSED running `sekaid tx customslashing unpause` inside sekaid container.
// Then checking if transaction of validator unpausing was executed inside blockchain, if does -  again checking for validator status
func UnpauseValidator(ctx context.Context, cfg *config.KiraConfig, cm *docker.ContainerManager) error {
	log.Debugf("Unpausing validator\n")
	log.Debugf("Geting validator status\n")
	nodeStatus, err := GetValidatorStatus(ctx, cfg, cm)
	if err != nil {
		return err
	}
	if strings.ToLower(nodeStatus.Status) != types.Paused {
		log.Errorf("Validator status is:  %s\n ", strings.ToLower(nodeStatus.Status))
		return fmt.Errorf("cannot unpause validator, node status is not <%s>, curent status <%s>", types.Active, nodeStatus.Status)
	}
	//sekaid tx customslashing unpause --from validator --chain-id testnet-1 --keyring-backend=test --home  /data/.sekai --fees 100ukex --gas=1000000 --broadcast-mode=async --yes
	command := fmt.Sprintf("sekaid tx customslashing unpause --from %s --chain-id %s --keyring-backend=test --home  %s --fees 100ukex --gas=1000000 --broadcast-mode=async --yes --output json", types.ValidatorAccountName, cfg.NetworkName, cfg.SekaidHome)
	log.Debugf("Running command\n %s\n", command)
	out, err := cm.ExecCommandInContainer(ctx, cfg.SekaidContainerName, []string{"bash", "-c", command})
	if err != nil {
		return err
	}
	var data types.TxData
	log.Debugf("Unmarshaling out: \n%s\n", out)
	err = json.Unmarshal(out, &data)
	if err != nil {
		log.Errorf("Unmarshaling [%s]Error: %s", string(out), err)
		return err
	}
	hm := utils.NewHelperManager(cm, cfg)
	err = hm.AwaitNextBlock(ctx, cfg.TimeBetweenBlocks)
	if err != nil {
		return err
	}
	txData, err := hm.GetTxQuery(ctx, data.Txhash)
	if err != nil {
		return err
	}
	if txData.Code != 0 {
		log.Errorf("Propagating transaction '%s' error. Transaction status: %d\n%+v\n", data.Txhash, txData.Code, txData)
		return fmt.Errorf("unpausing validator error\nTransaction hash: '%s'.\nCode: '%d'", data.Txhash, txData.Code)
	}
	if strings.ToLower(nodeStatus.Status) != types.Paused {
		log.Errorf("Validator status is:  %s\n ", strings.ToLower(nodeStatus.Status))
		return fmt.Errorf("cannot unpause validator, node status is not <%s>, curent status <%s>", types.Active, nodeStatus.Status)
	}
	return nil
}

// ActivateValidator is geting validator status, if status IS INACTIVE running `sekaid tx customslashing unpause` inside sekaid container.
// Then checking if transaction of validator unpausing was executed inside blockchain
func ActivateValidator(ctx context.Context, cfg *config.KiraConfig, cm *docker.ContainerManager) error {
	log.Debugf("Activating validator\n")
	log.Debugf("Geting validator status\n")
	nodeStatus, err := GetValidatorStatus(ctx, cfg, cm)
	if err != nil {
		return err
	}
	if strings.ToLower(nodeStatus.Status) != types.Inactive {
		return fmt.Errorf("cannot activate validator, node status is not <%s>, curent status <%s>", types.Inactive, nodeStatus.Status)
	}
	//sekaid tx customslashing unpause --from validator --chain-id testnet-1 --keyring-backend=test --home  /data/.sekai --fees 100ukex --gas=1000000 --broadcast-mode=async --yes
	command := fmt.Sprintf("sekaid tx customslashing activate --from %s --chain-id %s --keyring-backend=test --home  %s --fees 100ukex --gas=1000000 --broadcast-mode=async --yes --output json", types.ValidatorAccountName, cfg.NetworkName, cfg.SekaidHome)
	log.Debugf("Running command\n %s\n", command)
	out, err := cm.ExecCommandInContainer(ctx, cfg.SekaidContainerName, []string{"bash", "-c", command})
	if err != nil {
		return err
	}
	var data types.TxData
	log.Debugf("Unmarshaling out: \n%s\n", out)
	err = json.Unmarshal(out, &data)
	if err != nil {
		log.Errorf("Unmarshaling [%s]Error: %s", string(out), err)
		return err
	}
	hm := utils.NewHelperManager(cm, cfg)
	err = hm.AwaitNextBlock(ctx, cfg.TimeBetweenBlocks)
	if err != nil {
		return err
	}
	txData, err := hm.GetTxQuery(ctx, data.Txhash)
	if err != nil {
		return err
	}
	if txData.Code != 0 {
		log.Errorf("Propagating transaction '%s' error. Transaction status: %d\n%+v\n", data.Txhash, txData.Code, txData)
		return fmt.Errorf("pausing validator error\nTransaction hash: '%s'.\nCode: '%d'", data.Txhash, txData.Code)
	}
	return nil
}
