package main

import (
	"fmt"
)

type node struct {
	Value    int
	Previous *node
	Next     *node
}

// head points to head of doubly linked list
var head = &node{}

func addNode(t *node, v int) int {
	if head == nil {
		t = &node{v, nil, nil}
		head = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return -1
	}

	if t.Next == nil {
		temp := t
		t.Next = &node{v, temp, nil}
		return -2
	}
	return addNode(t.Next, v)
}

func traverse(t *node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}

	fmt.Printf("\n")
}

func reversePrint(t *node) {
	if t == nil {
		fmt.Println("-> Empty list")
		return
	}

	tail := t

	for t != nil {
		tail = t
		t = t.Next
	}

	for tail != nil {
		fmt.Printf("%d -> ", tail.Value)
		tail = tail.Previous
	}
}

func size(t *node) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	n := 0
	for t != nil {
		n++
		t = t.Next
	}
	return n
}

func lookupNode(t *node, v int) bool {
	if head == nil {
		return false
	}

	if v == t.Value {
		return true
	}
	if t.Next == nil {
		return false
	}
	return lookupNode(t, v)
}

func main() {
	fmt.Println(head)
	head = nil
	fmt.Println(head)
	traverse(head)

	addNode(head, 1)
	addNode(head, 1)

	traverse(head)

	addNode(head, 10)
	addNode(head, 5)
	addNode(head, 0)
	addNode(head, 0)

	traverse(head)

	addNode(head, 100)

	fmt.Println("Size:", size(head))

	traverse(head)
	reversePrint(head)
}
