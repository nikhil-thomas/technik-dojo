package errhndl

import (
	"fmt"
	"os"
)

// ExitErrorf formats error and print to Stderr
func ExitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
