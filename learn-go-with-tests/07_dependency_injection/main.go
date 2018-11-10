package main

import (
	"fmt"
	"io"
	"os"
)

// Greet write hello message to io.writer
func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}
func main() {
	Greet(os.Stdout, "James")
}
