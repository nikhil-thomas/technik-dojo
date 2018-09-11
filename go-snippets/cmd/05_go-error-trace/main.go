package main

import (
	"fmt"

	"github.com/pkg/errors"
)

// CustomError defines an error with an error code
type CustomError struct {
	Code int
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("Failed with code %d", c.Code)
}

func main() {
	err := &CustomError{Code: 321}

	wrappedError := errors.Wrap(err, "[1] failed with error")
	doubleWrappedError := errors.Wrap(wrappedError, "[2] failed with error")

	fmt.Println(doubleWrappedError)
	if originalError, ok := errors.Cause(doubleWrappedError).(*CustomError); ok {
		fmt.Printf("%v, %#v\n", originalError, originalError)
		fmt.Println("Original error is:", originalError)
	}
}
