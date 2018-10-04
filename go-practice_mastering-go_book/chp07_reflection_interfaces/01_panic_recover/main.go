package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("calling f")
	f()
	delay1S()
	fmt.Println("returned normally from f()")

}

func f() {
	defer func() {
		if r := recover(); r != nil {
			delay1S()
			fmt.Println("recovered in f, panic:", r)
		}
	}()
	delay1S()
	fmt.Println("calling g")
	g(0)
	delay1S()
	fmt.Println("returned normally from g")
}

func g(v int) {
	if v > 3 {
		delay1S()
		fmt.Println("panicking from g")
		panic(fmt.Sprintf("code %v", v))
	}
	defer func() {
		delay1S()
		fmt.Println("defer in g", v)
	}()
	delay1S()
	fmt.Println("printing in g", v)
	g(v + 1)
}

func delay1S() {
	time.Sleep(1 * time.Second)
}
