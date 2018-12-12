package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	args := os.Args
	if len(args) == 1 {
		log.Fatalln("port required")
	}

	listenOn := args[1]

	addr, err := net.ResolveUDPAddr("udp4", listenOn)
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Print("-> ", string(buffer[:n]))

		if strings.TrimSpace(string(buffer[:n])) == "STOP" {
			fmt.Println("exiting UDP server")
			break
		}

		data := []byte(strconv.Itoa(random(1, 1001)))
		fmt.Printf("data: %s\n", string(data))
		_, err = conn.WriteToUDP(data, addr)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
