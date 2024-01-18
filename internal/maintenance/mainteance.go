package maintenance

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/types"
	"github.com/mrlutik/kira2.0/internal/utils"
)

type (
	ValidatorManager struct {
		config            *config.KiraConfig
		helper            *utils.HelperManager
		containerExecutor ContainerExecutor

		log *logging.Logger
	}
	ContainerExecutor interface {
		ExecCommandInContainer(ctx context.Context, containerID string, command []string) ([]byte, error)
	}
)

func NewValidatorManager(helper *utils.HelperManager, containerExecutor ContainerExecutor, config *config.KiraConfig, logger *logging.Logger) *ValidatorManager {
	return &ValidatorManager{
		config:            config,
		helper:            helper,
		containerExecutor: containerExecutor,
		log:               logger,
	}
}

// GetValidatorStatus is getting status of validator from `sekaid query customstaking validator“, returns *ValidatorStatus
func (v *ValidatorManager) GetValidatorStatus(ctx context.Context) (*types.ValidatorStatus, error) {
	v.log.Debugf("Getting validator status")

	kiraAddr, err := v.helper.GetAddressByName(ctx, "validator")
	if err != nil {
		return &types.ValidatorStatus{}, err
	}

	command := fmt.Sprintf("sekaid query customstaking validator --addr %s --output=json", kiraAddr)
	out, err := v.containerExecutor.ExecCommandInContainer(ctx, v.config.SekaidContainerName, []string{"bash", "-c", command})
	if err != nil {
		return &types.ValidatorStatus{}, err
	}

	var data *types.ValidatorStatus
	err = json.Unmarshal(out, &data)
	if err != nil {
		return &types.ValidatorStatus{}, err
	}

	v.log.Debugf("Validator status:\n%+v", data)
	return data, nil
}

// PauseValidator is getting validator status, if status IS ACTIVE running `sekaid tx customslashing pause` inside sekaid container.
// Then checking if transaction of validator pausing was executed inside blockchain, if does -  again checking for validator status
func (v *ValidatorManager) PauseValidator(ctx context.Context, cfg *config.KiraConfig) error {
	v.log.Info("Pausing validator")
	nodeStatus, err := v.GetValidatorStatus(ctx)
	if err != nil {
		return err
	}
	v.log.Debugf("Validator status %s\n", strings.ToLower(nodeStatus.Status))
	if strings.ToLower(nodeStatus.Status) != types.Active {
		v.log.Errorf("Validator status is: %s\n ", strings.ToLower(nodeStatus.Status))
		return fmt.Errorf("cannot pause validator: %w", &MismatchStatusError{
			ExpectedStatus: types.Paused,
			CurrentStatus:  nodeStatus.Status,
		})
	}

	// Command:
	// sekaid tx customslashing unpause
	// --from validator
	// --chain-id testnet-1
	// --keyring-backend=test
	// --home /data/.sekai
	// --fees 100ukex
	// --gas=1000000
	// --broadcast-mode=async --yes
	command := fmt.Sprintf("sekaid tx customslashing pause --from %s --chain-id %s --keyring-backend=test --home  %s --fees 100ukex --broadcast-mode=async --yes --output json", types.ValidatorAccountName, cfg.NetworkName, cfg.SekaidHome)
	v.log.Debugf("Running command\n %s\n", command)
	out, err := v.containerExecutor.ExecCommandInContainer(ctx, cfg.SekaidContainerName, []string{"bash", "-c", command})
	if err != nil {
		v.log.Errorf("Command executing error: %s\n %s\n", command, err)
		return err
	}

	var data types.TxData
	v.log.Debugf("Unmarshaling out: \n%s\n", out)
	err = json.Unmarshal(out, &data)
	if err != nil {
		v.log.Errorf("Unmarshaling [%s]Error: %s", string(out), err)
		return err
	}

	err = v.helper.AwaitNextBlock(ctx, cfg.TimeBetweenBlocks)
	if err != nil {
		return err
	}
	txData, err := v.helper.GetTxQuery(ctx, data.Txhash)
	if err != nil {
		return err
	}
	if txData.Code != 0 {
		v.log.Errorf("Propagating transaction '%s' error\nTransaction status: %d\nData: %+v\n", data.Txhash, txData.Code, txData)
		return &TransactionError{
			TxHash: data.Txhash,
			Code:   txData.Code,
		}
	}

	v.log.Debugf("Getting validator status second time")
	nodeStatus, err = v.GetValidatorStatus(ctx)
	if err != nil {
		return err
	}
	v.log.Debugf("Validator status: %s", strings.ToLower(nodeStatus.Status))
	if strings.ToLower(nodeStatus.Status) != types.Paused {
		return fmt.Errorf("cannot pause validator: %w", &MismatchStatusError{
			ExpectedStatus: types.Paused,
			CurrentStatus:  nodeStatus.Status,
		})
	}
	return nil
}

// UnpauseValidator is getting validator status, if status IS PAUSED running `sekaid tx customslashing unpause` inside sekaid container.
// Then checking if transaction of validator unpausing was executed inside blockchain, if does -  again checking for validator status
func (v *ValidatorManager) UnpauseValidator(ctx context.Context, cfg *config.KiraConfig) error {
	v.log.Info("Unpausing validator")
	nodeStatus, err := v.GetValidatorStatus(ctx)
	if err != nil {
		return err
	}
	if strings.ToLower(nodeStatus.Status) != types.Paused {
		v.log.Errorf("Validator status is:  %s\n ", strings.ToLower(nodeStatus.Status))
		return fmt.Errorf("cannot unpause validator: %w", &MismatchStatusError{
			ExpectedStatus: types.Inactive,
			CurrentStatus:  nodeStatus.Status,
		})
	}

	// Command:
	// sekaid tx customslashing unpause
	// --from validator
	// --chain-id testnet-1
	// --keyring-backend=test
	// --home /data/.sekai
	// --fees 100ukex
	// --gas=1000000
	// --broadcast-mode=async --yes
	command := fmt.Sprintf("sekaid tx customslashing unpause --from %s --chain-id %s --keyring-backend=test --home  %s --fees 100ukex --broadcast-mode=async --yes --log_format=json --output json", types.ValidatorAccountName, cfg.NetworkName, cfg.SekaidHome)
	v.log.Debugf("Running command\n %s\n", command)
	out, err := v.containerExecutor.ExecCommandInContainer(ctx, cfg.SekaidContainerName, []string{"bash", "-c", command})
	if err != nil {
		return err
	}

	var data types.TxData
	v.log.Debugf("Unmarshaling out: \n%s\n", out)
	err = json.Unmarshal(out, &data)
	if err != nil {
		v.log.Errorf("Unmarshaling [%s]Error: %s", string(out), err)
		return err
	}

	err = v.helper.AwaitNextBlock(ctx, cfg.TimeBetweenBlocks)
	if err != nil {
		return err
	}
	txData, err := v.helper.GetTxQuery(ctx, data.Txhash)
	if err != nil {
		return err
	}
	if txData.Code != 0 {
		v.log.Errorf("Propagating transaction '%s' error\nTransaction status: %d\nData: %+v\n", data.Txhash, txData.Code, txData)
		return &TransactionError{
			TxHash: data.Txhash,
			Code:   txData.Code,
		}
	}
	return nil
}

// ActivateValidator is getting validator status, if status IS INACTIVE running `sekaid tx customslashing activate` inside sekaid container.
// Then checking if transaction of validator activating was executed inside blockchain
func (v *ValidatorManager) ActivateValidator(ctx context.Context, cfg *config.KiraConfig) error {
	v.log.Info("Activating validator")
	nodeStatus, err := v.GetValidatorStatus(ctx)
	if err != nil {
		return err
	}
	if strings.ToLower(nodeStatus.Status) != types.Inactive {
		return fmt.Errorf("cannot activate validator: %w", &MismatchStatusError{
			ExpectedStatus: types.Inactive,
			CurrentStatus:  nodeStatus.Status,
		})
	}
	// Command:
	// sekaid tx customslashing activate
	// --from validator
	// --chain-id chaosnet-1
	// --keyring-backend=test
	// --home /data/.sekai/
	// --fees 1000ukex --yes
	// --broadcast-mode=async
	// --log_format=json
	// --output=json
	command := fmt.Sprintf(`sekaid tx customslashing activate --from %s --chain-id %s --keyring-backend=test --home  %s --fees 1000ukex --broadcast-mode=async --yes --output json --log_format=json`, types.ValidatorAccountName, cfg.NetworkName, cfg.SekaidHome)
	v.log.Debugf("Running command\n %s\n", command)
	out, err := v.containerExecutor.ExecCommandInContainer(ctx, cfg.SekaidContainerName, []string{"bash", "-c", command})
	if err != nil {
		return err
	}

	var data types.TxData
	v.log.Debugf("Unmarshaling out: \n%s\n", out)
	err = json.Unmarshal(out, &data)
	if err != nil {
		v.log.Errorf("Unmarshaling [%s]Error: %s", string(out), err)
		return err
	}

	err = v.helper.AwaitNextBlock(ctx, cfg.TimeBetweenBlocks)
	if err != nil {
		return err
	}
	txData, err := v.helper.GetTxQuery(ctx, data.Txhash)
	if err != nil {
		return err
	}
	if txData.Code != 0 {
		v.log.Errorf("Propagating transaction '%s' error\nTransaction status: %d\nData: %+v\n", data.Txhash, txData.Code, txData)
		return fmt.Errorf("pausing validator error: %w", &TransactionError{
			TxHash: data.Txhash,
			Code:   txData.Code,
		})
	}
	return nil
}
