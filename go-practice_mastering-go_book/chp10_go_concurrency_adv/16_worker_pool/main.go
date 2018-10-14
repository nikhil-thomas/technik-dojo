package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

var workerid = 0

type client struct {
	id  int
	val int
}

type data struct {
	job    client
	result int
}

var (
	size       = 10
	clientChan = make(chan client, size)
	dataChan   = make(chan data, size)
)

func worker(w *sync.WaitGroup) {
	defer w.Done()
	workerid++
	wid := workerid
	for c := range clientChan {
		fmt.Println("worker:", wid, "processing", c.val)
		square := c.val * c.val
		output := data{c, square}
		dataChan <- output
		time.Sleep(1 * time.Second)
	}
}

func makeWorkerPool(n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(dataChan)
}

func create(n int) {
	for i := 1; i <= n; i++ {
		c := client{i, i}
		clientChan <- c
	}
	close(clientChan)
}

func main() {
	flag.Parse()

	if len(flag.Args()) != 2 {
		log.Fatalln("num jobs and num workers required")
	}

	fmt.Println("capacity of client channel", cap(clientChan))
	fmt.Println("capacity of data channel", cap(dataChan))

	nJobs, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}

	nWorkers, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatalln(err)
	}

	go create(nJobs)
	finishedChan := make(chan interface{})

	go func() {
		for d := range dataChan {
			fmt.Printf("client id: %d\tval: %d\n", d.job.id, d.job.val)
			fmt.Printf("val: %d\tsquare: %d\n", d.job.val, d.result)
		}
		finishedChan <- true
	}()

	makeWorkerPool(nWorkers)
	fmt.Printf(": %v\n", <-finishedChan)
}
