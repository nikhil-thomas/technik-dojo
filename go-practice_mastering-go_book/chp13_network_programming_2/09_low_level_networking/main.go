package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	netaddr, err := net.ResolveIPAddr("ip4", "127.0.0.1")
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := net.ListenIP("ip4:icmp", netaddr)
	if err != nil {
		log.Fatalln(err)
	}

	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFrom(buffer)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("% X\n", buffer[:n])
}
