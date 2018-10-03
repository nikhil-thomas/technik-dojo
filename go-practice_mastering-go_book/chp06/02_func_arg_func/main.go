package main

import (
	"fmt"
)

func doubleFunc(i int) int {
	return i + i
}

func squareFunc(i int) int {
	return i * i
}

func funcArgFunc(f func(int) int, v int) int {
	return f(v)
}

func main() {
	fmt.Println("double 2:", funcArgFunc(doubleFunc, 2))
	fmt.Println("square 2:", funcArgFunc(squareFunc, 2))
	fmt.Println("inline (i * i * i) 2:", funcArgFunc(func(i int) int {
		return i * i * i
	}, 2))
}
