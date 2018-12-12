package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nikhil-thomas/technik-dojo/go-snippets/cmd/08_goroutines-channels/02_crawler-1/links"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	links, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return links
}

func main() {
	// number of pending sends to worklist
	var n int
	worklist := make(chan []string)

	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(links string) {
					newLinks := crawl(link)
					worklist <- newLinks
				}(link)

			}
		}
	}
}
