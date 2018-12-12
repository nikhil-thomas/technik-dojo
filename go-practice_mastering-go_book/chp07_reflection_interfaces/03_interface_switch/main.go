package main

import (
	"fmt"
)

type square struct {
	x float64
}

type circle struct {
	r float64
}

type rectangle struct {
	x float64
	y float64
}

func checkType(v interface{}) {
	switch t := v.(type) {
	case square:
		fmt.Printf("%+v is a square\n", t)
	case circle:
		fmt.Printf("%+v is a circle\n", t)
	case rectangle:
		fmt.Printf("%+v is a rectangle\n", t)
	default:
		fmt.Print("unknown type\n")
	}
}

func main() {
	c := circle{}
	checkType(c)
	s := square{}
	checkType(s)
	rec := rectangle{}
	checkType(rec)
	checkType(123)
}
