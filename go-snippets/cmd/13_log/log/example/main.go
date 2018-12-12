package main

import (
	"fmt"

	"github.com/nikhil-thomas/technik-dojo/go-snippets/cmd/13_log/log"
)

func main() {
	fmt.Println("basic logging and modification of logger")
	log.Log()
	fmt.Println("logging handled errors")
	log.Handler()
}
