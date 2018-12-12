package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var dateStr string
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s string\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	dateStr = os.Args[1]

	d, err := time.Parse("2006 Jan 02", dateStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Full:", d)
	fmt.Println("Time:", d.Day(), d.Month(), d.Year())
}
