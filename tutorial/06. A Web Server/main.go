package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// To avode race condition, meaning if two or more concurrent requests try to
// update /count at the same time, we must ensure that at most one goroutine
// accesses the variable at a time, which is the purpose of the mu.Lock() and
// mu.Unlock()
var mu sync.Mutex
var count int

// handler echoes the path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL path visited: %q", r.URL.Path)
}

// counter echoes the number of calls so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
