package main

import (
	"fmt"
	"log"
	"net/http"
)

func checkStatusOk(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Status: OK")
}

func statusNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "serving: %s\n", r.URL.Path)
	fmt.Printf("served: %s\n", r.Host)
}

func main() {
	http.HandleFunc("/status", checkStatusOk)
	http.HandleFunc("/notfound", statusNotFound)
	http.HandleFunc("/", myHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
