package manager

import (
	"errors"
	"fmt"
)

type (
	ProcessNotRunningError struct {
		ProcessName   string
		ContainerName string
	}
	StringPrefixError struct {
		StringValue string
		Prefix      string
	}
)

var (
	ErrEmptyNecessaryConfigs    = errors.New("cannot apply empty necessary configs for joiner node")
	ErrSHA256ChecksumMismatch   = errors.New("sha256 checksum is not the same")
	ErrFilesContentNotIdentical = errors.New("files content are not identical")
	ErrInvalidIPPortFormat      = errors.New("invalid IP and Port format in seed")
	ErrInvalidSeedFormat        = errors.New("invalid seed format")
)

func (e *ProcessNotRunningError) Error() string {
	return fmt.Sprintf("process '%s' is not running in '%s' container", e.ProcessName, e.ContainerName)
}

func (e *StringPrefixError) Error() string {
	return fmt.Sprintf("string '%s' does not have prefix '%s'", e.StringValue, e.Prefix)
}
