package main

import (
	"fmt"
)

const size = 15

type node struct {
	Value int
	Next  *node
}

type hashTable struct {
	Table map[int]*node
	Size  int
}

func hashFunction(i, size int) int {
	return i % size
}

func insert(hash *hashTable, value int) int {
	index := hashFunction(value, hash.Size)
	element := node{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

func traverse(hash *hashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
			fmt.Printf("\n")
		}
	}
}

func lookUp(hash *hashTable, value int) bool {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				if t.Value == value {
					return true
				}
				t = t.Next
			}
		}
	}
	return false
}

func main() {
	table := make(map[int]*node, size)
	hash := &hashTable{Table: table, Size: size}

	fmt.Println("Number of spaces:", hash.Size)

	for i := 0; i < 120; i++ {
		insert(hash, i)
	}
	fmt.Println("120 values inserted.")
	for i := 10; i < 130; i++ {
		fmt.Printf("lookup %d: %v\n", i, lookUp(hash, i))
	}
}
