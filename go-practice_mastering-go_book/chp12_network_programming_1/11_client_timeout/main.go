package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

var timeout = time.Duration(1 * time.Second)

func timeoutConn(network, host string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, host, timeout)
	if err != nil {
		return nil, err
	}
	conn.SetDeadline(time.Now().Add(timeout))
	return conn, nil
}

func main() {
	if len(os.Args) == 1 {
		log.Fatalln("specify url and timeout")
	}
	if len(os.Args) == 3 {
		temp, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("using default timeout:", timeout)
		} else {
			timeout = time.Duration(time.Duration(temp) * time.Second)
		}
	}

	url := os.Args[1]
	t := http.Transport{
		Dial: timeoutConn,
	}

	client := http.Client{
		Transport: &t,
	}

	data, err := client.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer data.Body.Close()
	_, err = io.Copy(os.Stdout, data.Body)
	if err != nil {
		log.Fatalln(err)
	}
}
