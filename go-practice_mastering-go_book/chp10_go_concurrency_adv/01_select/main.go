package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

func genRand(done <-chan bool, min, max int) <-chan int {
	intChan := make(chan int)
	go func() {
		timeoutChan := time.After(4 * time.Second)
		for {
			select {
			case intChan <- rand.Intn(max-min) + min:
			case <-done:
				return
			case <-timeoutChan:
				fmt.Println("timeout")
			}
		}
	}()
	return intChan
}

func main() {
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	if len(os.Args) != 2 {
		log.Fatalln("upper limt for random not specified")
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Starting to create %d random numbers\n", n)

	rand.Seed(time.Now().Unix())
	done := make(chan bool)

	nChan := genRand(done, 0, 2*n)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", <-nChan)
		time.Sleep(1 * time.Second)
	}
	fmt.Print("\nafter forloop\n")

	time.Sleep(5 * time.Second)

	done <- true
	fmt.Println("Exiting")

}
