package word

import "testing"

func TestPalindrome(t *testing.T) {
	if !IsPalidrome("detartrated") {
		t.Error(`IsPalindrom("detartrated") = false`)
	}
	if !IsPalidrome("kayak") {
		t.Error(`IsPalindrom("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalidrome("palindrome") {
		t.Error(`IsPalindrom("palindrome") = true`)
	}
}
