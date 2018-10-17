package main

import (
	"fmt"
	"os"
	"testing"
)

var errCreate error

func benchmarkCreate(b *testing.B, buffersize, filesize int) {
	var err error
	for i := 0; i < b.N; i++ {
		err = create("/tmp/randomfile", buffersize, filesize)
	}
	errCreate = err
	err = os.Remove("/tmp/randomfile")
	if err != nil {
		fmt.Println(err)
	}
}

func BenchmarkCreate1(b *testing.B) {
	benchmarkCreate(b, 1, 1000000)
}

func BenchmarkCreate2(b *testing.B) {
	benchmarkCreate(b, 2, 1000000)
}

func BenchmarkCreate3(b *testing.B) {
	benchmarkCreate(b, 4, 1000000)
}

func BenchmarkCreate4(b *testing.B) {
	benchmarkCreate(b, 10, 1000000)
}

func BenchmarkCreate5(b *testing.B) {
	benchmarkCreate(b, 1000, 1000000)
}
