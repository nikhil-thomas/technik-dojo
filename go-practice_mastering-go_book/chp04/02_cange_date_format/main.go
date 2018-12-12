package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("please provide one text file to parse")
		os.Exit(1)
	}

	filename := args[1]

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file %s", err)
		os.Exit(1)
	}
	defer f.Close()

	notAMatch := 0

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error reading file:", err)
		}
		fmt.Printf("\nline orig: %s", line)
		r1 := regexp.MustCompile(`.*\[(\d\d\/\w+\/\d\d\d\d:\d\d:\d\d:\d\d.*)\] .*`)
		if r1.MatchString(line) {
			match := r1.FindStringSubmatch(line)
			d1, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])
			if err != nil {
				notAMatch++
				continue
			}
			newFormat := d1.Format(time.Stamp)
			fmt.Printf("line edit: %s\n", strings.Replace(line, match[1], newFormat, 1))
		}

		r2 := regexp.MustCompile(`.*\[(\w+\-\d\d-\d\d:\d\d:\d\d:\d\d.*)\] .*`)
		if r2.MatchString(line) {
			match := r2.FindStringSubmatch(line)
			d1, err := time.Parse("Jan-02-06:15:04:05 -0700", match[1])
			if err != nil {
				notAMatch++
				continue
			}
			newFormat := d1.Format(time.Stamp)
			fmt.Printf("line edit: %s\n", strings.Replace(line, match[1], newFormat, 1))
		}
	}
	fmt.Printf("%d lines did not match\n", notAMatch)
}
