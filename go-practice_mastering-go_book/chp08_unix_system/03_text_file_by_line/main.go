package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func lineByLine(file string) error {
	var err error

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error reading file", err)
			break
		}
		fmt.Print(line)
	}
	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Fatalln("Provide atleast one file to process")
	}

	for _, file := range flag.Args() {
		err := lineByLine(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
