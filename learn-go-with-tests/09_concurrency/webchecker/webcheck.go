package webchecker

import (
	"sync"
)

// WebsiteChecker defines a website chacker function type
type WebsiteChecker func(string) bool

// CheckWebsites checks a set of urls
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			mu.Lock()
			results[u] = wc(u)
			mu.Unlock()
		}(url)
	}
	wg.Wait()
	return results
}

type result struct {
	url  string
	pass bool
}

// CheckWebsitesWithChan checks a set of urls
func CheckWebsitesWithChan(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChan := make(chan result)
	for _, url := range urls {
		go func(u string) {
			resultChan <- result{
				url:  u,
				pass: wc(u),
			}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChan
		results[r.url] = r.pass
	}

	return results
}
