package main

import (
	"fmt"
)

func funcReturnFunc() func() int {
	i := 0
	fmt.Println("funcReturnFunc, i:", i)
	return func() int {
		i++
		fmt.Println("func () int, i:", i)
		return i
	}
}

func main() {
	f1 := funcReturnFunc()
	f2 := funcReturnFunc()

	fmt.Println("f1.1:", f1())
	fmt.Println("f2.1:", f2())

	fmt.Println("f1.2:", f1())
	fmt.Println("f1.3:", f1())

	fmt.Println("f2.2:", f2())
	fmt.Println("f2.2:", f2())
}
