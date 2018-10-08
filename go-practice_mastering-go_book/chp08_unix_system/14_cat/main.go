package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func printFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		io.WriteString(os.Stdout, scanner.Text())
		io.WriteString(os.Stdout, "\n")
	}
	return nil
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for _, file := range args {
		err := printFile(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
