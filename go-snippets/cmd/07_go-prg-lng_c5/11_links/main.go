package main

import (
	"log"
	"os"

	"github.com/nikhil-thomas/technik-dojo/go-snippets/cmd/07_go-prg-lng_c5/11_links/links"
)

func main() {
	for _, url := range os.Args[1:] {
		l, err := links.Extract(url)
		if err != nil {
			log.Println(err)
			continue
		}
		for _, link := range l {
			log.Println(link)
		}
	}
}
