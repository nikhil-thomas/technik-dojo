package log

import (
	"log"

	"github.com/pkg/errors"
)

// OriginalError returns first/original error
func OriginalError() error {
	return errors.New("error occurred")
}

// PassThroughError returns a wrapped error
func PassThroughError() error {
	err := OriginalError()
	return errors.Wrap(err, "in passthrougherror")
}

// Handler gets an error from PassThroughError and prints it
func Handler() {
	err := PassThroughError()
	if err != nil {
		log.Printf("an error occurred: %s\n", err.Error())
	}
}
