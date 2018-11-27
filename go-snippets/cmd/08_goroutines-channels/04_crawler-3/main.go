package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nikhil-thomas/technik-dojo/go-snippets/cmd/08_goroutines-channels/02_crawler-1/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	links, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return links
}

func main() {
	worklist := make(chan []string)
	unseenlinks := make(chan string)

	go func() {
		worklist <- os.Args[1:]
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenlinks {
				foundlinks := crawl(link)
				go func() {
					worklist <- foundlinks
				}()
			}
		}()
	}

	seen := make(map[string]bool)

	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenlinks <- link
			}
		}
	}
}
