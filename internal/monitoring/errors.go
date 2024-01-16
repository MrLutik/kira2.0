package monitoring

import "fmt"

type HTTPRequestFailedError struct {
	StatusCode int
}

func (e *HTTPRequestFailedError) Error() string {
	return fmt.Sprintf("HTTP request failed with status: %d", e.StatusCode)
}
