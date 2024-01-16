package utils

import "fmt"

type (
	TransactionExecutionError struct {
		TxHash string
		Code   int
	}
	PermissionAddingError struct {
		PermissionToAdd int
		Address         string
		TxHash          string
		Code            int
	}
	TimeoutError struct {
		TimeoutSeconds float64
	}
	ConfigurationVariableNotFoundError struct {
		VariableName string
		Tag          string
	}
	EnvVariableNotFoundError struct {
		VariableName string
	}
)

func (e *TransactionExecutionError) Error() string {
	return fmt.Sprintf("the '%s' transaction was executed with error. Code: %d", e.TxHash, e.Code)
}

func (e *PermissionAddingError) Error() string {
	return fmt.Sprintf("adding '%d' permission to '%s' address error.\nTransaction hash: '%s'.\nCode: '%d'", e.PermissionToAdd, e.Address, e.TxHash, e.Code)
}

func (e *TimeoutError) Error() string {
	return fmt.Sprintf("timeout, failed to await next block within %0.2f s limit", e.TimeoutSeconds)
}

func (e *ConfigurationVariableNotFoundError) Error() string {
	return fmt.Sprintf("the configuration does NOT contain a variable name '%s' occurring after the tag '%s'", e.VariableName, e.Tag)
}

func (e *EnvVariableNotFoundError) Error() string {
	return fmt.Sprintf("env variable '%s' not found", e.VariableName)
}
