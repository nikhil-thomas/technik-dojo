package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Served: %s\n", r.Host)
}

func timehandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "serving: %s\n", r.URL.Path)
	fmt.Printf("served time for: %s\n", r.Host)
}

func main() {
	port := ":8080"
	args := os.Args
	if len(args) > 1 {
		port = ":" + args[1]
	}

	// cannot use default mux for http profiling
	r := http.NewServeMux()
	r.HandleFunc("/time", timehandler)
	r.HandleFunc("/", myHandler)

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline/", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile/", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol/", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace/", pprof.Trace)

	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalln(err)
	}
}
