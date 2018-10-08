package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Printf("usage\n%s <dir>  [<dir>...]", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	for _, d := range args {
		filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
			fmt.Printf("%s : %s : %d\n", path, info.Name(), info.Size())
			return nil
		})
	}
}
