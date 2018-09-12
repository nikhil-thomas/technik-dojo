package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Give one or more floats")
		os.Exit(1)
	}

	arguments := os.Args
	var err error = errors.New("An error")
	k := 1
	var n float64

	// check if there is atleast one float
	for err != nil {
		if k >= len(arguments) {
			fmt.Println("None of the arguments is a float")
			return
		}
		n, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}
	min, max := n, n

	for i, v := range arguments {
		if i == 0 {
			continue
		}
		n, err = strconv.ParseFloat(v, 64)
		if err == nil {
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
