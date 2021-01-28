package main

func reverse(s []int) {
	for i := 0; i <= len(s)-1; i++ {
		for j := 0; j < i; j++ {
			s[i], s[j] = s[j], s[i]
		}
	}
}
