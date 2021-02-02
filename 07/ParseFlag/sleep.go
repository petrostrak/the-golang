package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", time.Second, "sleep period")

// go run sleep.go -period 2m50s
func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
