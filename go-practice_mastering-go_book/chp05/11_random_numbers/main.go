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

	TOTAL := 100
	SEED := time.Now().Unix()

	args := os.Args

	switch len(args) {
	case 2:
		fmt.Println("Usage: ./main MIN MAX TOTAL SEED")
		var err error
		MIN, err = strconv.Atoi(args[1])
		if err != nil {
			log.Fatalln(err)
		}
		MAX = MIN + 10
	case 3:
		fmt.Println("Usage: ./main MIN MAX TOTAL SEED")
		var err error
		MIN, err = strconv.Atoi(args[1])
		if err != nil {
			log.Fatalln(err)
		}
		MAX, err = strconv.Atoi(args[2])
		if err != nil {
			log.Fatalln(err)
		}
	case 4:
		fmt.Println("Usage: ./main MIN MAX TOTAL SEED")
		var err error
		MIN, err = strconv.Atoi(args[1])
		if err != nil {
			log.Fatalln(err)
		}
		MAX, err = strconv.Atoi(args[2])
		if err != nil {
			log.Fatalln(err)
		}
		TOTAL, err = strconv.Atoi(args[3])
		if err != nil {
			log.Fatalln(err)
		}
	case 5:
		fmt.Println("Usage: ./main MIN MAX TOTAL SEED")
		var err error
		MIN, err = strconv.Atoi(args[1])
		if err != nil {
			log.Fatalln(err)
		}
		MAX, err = strconv.Atoi(args[2])
		if err != nil {
			log.Fatalln(err)
		}
		TOTAL, err = strconv.Atoi(args[3])
		if err != nil {
			log.Fatalln(err)
		}
		SEED, err = strconv.ParseInt(args[4], 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
	default:
		fmt.Println("Usage: ./main MIN MAX TOTAL SEED")
		fmt.Println("Using default values")
	}

	fmt.Println("MIN:", MIN, "MAX:", MAX, "TOTAL:", TOTAL, "SEED:", SEED)

	rand.Seed(SEED)
	for i := 0; i < TOTAL; i++ {
		rndNum := random(MIN, MAX)
		fmt.Print(rndNum, " ")
	}
	fmt.Printf("\n\n")

}
