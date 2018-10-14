package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

var (
	m  sync.Mutex
	v1 int
)

func changeInt(i int) {
	m.Lock()
	time.Sleep(500 * time.Millisecond)
	v1++
	if v1%10 == 0 {
		v1 -= 10 * i
	}
	m.Unlock()
}

func read() int {
	m.Lock()
	a := v1
	m.Unlock()
	return a
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatalln("integer argument required")
	}

	numGR, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup

	fmt.Printf("%d ", read())
	for i := 0; i < numGR; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			changeInt(i)
			fmt.Printf("-> %d ", read())
		}(i)
	}

	wg.Wait()
	fmt.Printf("-> %d\n", read())
}
