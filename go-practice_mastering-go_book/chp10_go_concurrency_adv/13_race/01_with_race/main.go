package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"sync"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatalln("one integer argument required")
	}
	n, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup
	var i int

	k := make(map[int]int)
	k[1] = 12

	for i = 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			k[i] = i
		}()
	}

	k[2] = 10
	wg.Wait()
	fmt.Printf("k = %#v\n", k)
}
