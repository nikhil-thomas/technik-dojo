package main

import (
	"fmt"
)

const helloPrefix = "Hello, "
const spanishHelloprefix = "Hola, "
const frenchHelloprefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

// Hello returns greeting message
func Hello(name string) string {
	return helloPrefix + name
}

// Hello2 returns greeting message based on input name
func Hello2(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name
}

// Hello3 returns greeting message based on input language and name
func Hello3(language, name string) string {
	if name == "" {
		name = "World"
	}
	prefix := greetingPrefix(language)
	return prefix + name
}

func greetingPrefix(language string) string {
	switch language {
	case spanish:
		return spanishHelloprefix
	case french:
		return frenchHelloprefix
	default:
		return helloPrefix
	}
}

func main() {
	fmt.Println(Hello("world"))
}
