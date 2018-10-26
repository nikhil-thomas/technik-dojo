package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		log.Fatalln("port required")
	}

	connectTo := args[1]

	s, err := net.ResolveUDPAddr("udp4", connectTo)
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	fmt.Println("UDP server:", conn.RemoteAddr().String())

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		data := []byte(text + "\n")
		_, err = conn.Write([]byte(data))
		if err != nil {
			log.Fatalln(err)
		}
		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("exiting UDP client")
			break
		}

		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Printf("reply: %s\n", string(buffer[:n]))
	}
}
