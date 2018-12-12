package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Sleeper defines a sleeper behavior
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper defines a customizable sleeper
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep invokes sleep of configurable sleeper
func (cs *ConfigurableSleeper) Sleep() {
	cs.sleep(cs.duration)
}

// DefaultSleeper defins a default sleeper
type DefaultSleeper struct{}

// Sleep sleeps for one second
func (ds *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Countdown counts down from 3
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprintf(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{
		duration: 1 * time.Second,
		sleep:    time.Sleep,
	}
	Countdown(os.Stdout, sleeper)
}
