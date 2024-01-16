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
