package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptrace"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("specify a url")
	}
	url := os.Args[1]
	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			fmt.Println("first response byte!")
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("got conn: %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("dns info: %+v\n", dnsInfo)
		},
		ConnectStart: func(network, addr string) {
			fmt.Println("dial start")
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Println("dial done")
		},
		WroteHeaders: func() {
			fmt.Println("wrote headers")
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	fmt.Println("requesting data from server")
	_, err = http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatalln(err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("\nResponse\n")
	io.Copy(os.Stdout, response.Body)
}
