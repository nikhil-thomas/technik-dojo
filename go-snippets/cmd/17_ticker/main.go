package main

import (
    "time"
    "log"

)

func main() {
    ticker := time.NewTicker(500 * time.Millisecond)
    go func() {
	for t := range ticker.C {
            log.Println("Tick at", t)
	}
	log.Println("End of loop")
    }()
    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    log.Println("Ticker stopped")
    time.Sleep(5 * time.Second)
}

