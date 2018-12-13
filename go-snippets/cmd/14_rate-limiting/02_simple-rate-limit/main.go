package main

import (
	"context"
	"log"
	"os"
	"sync"

	"golang.org/x/time/rate"
)

func main() {
	defer log.Printf("Done.")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	apiConnection := Open()
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ReadFiles(context.Background())
			if err != nil {
				log.Printf("cannot Resolve Address: %v", err)
			}
			log.Printf("ReadFile")
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("cannot Resolve Address: %v", err)
			}
			log.Printf("ResolveAddress")
		}()
	}
	wg.Wait()
}

// APIConnection defines an API connection struct
type APIConnection struct {
	rateLimiter *rate.Limiter
}

// Open represents a new API connection
func Open() *APIConnection {
	return &APIConnection{
		rateLimiter: rate.NewLimiter(rate.Limit(1), 1),
	}
}

// ReadFiles represents reading files
func (a *APIConnection) ReadFiles(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	// Pretend we do work here

	return nil
}

// ResolveAddress represents resolving address
func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return nil
	}
	// Pretend we do work here

	return nil
}
