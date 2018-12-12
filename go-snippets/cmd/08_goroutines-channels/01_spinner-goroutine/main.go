package main

import (
	"fmt"
	"time"
)

func main() {
	go func(delay time.Duration) {
		for {
			for _, v := range `-\|/` {
				fmt.Printf("\r%[1]c%[1]c%[1]c%[1]c%[1]c%[1]c", v)
				time.Sleep(delay)
			}
		}
	}(500 * time.Millisecond)

	time.Sleep(10 * time.Second)
}
