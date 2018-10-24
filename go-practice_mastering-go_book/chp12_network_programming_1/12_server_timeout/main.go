package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "serving: %s\n", r.URL.Path)
	fmt.Printf("server: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	body := "the current time is"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>", t)
	fmt.Fprintf(w, "serving %s\n", r.URL.Path)
	fmt.Printf("served time for: %s\n", r.Host)
}

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", handler1)
	m.HandleFunc("/time", timeHandler)

	srv := http.Server{
		Addr:         ":8002",
		Handler:      m,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}

}
