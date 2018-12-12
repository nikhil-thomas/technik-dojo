package main

import (
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

	listenOn := "localhost:" + args[1]

	s, err := net.ResolveTCPAddr("tcp", listenOn)
	if err != nil {
		log.Fatalln(err)
	}

	l, err := net.ListenTCP("tcp", s)
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	buffer := make([]byte, 1024)
	conn, err := l.Accept()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println(err)
			break
		}
		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("exiting TCP server")
			break
		}
		fmt.Println("-> ", string(buffer[:n-1]))
		n, err = conn.Write(buffer)
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println("echo len:", n)
	}
}
