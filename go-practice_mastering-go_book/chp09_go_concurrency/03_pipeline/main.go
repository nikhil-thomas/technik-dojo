package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var closeA bool
var data = make(map[int]bool)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func stage1(min, max int) <-chan int {
	randChan := make(chan int)
	go func() {
		for {
			if closeA {
				close(randChan)
				return
			}
			randChan <- random(min, max)
		}
	}()
	return randChan
}

func stage2(in <-chan int) <-chan int {
	intChan := make(chan int)
	go func() {
		for v := range in {
			fmt.Print(v, " ")
			_, ok := data[v]
			if ok {
				closeA = true
				close(intChan)
				return
			}
			data[v] = true
			intChan <- v
		}
	}()
	return intChan
}

func stage3(in <-chan int) {
	sum := 0
	for v := range in {
		sum += v
	}
	fmt.Println("sum of randon numbers:", sum)
}

func main() {
	flag.Parse()
	min, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	max, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}
	if min >= max {
		log.Fatalln("min should not be greater than max")
	}

	rand.Seed(time.Now().Unix())

	c1 := stage1(min, max)
	c2 := stage2(c1)
	stage3(c2)
}
