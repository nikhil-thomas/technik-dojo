package main

import (
	"context"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	log.Println("main started")
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go action(ctx)
	time.Sleep(1 * time.Second)
	cancel()
	// calling cancel again will not cause a panic
	// cancel checks if it is already canceled,
	// if yes it immediately returns
	cancel()
	wg.Wait()
	log.Println("main completed")
}

func action(ctx context.Context) {
	defer wg.Done()
	log.Println("action started")
	select {
	case <-time.After(3 * time.Second):
		log.Println("action success")
	case <-ctx.Done():
		log.Println("action aborted")
	}
	log.Println("action end")
}
