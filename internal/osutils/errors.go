package osutils

import "fmt"

type (
	InvalidPortRangeError struct {
		Port int
	}

	InvalidIPError struct {
		Input string
	}
)

func (e *InvalidPortRangeError) Error() string {
	return fmt.Sprintf("'%d' port is not in a valid range", e.Port)
}

func (e *InvalidIPError) Error() string {
	return fmt.Sprintf("'%s' is not a valid IP", e.Input)
}
