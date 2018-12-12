package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func fib(n int) int {
	fMap := make(map[int]int)
	for i := 0; i <= n; i++ {
		var val int
		if i <= 2 {
			val = n
		} else {
			val = fib(i-1) + fib(i-2)
		}
		fMap[i] = val
	}
	return fMap[n]
}

func handleConn(c net.Conn) {
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

		val := "-1\n"
		n, err := strconv.Atoi(temp)
		if err != nil {
			log.Fatalln(err)
		}
		val = strconv.Itoa(fib(n)) + "\n"
		fmt.Println(n, ":", val)
		c.Write([]byte(val))
	}
	time.Sleep(5 * time.Second)
	c.Close()
}

func main() {
	args := os.Args
	if len(args) == 1 {
		log.Fatalln("provide port number")
	}

	listenOn := ":" + args[1]

	l, err := net.Listen("tcp4", listenOn)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			break
		}
		go handleConn(conn)
	}
}
