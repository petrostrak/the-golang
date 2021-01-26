package main

func comma(s string) string {
	n := len(s)
	if n < 3 {
		return s // no comma is necessary
	}
	// calls itself recursively with a substring of all but
	// the last 3 characters, appends the comma and the last
	// 3 characters
	return comma(s[:n-3]) + "," + s[n-3:]
}
