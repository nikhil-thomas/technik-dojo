package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func printList(l *list.List) {
	for t := l.Back(); t != nil; t = t.Prev() {
		fmt.Print(t.Value, " ")
	}
	fmt.Printf("\n")

	for t := l.Front(); t != nil; t = t.Next() {
		fmt.Print(t.Value, " ")
	}
	fmt.Printf("\n")
}

func main() {
	values := list.New()

	e1 := values.PushBack("One")
	e2 := values.PushBack("Two")

	fmt.Println("e1", e1, "e2", e2)
	fmt.Printf("%v\n %q\n", e1, e1)
	printList(values)

	values.PushFront("Three")
	values.InsertBefore("Four", e1)
	values.InsertAfter("Five", e2)

	fmt.Println("push back")
	printList(values)
	values.PushBackList(values)
	printList(values)

	fmt.Printf("\ninit\n")
	values.Init()
	printList(values)

	fmt.Printf("After Init(): %v\n", values)

	fmt.Printf("\n\nloop push front\n")
	for i := 0; i < 20; i++ {
		values.PushFront(strconv.Itoa(i))
	}
	printList(values)
}
