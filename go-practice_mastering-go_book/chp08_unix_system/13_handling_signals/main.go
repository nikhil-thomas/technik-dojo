package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func handleSignal(chanName string, signal os.Signal) {
	fmt.Println(chanName, "Caught:", signal)
}

func main() {
	wg := sync.WaitGroup{}

	sigsChan1 := make(chan os.Signal, 1)
	sigsChan2 := make(chan os.Signal, 1)
	signal.Notify(sigsChan1, os.Interrupt, syscall.SIGTERM)
	signal.Notify(sigsChan2)

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			sig := <-sigsChan1
			switch sig {
			case os.Interrupt:
				fmt.Println("Chan1 caught:", sig)
			case syscall.SIGTERM:
				handleSignal("chan1", sig)
				return
			}
		}
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			sig := <-sigsChan2
			switch sig {
			case os.Interrupt:
				handleSignal("chan2", sig)
			case syscall.SIGTERM:
				handleSignal("chan2", sig)
				fmt.Println("exiting")
				os.Exit(0)
			case syscall.SIGUSR2:
				fmt.Println("handling syscall.SIGUSR2")
			}
		}
	}()

	wg.Wait()
}
