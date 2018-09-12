package main

import (
	"fmt"
	"unsafe"
)

// note: Go compiler implements the unsafe package
// when you import it into your programs.

func main() {
	var value int64 = 5
	var p1 = &value
	var p2 = (*int32)(unsafe.Pointer(p1))

	fmt.Println("*p1:", *p1)
	fmt.Println("*p2:", *p2)

	*p1 = 1234567890123456789
	fmt.Println("*p1:", *p1)

	// value will be 32 bit
	// as 32 bit int pointer cannot store 64 bit integer
	fmt.Println("*p2:", *p2)

	*p1 = 12345678
	fmt.Println("*p1:", *p1)
	fmt.Println("*p2:", *p2)
}
