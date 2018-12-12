package main

import (
	"fmt"
	"unicode"
)

func main() {
	const str = "\x99\x00ab\x50\x00\x23\x50\x29\x9c"

	for i := 0; i < len(str); i++ {
		if unicode.IsPrint(rune(str[i])) {
			fmt.Printf("%d : %x : %c\n", i, str[i], str[i])
		} else {
			fmt.Printf("not printable %d : %x\n", i, str[i])
		}
	}
}
