package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("usage: select_column <column_number> <file_name>\n")
		os.Exit(1)
	}

	column, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	if column < 0 {
		fmt.Println("invalid column number")
		os.Exit(1)
	}

	for _, filename := range arguments[2:] {
		fmt.Println("processing:", filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error opening file: %s", err)
			continue
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')

			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Printf("Error opening file: %s", err)
				continue
			}
			data := strings.Fields(line)
			if len(data) >= column {
				fmt.Println(data[column-1])
			}
		}
	}
}
