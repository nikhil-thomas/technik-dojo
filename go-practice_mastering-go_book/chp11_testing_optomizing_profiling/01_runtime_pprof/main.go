package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func fibo1(n int) int64 {
	if n == 0 || n == 1 {
		return int64(n)
	}
	time.Sleep(1 * time.Millisecond)
	return int64(fibo1(n-1)) + int64(fibo1(n-2))
}

func fibo2(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int

		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	time.Sleep(50 * time.Millisecond)
	return fn[n]
}

func prime1(n int) bool {
	k := math.Floor(float64(n/2 + 1))
	for i := 2; i < int(k); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func prime2(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	cpuProfileFile, err := os.Create("/tmp/cpuprofile.out")
	if err != nil {
		log.Fatalln(err)
	}
	defer cpuProfileFile.Close()

	pprof.StartCPUProfile(cpuProfileFile)
	defer pprof.StopCPUProfile()

	total := 0
	for i := 2; i < 100000; i++ {
		n := prime1(i)
		if n {
			total = total + 1
		}
	}

	fmt.Println("prime1() -> Total primes:", total)

	total = 0
	for i := 2; i < 100000; i++ {
		n := prime2(i)
		if n {
			total = total + 1
		}
	}

	fmt.Println("prime2() -> Total primes:", total)

	for i := 1; i < 20; i++ {
		n := fibo1(i)
		fmt.Print(n, " ")
	}
	fmt.Print("\n")

	for i := 1; i < 20; i++ {
		n := fibo2(i)
		fmt.Print(n, " ")
	}
	fmt.Print("\n")

	runtime.GC()

	// memory profiling
	memoryProfileFile, err := os.Create("/tmp/memoryprofile.out")
	if err != nil {
		log.Fatalln(err)
	}
	defer memoryProfileFile.Close()

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			log.Fatalln("operation failed")
		}
		time.Sleep(50 * time.Millisecond)
	}

	err = pprof.WriteHeapProfile(memoryProfileFile)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("end of main")
}
