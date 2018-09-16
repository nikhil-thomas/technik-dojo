package main

import (
	"fmt"
)

func main() {
	a6 := []int{1, 2, 3, 4, 5, 6}
	a4 := []int{10, 20, 30, 40}
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)

	copy(a6, a4)
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	fmt.Println()

	a6 = []int{1, 2, 3, 4, 5, 6}
	a4 = []int{10, 20, 30, 40}
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)

	copy(a4, a6)
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	fmt.Println()

	array6 := [6]int{1, 2, 3, 4, 5, 6}
	s4 := []int{10, 20, 30, 40}
	fmt.Println("array6:", array6)
	fmt.Println("s4:", s4)

	copy(s4, array6[:])
	fmt.Println("array6:", array6)
	fmt.Println("s4:", s4)
	fmt.Println()

	array4 := [4]int{1, 2, 3, 4}
	s6 := []int{10, 20, 30, 40, 50, 60}
	fmt.Println("array4:", array4)
	fmt.Println("s6:", s6)

	copy(s6, array4[:])
	fmt.Println("array4:", array4)
	fmt.Println("s6:", s6)
	fmt.Println()
}
