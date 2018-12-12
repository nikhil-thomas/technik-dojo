package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	MIN := 0
	MAX := 100
	SEED := time.Now().Unix()

	var LENGTH int64 = 8

	args := os.Args

	switch len(args) {
	case 2:
		var err error
		LENGTH, err = strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
	default:
		fmt.Println("Using default values")
	}

	rand.Seed(SEED)

	startChar := "!"
	var i int64 = 1

	for {
		randNum := random(MIN, MAX)
		newChar := string(startChar[0] + byte(randNum))
		fmt.Print(newChar)
		if i == LENGTH {
			break
		}
		i++
	}
	fmt.Print("\n")
}
