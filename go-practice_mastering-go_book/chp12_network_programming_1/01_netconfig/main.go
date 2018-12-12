package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, intf := range interfaces {
		fmt.Printf("interface: %v\n", intf.Name)
		byname, err := net.InterfaceByName(intf.Name)
		if err != nil {
			fmt.Println(err)
		}
		addresses, err := byname.Addrs()
		for k, v := range addresses {
			fmt.Printf("interface address #%v: %v\n", k, v.String())
		}
		fmt.Print("\n")
	}
}
