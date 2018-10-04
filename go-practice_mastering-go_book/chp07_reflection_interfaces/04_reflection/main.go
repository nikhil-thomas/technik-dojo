// Fields should be exported in this example
// reason
// panic: reflect.Value.Interface:
//        cannot return value obtained from unexported field or method

package main

import (
	"fmt"
	"os"
	"reflect"
)

type a struct {
	X int
	Y float64
	Z string
}

type b struct {
	F int
	G int
	H string
	I float64
}

func main() {
	x := 100
	xRefl := reflect.ValueOf(&x)
	fmt.Println("reflect.ValueOf(&x):", xRefl)
	fmt.Printf("type reflect.ValueOf(&x): %T\n", xRefl)
	fmt.Println("Elem():", xRefl.Elem())
	fmt.Println("Type():", xRefl.Type())
	fmt.Println("::::::::::")

	p := a{100, 200.1, "struct a"}
	q := b{1, 2, "struct b", -1.2}
	var r reflect.Value

	args := os.Args
	if len(args) == 1 {
		r = reflect.ValueOf(&p).Elem()
	} else {
		r = reflect.ValueOf(&q).Elem()
	}

	iType := r.Type()
	fmt.Printf("Type: %s\n", iType)
	fmt.Printf("Num fields: %d\n", r.NumField())

	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("Name: %s, ", iType.Field(i).Name)
		fmt.Printf("Type: %s, ", r.Field(i).Type())
		fmt.Printf("Value: %v, \n", r.Field(i).Interface())
	}

}
