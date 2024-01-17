package maintenance

import "fmt"

type (
	TransactionError struct {
		TxHash string
		Code   int
	}
	MismatchStatusError struct {
		ExpectedStatus string
		CurrentStatus  string
	}
)

func (e *TransactionError) Error() string {
	return fmt.Sprintf("transaction error\nHash: '%s'\nCode: '%d'", e.TxHash, e.Code)
}

func (e *MismatchStatusError) Error() string {
	return fmt.Sprintf("node status is not '%s', current status is '%s'", e.ExpectedStatus, e.CurrentStatus)
}
