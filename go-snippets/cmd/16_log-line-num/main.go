package main

import (
    "log"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s served", r.URL.Path)
    w.Write([]byte("hello"))
}

func logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        start := time.Now()
	next.ServeHTTP(w, r)
        log.Printf(
		"%s\t\t%s\t\t%s",
		r.Method,
		r.RequestURI,
		time.Since(start),
	)
    })
}

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.Print("123")
    http.Handle("/", logger(http.HandlerFunc(handler)))
    http.ListenAndServe(":8080", nil)
}

