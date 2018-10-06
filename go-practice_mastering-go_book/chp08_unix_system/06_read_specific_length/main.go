package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func readSize(f *os.File, size int) ([]byte, error) {
	buffer := make([]byte, size)

	n, err := f.Read(buffer)
	if err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}
	return buffer[:n], nil
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Printf("usage\n%s <buffer size> <file name>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	bufferSize, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(err)
	}

	filename := args[1]

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	for {
		data, err := readSize(f, bufferSize)
		if err != nil {
			log.Fatalln(err)
		}
		if data == nil {
			break
		}
		fmt.Print(string(data))
	}
}
