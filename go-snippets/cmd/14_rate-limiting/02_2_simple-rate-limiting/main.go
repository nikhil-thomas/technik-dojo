package main

import (
    "context"
    "golang.org/x/time/rate"
    "log"
    "os"
    "sync"
)

type APIConnection struct {
    rateLimiter *rate.Limiter
}

func Open() *APIConnection {
    return &APIConnection{
	    rateLimiter: rate.NewLimiter(rate.Limit(1), 4),
    }
}

func (a *APIConnection) ReadFile(ctx context.Context) error {
    if err := a.rateLimiter.Wait(ctx); err != nil {
	return err
    }
    // do work
    return nil
}

func (a *APIConnection) ResolveAddress(ctx context.Context) error {
    if err := a.rateLimiter.Wait(ctx); err != nil {
        return err
    }
    // do work
    return nil
}



func main() {
    defer log.Printf("Main done")
    log.SetOutput(os.Stdout)
    log.SetFlags(log.Ltime | log.LUTC)

    apiConnection := Open()
    var wg sync.WaitGroup
    wg.Add(40)

    for i := 0; i < 20; i++ {
        go func() {
	    defer wg.Done()
	    err := apiConnection.ReadFile(context.Background())
	    if err != nil {
		log.Printf("cannot ReadFile: %v", err)
            }
	    log.Printf("ReadFile")
	}()
    }


    for i := 0; i < 20; i++ {
	go func() {
	   defer wg.Done()
	   err := apiConnection.ResolveAddress(context.Background())
	   if err != nil {
		log.Printf("cannot ResolveAddress: %v", err)
	   }
	   log.Printf("Resolve Address")
	}()
    }
    wg.Wait()
}

