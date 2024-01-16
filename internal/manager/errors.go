package manager

import (
	"errors"
	"fmt"
)

type ProcessNotRunningError struct {
	ProcessName   string
	ContainerName string
}

var ErrEmptyNecessaryConfigs = errors.New("cannot apply empty necessary configs for joiner node")

func (e *ProcessNotRunningError) Error() string {
	return fmt.Sprintf("process '%s' is not running in '%s' container", e.ProcessName, e.ContainerName)
}
