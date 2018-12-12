package main

import (
	"fmt"
)

const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"

var s2 = "€£³"

func main() {
	fmt.Println("sLiteral:", sLiteral)
	fmt.Println("s2:", s2)
	fmt.Printf("\n")

	fmt.Print("sLiteral '%''x': ")
	fmt.Printf("%x\n", sLiteral)
	fmt.Print("s2 '%''x': ")
	fmt.Printf("%x\n", s2)
	fmt.Printf("\n")

	fmt.Println("len sLiteral:", len(sLiteral))
	fmt.Println("len s2:", len(s2))
	fmt.Printf("\n")

	for i := 0; i < len(sLiteral); i++ {
		fmt.Print("sLiteral '%''x': ")
		fmt.Printf("%d: %x\n", i, sLiteral[i])
	}
	fmt.Printf("\n")
	for i := 0; i < len(s2); i++ {
		fmt.Print("s2 '%''x': ")
		fmt.Printf("%d: %x\n", i, s2[i])
	}
	fmt.Printf("\n")

	fmt.Printf("sLiteral q: %q\n", sLiteral)
	fmt.Printf("s2 q: %q\n", s2)

	fmt.Printf("\n")
	fmt.Printf("sLiteral +q: %+q\n", sLiteral)
	fmt.Printf("s2 +q: %+q\n", s2)

	fmt.Printf("\n")
	fmt.Printf("sLiteral ( x): % x\n", sLiteral)
	fmt.Printf("s2 ( x): % x\n", s2)

	fmt.Printf("\n")
	fmt.Printf("sLiteral s: %s\n", sLiteral)
	fmt.Printf("s2 s: %s\n", s2)

}
