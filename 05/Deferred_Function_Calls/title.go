package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}
	return nil
}

var mu sync.Mutex
var m = make(map[string]int)

func lockup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}
