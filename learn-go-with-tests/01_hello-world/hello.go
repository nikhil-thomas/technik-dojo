package main

import (
	"fmt"
)

const helloPrefix = "Hello, "

// Hello returns greeting message
func Hello(name string) string {
	return helloPrefix + name
}

// Hello2 returns greeting message based in input
func Hello2(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
