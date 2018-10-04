package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered panic:", r)
		}
	}()

	var typeInt interface{} = 123
	v, ok := typeInt.(int)
	if ok {
		fmt.Println("type assertion success, value:", v)
	} else {
		fmt.Println("type assertion failed without panicking")
	}

	v2, ok := typeInt.(float64)
	if ok {
		fmt.Println("type assertion success, value:", v2)
	} else {
		fmt.Println("type assertion failed without panicking")
	}

	v3 := typeInt.(int)
	fmt.Println("type assertion success, value:", v3)

	v4 := typeInt.(float64)
	fmt.Println("type assertion success, value:", v4)

}
