package main

func nonEmpty(str []string) []string {
	i := 0
	for _, s := range str {
		if s != "" {
			str[i] = s
			i++
		}
	}
	return str[:i]
}

func nonEmptyAlt(str []string) []string {
	out := str[:0] // creating an zero-length slice of original
	for _, s := range str {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
