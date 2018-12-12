package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage\n%s <filename>\n", filepath.Base(args[0]))
		os.Exit(1)
	}

	filename := args[1]
	info, err := os.Stat(filename)
	if err != nil {
		log.Fatalln(err)
	}
	mode := info.Mode()
	fmt.Println(filename, "mode is ", mode.String()[1:10])
}
