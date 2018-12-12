package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument"
	} else {
		myString = arguments[1]
	}
	io.WriteString(os.Stdout, "Stdout: This is Stdout\n")
	io.WriteString(os.Stderr, "Stderr: "+myString)
	io.WriteString(os.Stderr, "\n")
}
