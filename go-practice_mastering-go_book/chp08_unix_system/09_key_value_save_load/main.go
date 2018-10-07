package main

import (
	"bufio"
	"encoding/gob"
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
var datafile = "/tmp/dataFile.gob"

func save() error {
	fmt.Println("Saving", datafile)

	err := os.Remove(datafile)
	if err != nil {
		fmt.Println(err)
	}
	saveTo, err := os.Create(datafile)
	if err != nil {
		return err
	}
	defer saveTo.Close()

	encoder := gob.NewEncoder(saveTo)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func load() error {
	fmt.Println("Loading", datafile)
	loadFrom, err := os.Open(datafile)
	if err != nil {
		return err
	}
	defer loadFrom.Close()

	decoder := gob.NewDecoder(loadFrom)
	err = decoder.Decode(&data)
	if err != nil {
		return err
	}
	return nil
}

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
	err := load()
	if err != nil {
		fmt.Println(err)
		fmt.Println("starting blank store")
	}

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
			err = save()
			if err != nil {
				fmt.Println(err)
			}
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
