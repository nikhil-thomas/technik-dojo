package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nikhil-thomas/technik-dojo/go-snippets/cmd/07_go-prg-lng_c5/12_links-3/links"
)

var depth = 5
var width = 2

func breadthFirst(f func(item string) []string, urlList []string) []string {

	seen := make(map[string]bool)
	for len(urlList) > 0 {
		if depth <= 0 {
			break
		}
		fmt.Println("depth:", depth)
		depth--

		items := urlList
		urlList = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true

				newURLs := f(item)
				i := 0
				localSeen := map[string]bool{}
				for _, newURL := range newURLs {
					if i >= width {
						break
					}
					if !seen[newURL] {
						if !localSeen[newURL] {
							localSeen[newURL] = true
							fmt.Println(newURL)
							urlList = append(urlList, newURL)
							i++
						}
					}
				}

			}
		}
	}

	return urlList
}

func crawl(url string) []string {
	//fmt.Println(url)
	urls, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return urls
}

func main() {
	list := os.Args[1:]
	list = breadthFirst(crawl, list)
	// list, _ := links.Extract(os.Args[1])
	for _, url := range list {
		fmt.Println(url)
	}
}
