package main

import (
	"fmt"
)

func b() {
	fmt.Println("inside b()")
	panic("Panic in b()")
	fmt.Println("exiting b()")
}

func a() {
	fmt.Println("inside a()")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside a()")
		}
	}()
	fmt.Println("about to call b()")
	b()
	fmt.Println("inside a() b() exited")
	fmt.Println("exiting a()")
}

func main() {
	fmt.Println("start: main")
	a()
	fmt.Println("end: main")
}
