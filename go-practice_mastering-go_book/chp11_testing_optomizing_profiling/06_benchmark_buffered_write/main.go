package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var buffersize int
var filesize int

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func createBuffer(buf *[]byte, count int) {
	*buf = []byte{}
	if count == 0 {
		return
	}

	for i := 0; i < count; i++ {
		randInt := random(0, 100)
		intByte := byte(randInt)
		if len(*buf) > count {
			return
		}
		*buf = append(*buf, intByte)
	}
}

func create(dst string, b, f int) error {
	var err error
	_, err = os.Stat(dst)
	if err == nil {
		return fmt.Errorf("file %s already exists", dst)
	}
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	buf := make([]byte, 0)
	for {
		createBuffer(&buf, b)
		buf = buf[:b]
		if _, err := destination.Write(buf); err != nil {
			return err
		}
		if f < 0 {
			break
		}
		f = f - len(buf)
	}
	return err
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage\n%s <buffersize> <filesize>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	rand.Seed(time.Now().Unix())
	output := "/tmp/randomFile"
	var err error
	buffersize, err = strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	filesize, err = strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}

	err = create(output, buffersize, filesize)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.Remove(output)
	if err != nil {
		log.Fatalln(err)
	}
}
