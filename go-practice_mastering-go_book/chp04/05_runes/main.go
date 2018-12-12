package main

import "fmt"

func main() {
	const r1 = '€'
	fmt.Println("int32(r1):", r1)
	fmt.Printf("Hex r1: %x\n", r1)
	fmt.Printf("as a string: %s\n", r1)
	fmt.Printf("as a character r1: %c\n", r1)

	str := []byte("John")
	for i, v := range str {
		fmt.Println(i, v)
		fmt.Printf("Char: %c\n", str[i])
	}
	fmt.Printf("%s\n", str)
}
