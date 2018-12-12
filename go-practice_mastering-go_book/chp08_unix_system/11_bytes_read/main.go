package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("This is"))
	fmt.Fprintf(&buffer, " a string!\n")
	buffer.WriteTo(os.Stdout)
	buffer.WriteTo(os.Stdout)
	fmt.Println("::::::::::")

	buffer.Reset()
	buffer.Write([]byte("Mastering Go!"))
	r := bytes.NewReader([]byte(buffer.String()))
	fmt.Println(buffer.String())

	for {
		b := make([]byte, 3)
		n, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			continue
		}
		fmt.Printf("Read %s, byte coutn %d\n", b, n)
	}
}
