package webchecker

import (
	"errors"
	"net/http"
	"time"
)

// Racer compares fetch speed of two urls
func Racer(url1, url2 string) string {

	duration1 := measureTime(url1)
	duration2 := measureTime(url2)

	if duration1 < duration2 {
		return url1
	}
	return url2
}

func measureTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	duration := time.Since(start)
	return duration
}

// RacerPing cmpares fetch speed of two urls
func RacerPing(url1, url2 string) (string, error) {

	pingChan1 := ping(url1)
	pingChan2 := ping(url2)

	timeout := time.After(5 * time.Millisecond)

	select {
	case <-timeout:
		return "", errors.New("raceping timeout")
	case <-pingChan1:
		return url1, nil
	case <-pingChan2:
		return url2, nil
	}
}

func ping(url string) <-chan bool {
	pChan := make(chan bool)
	go func() {
		http.Get(url)
		pChan <- true
		close(pChan)
	}()
	return pChan
}

var tenSecondTimeout = 10 * time.Second

// Racer10s uses a Racer with 10s timeout
func Racer10s(url1, url2 string) (string, error) {
	return ConfigurableRacer(url1, url2, tenSecondTimeout)
}

// ConfigurableRacer is a Racer whose timeout can be configured
func ConfigurableRacer(url1, url2 string, t time.Duration) (string, error) {
	pingChan1 := ping(url1)
	pingChan2 := ping(url2)

	timeout := time.After(t)

	select {
	case <-timeout:
		return "", errors.New("raceping timeout")
	case <-pingChan1:
		return url1, nil
	case <-pingChan2:
		return url2, nil
	}
}
