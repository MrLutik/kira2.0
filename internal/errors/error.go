package errors

import "github.com/mrlutik/kira2.0/internal/logging"

var log = logging.Log

// Error handling function
func HandleErr(msg string, err error) {
	if err != nil {
		log.Fatalf("%s, error: %s", msg, err)
	}
}

// TODO add all custom errors here!
