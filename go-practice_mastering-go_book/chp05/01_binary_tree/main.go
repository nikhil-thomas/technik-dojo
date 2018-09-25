package main

import (
	"fmt"
	"math/rand"
	"time"
)

type tree struct {
	Left  *tree
	Value int
	Right *tree
}

func traverse(t *tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value, ", ")
	traverse(t.Right)
}

// create creates a tree with random integers
func create(n int) *tree {
	var t *tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		if t != nil {
			fmt.Print("traverse: ")
			traverse(t)
			fmt.Print("\n")
		}
		temp := rand.Intn(n * 2)
		fmt.Println("insert:", temp)
		t = insert(t, temp)
	}
	return t
}

func insert(t *tree, n int) *tree {
	if t == nil {
		return &tree{nil, n, nil}
	}
	if n == t.Value {
		fmt.Println(n, "already exists")
		return t
	}

	if n < t.Value {
		t.Left = insert(t.Left, n)
		return t
	}
	t.Right = insert(t.Right, n)
	return t
}

func main() {
	t := create(10)
	traverse(t)
	fmt.Print("\n\n")
	fmt.Println("insert -10 and -2")
	insert(t, -10)
	insert(t, -2)
	traverse(t)
	fmt.Print("\n")
}
