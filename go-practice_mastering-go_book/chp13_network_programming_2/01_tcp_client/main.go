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
		log.Fatalln("host:port required")
	}

	connectTo := args[1]
	c, err := net.Dial("tcp", connectTo)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Fprintf(c, text+"\n")

		message, err := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			break
		}
	}
}
