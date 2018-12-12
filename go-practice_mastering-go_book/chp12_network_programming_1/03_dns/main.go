package main

import (
	"fmt"
	"net"
	"os"
)

func lookUpIP(address string) ([]string, error) {
	hosts, err := net.LookupAddr(address)
	if err != nil {
		return nil, err
	}
	return hosts, nil
}

func lookUpHostName(hostname string) ([]string, error) {
	ips, err := net.LookupHost(hostname)
	if err != nil {
		return nil, err
	}
	return ips, nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("provide an argument")
		os.Exit(1)
	}

	input := args[1]
	IPAddress := net.ParseIP(input)

	if IPAddress == nil {
		IPs, err := lookUpHostName(input)
		if err != nil {
			fmt.Println(err)
		} else {
			for _, ip := range IPs {
				fmt.Println(ip)
			}
		}
	} else {
		hosts, err := lookUpIP(input)
		if err != nil {
			fmt.Println(err)
		} else {
			for _, host := range hosts {
				fmt.Println(host)
			}
		}
	}
}
