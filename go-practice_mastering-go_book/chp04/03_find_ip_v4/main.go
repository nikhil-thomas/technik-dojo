package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

func findIP(s string) string {
	partIP := "(255[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(s)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("usage: %s <logfile>\n", filepath.Base(args[0]))
		os.Exit(1)
	}

	for _, filename := range args[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error openning file %s: %s", filename, err)
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
				fmt.Println("Error reading line:", err)
				continue
			}
			ip := findIP(line)
			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			}
			fmt.Println(ip)
		}
	}
}
