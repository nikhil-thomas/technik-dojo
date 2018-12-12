package main

import (
	"container/heap"
	"fmt"
)

type heapFloat32 []float32

func (n *heapFloat32) Pop() interface{} {
	old := *n
	x := old[len(old)-1]
	new := old[0 : len(old)-1]
	*n = new
	return x
}

func (n *heapFloat32) Push(x interface{}) {
	*n = append(*n, x.(float32))
}

func (n heapFloat32) Len() int {
	return len(n)
}

func (n heapFloat32) Less(a, b int) bool {
	return n[a] < n[b]
}

func (n heapFloat32) Swap(a, b int) {
	n[a], n[b] = n[b], n[a]
}

func main() {
	hp := &heapFloat32{1.2, 2.1, 3.1, -100.1}
	fmt.Printf("%v\n", hp)
	heap.Init(hp)
	size := len(*hp)
	fmt.Println("Heap size:", size)
	fmt.Printf("%v\n", hp)

	hp.Push(float32(-100.2))
	hp.Push(float32(0.2))

	fmt.Println("Heap size:", size)
	fmt.Printf("%v\n", hp)

	heap.Init(hp)
	fmt.Printf("%v\n", hp)

	fmt.Println(hp.Pop())
	fmt.Println(hp.Pop())
	fmt.Println(hp.Pop())
	fmt.Println(hp.Pop())
}
