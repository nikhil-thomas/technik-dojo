package main

import (
	"fmt"
)

type node struct {
	Value int
	Next  *node
}

var root = new(node)

func addNode(t *node, v int) int {
	if root == nil {
		t = &node{v, nil}
		root = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node alread exists:", v)
		return -1
	}

	if t.Next == nil {
		t.Next = &node{v, nil}
		return -2
	}

	return addNode(t.Next, v)
}

func traverse(t *node) {

	if t == nil {
		fmt.Println("empty list")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Printf("\n")
}

func lookUpNode(t *node, v int) bool {
	if root == nil {
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookUpNode(t.Next, v)
}

func size(t *node) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	i := 0

	for t != nil {
		i++
		t = t.Next
	}
	return i
}

func main() {
	fmt.Println(root)
	root = nil
	traverse(root)

	addNode(root, 1)
	addNode(root, -1)
	traverse(root)

	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 45)
	addNode(root, 5)
	addNode(root, 5)
	traverse(root)

	addNode(root, 100)
	traverse(root)

	if lookUpNode(root, 100) {
		fmt.Println("Node 100 exists!")
	} else {
		fmt.Println("Node 100 does not exists!")
	}

	if lookUpNode(root, -100) {
		fmt.Println("Node -100 exists!")
	} else {
		fmt.Println("Node -100 does not exists!")
	}

}
