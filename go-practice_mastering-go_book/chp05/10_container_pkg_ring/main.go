package main

import (
	"container/ring"
	"fmt"
)

var size int = 10

func main() {
	rng := ring.New(size + 1)
	fmt.Println("initial state: ring:", *rng)

	for i := 0; i < rng.Len()-1; i++ {
		rng.Value = i
		rng = rng.Next()
	}
	rng.Value = 2

	sum := 0

	rng.Do(func(x interface{}) {
		t := x.(int)
		sum = sum + t
	})
	fmt.Println("Sum:", sum)

	for i := 0; i < rng.Len()+2; i++ {
		rng = rng.Next()
		fmt.Print(rng.Value, " ")
	}
	fmt.Printf("\n")
}
