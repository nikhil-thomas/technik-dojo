package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()

	if err != nil {
		fmt.Print(err)
		return
	}

	for _, intf := range interfaces {
		fmt.Printf("Name: %v\n", intf.Name)
		fmt.Println("Interface flags:", intf.Flags.String())
		fmt.Println("Interface MTU:", intf.MTU)
		fmt.Println("Interface Hardware Address:", intf.HardwareAddr)
		fmt.Print("\n")
	}
}
