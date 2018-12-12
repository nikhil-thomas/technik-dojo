package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("arg url required")
	}
	data, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer data.Body.Close()

	written, err := io.Copy(os.Stdout, data.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\nWritten: %d\n", written)
}
