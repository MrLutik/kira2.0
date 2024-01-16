package systemd

import (
	"errors"
	"fmt"
)

type (
	DifferentPropertyValuesError struct {
		PropertyName  string
		ExpectedValue string
		ActualValue   string
	}
)

var (
	ErrServiceTimeout                    = errors.New("timeout reached while working with service")
	ErrServiceRestartCancellationTimeout = errors.New("restarting service was cancelled or timed out")
)

func (e *DifferentPropertyValuesError) Error() string {
	return fmt.Sprintf("different values for the property '%s': expected '%s', got '%s'", e.PropertyName, e.ExpectedValue, e.ActualValue)
}
