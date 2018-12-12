package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlaystore())

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalln(err)
	}
}
