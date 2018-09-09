package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please supply one or more floats")
		os.Exit(1)
	}
	args := os.Args
	min, _ := strconv.ParseFloat(args[1], 64)
	max, _ := strconv.ParseFloat(args[1], 64)

	for i, str := range args {
		if i <= 1 {
			continue
		}
		val, _ := strconv.ParseFloat(str, 64)

		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
