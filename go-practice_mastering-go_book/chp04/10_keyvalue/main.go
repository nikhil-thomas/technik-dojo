package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type myElement struct {
	Name    string
	Surname string
	ID      string
}

var data = make(map[string]myElement)

func add(k string, n myElement) bool {
	if k == "" {
		return false
	}

	if lookup(k) == nil {
		data[k] = n
		return true
	}
	return false
}

func deleteKey(k string) bool {
	if lookup(k) != nil {
		delete(data, k)
		return true
	}
	return false
}

func lookup(k string) *myElement {
	val, ok := data[k]
	if ok {
		return &val
	}
	return nil
}

func change(k string, n myElement) bool {
	data[k] = n
	return true
}

func print() {
	fmt.Println("Key-Value")
	for k, v := range data {
		fmt.Printf("%s: %v\n", k, v)
	}
	fmt.Printf("\n")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		tokens := strings.Fields(text)

		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "", "", "", "")
		case 2:
			tokens = append(tokens, "", "", "")
		case 3:
			tokens = append(tokens, "", "")
		case 4:
			tokens = append(tokens, "")
		}

		switch tokens[0] {
		case "PRINT":
			print()
		case "STOP":
			return
		case "DELETE":
			if !deleteKey(tokens[1]) {
				fmt.Println("Delete opereation failed")
			}
		case "ADD":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !add(tokens[1], n) {
				fmt.Println("Add operation failed")
			}
		case "LOOKUP":
			n := lookup(tokens[1])
			if n != nil {
				fmt.Printf("%s: %v\n", tokens[1], *n)
			} else {
				fmt.Println(tokens[1], "not found")
			}
		case "CHANGE":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !change(tokens[1], n) {
				fmt.Println("Update operation failed")
			}
		default:
			fmt.Println("Unknown command")
		}

	}
}
