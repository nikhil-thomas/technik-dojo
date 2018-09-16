package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 5)
	fmt.Println("s1:", s1)
	reslice := s1[1:3]
	fmt.Println("reslice:", reslice)

	reslice[0] = -123
	reslice[1] = 9999

	fmt.Println("s1:", s1)
	fmt.Println("reslice:", reslice)
}
