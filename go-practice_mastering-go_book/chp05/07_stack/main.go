package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

var stack *node
var size = 0

func push(v int) bool {
	if stack == nil {
		stack = &node{v, nil}
		size++
		return true
	}

	temp := &node{v, nil}
	temp.next = stack
	stack = temp
	size++
	return true
}

func pop() (int, bool) {
	if stack == nil {
		return 0, false
	}

	t := stack

	if size == 1 {
		stack = nil
		size--
		return t.value, true
	}

	stack = stack.next
	size--
	return t.value, true
}

func traverse() {
	if stack == nil {
		fmt.Println("Empty stack")
		return
	}

	t := stack

	for t != nil {
		fmt.Printf("%d -> ", t.value)
		t = t.next
	}
	fmt.Printf("\n")
}

func main() {
	v, ok := pop()

	if ok {
		fmt.Println("poped:", v)
	} else {
		fmt.Println("pop failed")
	}

	push(100)
	traverse()
	push(200)
	traverse()

	for i := 0; i < 10; i++ {
		push(i)
	}
	traverse()

	for i := 0; i < 15; i++ {
		v, ok = pop()

		if ok {
			fmt.Println("poped:", v)
		} else {
			break
		}
	}

	fmt.Printf("\n")
	traverse()

}
