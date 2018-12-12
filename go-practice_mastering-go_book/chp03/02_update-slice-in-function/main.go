package main

import (
	"fmt"
)

func main() {
	fmt.Println("Exp 1")
	nums := []int{1, 2, 3}
	fmt.Println("Inside main")
	printSlice(nums)
	add10(nums)
	fmt.Println("Inside main")
	printSlice(nums)

	add3Items(nums)
	fmt.Println("Inside main")
	printSlice(nums)

	nums = add3ItemsRtrn(nums)
	fmt.Println("Inside main")
	printSlice(nums)

	fmt.Printf("\nExp 2")
	nums = make([]int, 0, 6)
	nums = append(nums, 1, 2, 3)
	fmt.Println("Inside main")
	printSlice(nums)

	add10(nums)
	fmt.Println("Inside main")
	printSlice(nums)

	add3Items(nums)
	fmt.Println("Inside main")
	printSlice(nums)

	nums = add3ItemsRtrn(nums)
	fmt.Println("Inside main")
	printSlice(nums)

}

// printSlice prints the contents of an int slice
func printSlice(vals []int) {
	fmt.Println("cap:", cap(vals), "len:", len(vals))
	for _, v := range vals {
		fmt.Print(v, " ")
	}
	fmt.Printf("\n\n")
}

// add10 adds 10 to each element of an int slice
func add10(vals []int) {
	for i := 0; i < len(vals); i++ {
		vals[i] += 10
	}
	fmt.Println("inside add10")
	printSlice(vals)
}

// add3Items appends 3 integers to an int slice
func add3Items(vals []int) {
	for i := 100; i <= 300; i += 100 {
		vals = append(vals, i)
	}
	fmt.Println("inside doublLength")
	printSlice(vals)
}

// add3ItemsRtrn appends 3 integers to an int slice
// and returns the slice
func add3ItemsRtrn(vals []int) []int {
	for i := 100; i <= 300; i += 100 {
		vals = append(vals, i)
	}
	fmt.Println("inside doublLengthRtrn")
	printSlice(vals)
	return vals
}
