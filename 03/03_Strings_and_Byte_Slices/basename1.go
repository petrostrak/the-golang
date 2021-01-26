package main

func basenameSlash(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:] // +1 emits the '/' and ':' keeps the rest of the byteslice
			break
		}
	}
	return s
}

func basenameDot(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i] // preserves everything before '.'
			break
		}
	}
	return s
}
