package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)

	fmt.Println("Web Assembly initialized")
	registerCallbacks()

	<-c
}

func add(i []js.Value) {
	val1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	val2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(val1)
	int2, _ := strconv.Atoi(val2)

	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1+int2)
	println(int1 + int2)
}

func subtract(i []js.Value) {
	val1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	val2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(val1)
	int2, _ := strconv.Atoi(val2)

	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1-int2)
	println(int1 - int2)
}

func registerCallbacks() {
	js.Global().Set("add", js.NewCallback(add))
	js.Global().Set("sub", js.NewCallback(subtract))
}
