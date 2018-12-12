package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("usage: %s <number>", filepath.Base(args[0]))
	}
	num, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Error parsing number:", err)
	} else {
		switch {
		case num < 0:
			fmt.Println(num, "is less than 0")
		case num > 0:
			fmt.Println(num, "is greater than 0")
		default:
			fmt.Println(num, "is equal to 0")
		}
	}

	str := args[1]

	switch str {
	case "5":
		fmt.Println("five")
	case "0":
		fmt.Println("zero")
	default:
		fmt.Println("do not care")
	}

	negative := regexp.MustCompile(`-`)
	floatingPt := regexp.MustCompile(`\d?\.\d`)
	email := regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)

	switch {
	case negative.MatchString(str):
		fmt.Println("negative number")
	case floatingPt.MatchString(str):
		fmt.Println("floating pt number")
	case email.MatchString(str):
		fmt.Println("email string")
		fallthrough
	default:
		fmt.Println("something else")
	}
}
