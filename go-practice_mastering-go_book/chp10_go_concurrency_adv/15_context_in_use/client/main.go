package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	url   string
	delay = 5
	wg    sync.WaitGroup
)

type data struct {
	r   *http.Response
	err error
}

func connect(c context.Context) error {
	defer wg.Done()
	dataChan := make(chan data, 1)

	tr := &http.Transport{}
	httpClient := &http.Client{Transport: tr}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	go func() {
		response, err := httpClient.Do(req)
		var pack data
		if err != nil {
			fmt.Println(err)
			pack = data{nil, err}
		}
		pack = data{response, nil}
		dataChan <- pack
	}()

	select {
	case <-c.Done():
		tr.CancelRequest(req)
		d := <-dataChan
		fmt.Println("request was canceled, data:", d)
		return c.Err()
	case ok := <-dataChan:
		err := ok.err
		resp := ok.r

		if err != nil {
			fmt.Println("Error select:", err)
			return err
		}
		defer resp.Body.Close()

		httpData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error select:", err)
			return err
		}

		fmt.Printf("Server response: %s\n", httpData)
	}
	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 2 {
		log.Fatalln("specify url and delay")
	}
	t, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatalln(err)
	}
	delay = t
	url = flag.Arg(0)

	fmt.Println("url:", url)
	fmt.Println("delay:", delay)

	c := context.Background()
	c, cancel := context.WithTimeout(c, time.Duration(delay)*time.Second)
	defer cancel()

	wg.Add(1)
	go connect(c)
	wg.Wait()
	fmt.Println("end main")
}
