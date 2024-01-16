package monitoring

import "fmt"

type (
	HTTPRequestFailedError struct {
		StatusCode int
	}

	ValidatorAddressNotFoundError struct {
		Address string
	}
)

func (e *HTTPRequestFailedError) Error() string {
	return fmt.Sprintf("HTTP request failed with status: %d", e.StatusCode)
}

func (e *ValidatorAddressNotFoundError) Error() string {
	return fmt.Sprintf("can't find validator address '%s'", e.Address)
}
