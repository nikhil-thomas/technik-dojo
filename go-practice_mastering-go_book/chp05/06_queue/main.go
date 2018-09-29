package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

var size = 0
var queue = new(node)

func push(t *node, v int) bool {
	if queue == nil {
		queue = &node{v, nil}
		size++
		return true
	}

	t = &node{v, nil}
	t.next = queue
	queue = t
	size++
	return true
}

func pop(t *node) (int, bool) {
	if size == 0 {
		return 0, false
	}

	if size == 1 {
		queue = nil
		size--
		return t.value, true
	}

	prev := t

	for (t.next) != nil {
		prev = t
		t = t.next
	}
	v := (prev.next).value
	prev.next = nil
	size--
	return v, true
}

func traverse(t *node) {
	if size == 0 {
		fmt.Println("Empty Queue")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.value)
		t = t.next
	}
	fmt.Printf("\n")
}

func main() {
	queue = nil
	push(queue, 10)
	fmt.Println("size:", size)
	traverse(queue)

	v, ok := pop(queue)
	if ok {
		fmt.Println("Popped:", v)
	}
	fmt.Println("size:", size)
	traverse(queue)

	for i := 0; i < 5; i++ {
		push(queue, i)
	}
	traverse(queue)
	fmt.Println("size:", size)

	v, ok = pop(queue)
	if ok {
		fmt.Println("Popped:", v)
	}
	fmt.Println("size:", size)

	v, ok = pop(queue)
	if ok {
		fmt.Println("Popped:", v)
	}
	fmt.Println("size:", size)
	traverse(queue)
}
