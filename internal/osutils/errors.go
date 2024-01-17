package osutils

import "fmt"

type InvalidIPError struct {
	Input string
}

func (e *InvalidIPError) Error() string {
	return fmt.Sprintf("'%s' is not a valid IP", e.Input)
}
