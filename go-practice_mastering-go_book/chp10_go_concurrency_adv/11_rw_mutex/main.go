package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type secret struct {
	password string
	m        sync.Mutex
	rwm      sync.RWMutex
}

var password = secret{password: "password"}

func change(c *secret, pass string) {
	c.rwm.Lock()
	fmt.Println("change password")
	time.Sleep(5 * time.Second)
	c.password = pass
	c.rwm.Unlock()
}

func showRLock(c *secret) string {
	c.rwm.RLock()
	fmt.Println("show with read lock")
	time.Sleep(3 * time.Second)
	defer c.rwm.RUnlock()
	return c.password
}

func showLock(c *secret) string {
	c.m.Lock()
	fmt.Print("show with lock")
	time.Sleep(3 * time.Second)
	defer c.m.Unlock()
	return c.password
}

func main() {
	var showFunc func(*secret) string

	if len(os.Args) != 2 {
		fmt.Println("using sync.RWMutex")
		showFunc = showRLock
	} else {
		fmt.Println("using sync.Mutex")
		showFunc = showLock
	}

	var wg sync.WaitGroup

	fmt.Println("pass:", showFunc(&password))

	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("pass:", showFunc(&password))
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		change(&password, "123456")
	}()
	wg.Wait()
	fmt.Println("pass:", showFunc(&password))
}
