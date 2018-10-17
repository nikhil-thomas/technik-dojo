package main

import (
	"fmt"

	"github.com/pkg/profile"
)

var v int

func prime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func multiply(a, b int) int {
	if a == 1 {
		return b
	}
	if a == 0 || b == 0 {
		return 0
	}
	if a < 0 {
		return -multiply(-a, b)
	}
	return b + multiply(a-1, b)
}

func main() {
	//defer profile.Start(profile.ProfilePath("/tmp")).Stop()
	defer profile.Start(profile.MemProfile).Stop()

	total := 0
	for i := 2; i < 200000; i++ {
		n := prime(i)
		if n {
			total++
		}
	}
	fmt.Println("total:", total)

	total = 0

	for i := 0; i < 5000; i++ {
		for j := 0; j < 400; j++ {
			k := multiply(i, j)
			v = k
			total++
		}
	}
	fmt.Println("total:", total)

}
