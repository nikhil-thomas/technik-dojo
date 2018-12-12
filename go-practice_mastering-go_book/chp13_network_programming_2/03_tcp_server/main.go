package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		log.Fatalln("port required")
	}

	listenOn := ":" + args[1]

	l, err := net.Listen("tcp", listenOn)
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("exiting TCP server")
			break
		}
		fmt.Print("-> ", string(netData))
		t := time.Now().Format(time.RFC3339) + "\n"
		c.Write([]byte(t))
	}
}
