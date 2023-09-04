package errors

import (
	"fmt"

	"github.com/mrlutik/kira2.0/internal/logging"
)

var log = logging.Log

// HandleFatalErr handles fatal errors from functions
func HandleFatalErr(msg string, err error) {
	if err != nil {
		log.Fatalf("%s, error: %s", msg, err)
	}
}

func LogAndReturnErr(msg string, err error) error {
	if err != nil {
		log.Errorf("%s, error: %s", msg, err)
		return fmt.Errorf("%s, error: %w", msg, err)
	}

	return nil
}

// TODO add all custom errors here!
