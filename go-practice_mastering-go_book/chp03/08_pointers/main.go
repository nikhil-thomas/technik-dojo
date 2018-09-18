package main

import (
	"fmt"
)

func getPointer(n *int) {
	*n = *n * *n
}

func returnPointer(n int) *int {
	v := n * n
	return &v
}

func main() {
	i := -10
	j := 25

	pI := &i
	pJ := &j

	fmt.Println("pI mem", pI)
	fmt.Println("pJ mem", pJ)
	fmt.Println("pI val", *pI)
	fmt.Println("pJ val", *pJ)

	getPointer(pJ)
	fmt.Println("j:", j)
	k := returnPointer(12)
	fmt.Println(*k)
	fmt.Println(k)
}
