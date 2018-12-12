package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	s := []byte("Data to write\n")

	f1, err := os.Create("f1")
	if err != nil {
		log.Fatalln(err)
	}
	defer f1.Close()
	fmt.Fprintf(f1, string(s))

	f2, err := os.Create("f2")
	if err != nil {
		log.Fatalln(err)
	}
	defer f2.Close()
	n, err := f2.WriteString(string(s))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("wrote %d bytes\n", n)

	f3, err := os.Create("f3")
	if err != nil {
		log.Fatalln(err)
	}
	defer f3.Close()

	w := bufio.NewWriter(f3)
	n, err = w.WriteString(string(s))
	if err != nil {
		log.Fatalln(err)
	}
	w.Flush()
	fmt.Printf("f3: wrote %d bytes\n", n)

	f4 := "f4"
	err = ioutil.WriteFile(f4, s, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	f5, err := os.Create("f5")
	if err != nil {
		log.Fatalln(err)
	}
	defer f5.Close()

	n, err = io.WriteString(f5, string(s))
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Wrote %d bytes\n", n)
}
