package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("arg url required")
	}

	url, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		log.Fatalln(err)
	}

	httpData, err := c.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("status code:", httpData.Status)
	header, err := httputil.DumpResponse(httpData, false)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(header))

	contentType := httpData.Header.Get("Content-Type")
	fmt.Println(contentType)
	charSet := strings.SplitAfter(contentType, "charset=")
	if len(charSet) > 1 {
		fmt.Println("Character Set:", charSet[1])
	}

	if httpData.ContentLength == -1 {
		fmt.Println("content length unknown")
	} else {
		fmt.Println("content length:", httpData.ContentLength)
	}

	length := 0
	buffer := make([]byte, 1024)
	r := httpData.Body
	for {
		n, err := r.Read(buffer[0:])
		if err != nil {
			fmt.Println(err)
			break
		}
		length = length + n
	}
	fmt.Println("calculated respose data length:", length)

}
