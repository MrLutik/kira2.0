package monitoring

import (
	"errors"
	"fmt"
)

type (
	HTTPRequestFailedError struct {
		StatusCode int
	}

	ValidatorAddressNotFoundError struct {
		Address string
	}
)

var (
	ErrExtractingPublicIP     = errors.New("unable to extract public IP address")
	ErrGettingPublicIPAddress = errors.New("can't get the public IP address")
)

func (e *HTTPRequestFailedError) Error() string {
	return fmt.Sprintf("HTTP request failed with status: %d", e.StatusCode)
}

func (e *ValidatorAddressNotFoundError) Error() string {
	return fmt.Sprintf("can't find validator address '%s'", e.Address)
}
