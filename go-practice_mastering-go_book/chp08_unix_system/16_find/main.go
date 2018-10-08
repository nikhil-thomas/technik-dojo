package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var flagD bool
var flagF bool

func walkFunc(p string, info os.FileInfo, err error) error {
	mode := info.Mode()

	if mode.IsRegular() && flagF {
		fmt.Print("+ ")
	}

	if mode.IsDir() && flagF {
		fmt.Print("* ")
	}
	fmt.Printf("%s\n", p)
	return nil
}

func main() {
	starD := flag.Bool("d", false, "mark directories")
	starF := flag.Bool("f", false, "mark file")
	flag.Parse()

	path := "."
	if len(flag.Args()) == 1 {
		path = flag.Arg(0)
	}

	flagD = *starD
	flagF = *starF

	err := filepath.Walk(path, walkFunc)

	if err != nil {
		log.Fatalln(err)
	}
}
