package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 4}
	fmt.Println("a:", "len:", len(a), "cap:", cap(a))
	fmt.Println(a)
	fmt.Printf("\n")

	a = nil

	fmt.Println("a:", "len:", len(a), "cap:", cap(a))
	fmt.Println(a)
	fmt.Printf("\n")

	a2 := make([][]int, 3)
	fmt.Println("a2:", "len:", len(a2), "cap:", cap(a2))
	fmt.Println(a2)
	fmt.Printf("\n")

	// for i, _ := range a2 {
	// same as below
	for i := range a2 {
		for j := 1; j <= 3; j++ {
			a2[i] = append(a2[i], j)
		}
	}

	fmt.Println("a2:", "len:", len(a2), "cap:", cap(a2))
	fmt.Println(a2)
	fmt.Printf("\n")
}
