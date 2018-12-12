package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("provide an argument")
		os.Exit(1)
	}

	domain := args[1]
	MXs, err := net.LookupMX(domain)
	if err != nil {
		log.Fatalln(err)
	}

	for _, mx := range MXs {
		fmt.Println(mx.Host)
	}
}
