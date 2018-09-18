package main

import (
	"fmt"
)

// Digit is a named int type
type Digit int

// Power2 is a name dint type
type Power2 int

// PI is the math constant pi
const PI = 3.1415926

const (
	zero Digit = iota
	one
	two
	three
	four
)

const (
	p2_0 Power2 = 1 << iota
	_
	p2_2 Power2 = 1 << iota
	_
	p2_4 Power2 = 1 << iota
	_
	p2_6 Power2 = 1 << iota
	_
)

const (
	// C1 is a string constant
	C1 = "C1C1C1"

	// C2 is a string constant
	C2 = "C2C2C2"

	// C3 is a string constant
	C3 = "C3C3C3"
)

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(PI)
	fmt.Println(one, two)
	fmt.Printf("\n")
	fmt.Println("2^0:", p2_0)
	fmt.Println("2^2:", p2_2)
	fmt.Println("2^4:", p2_4)
	fmt.Println("2^6:", p2_6)
}
