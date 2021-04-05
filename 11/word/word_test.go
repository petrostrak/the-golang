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

func TestFrenchPalindrome(t *testing.T) {
	if IsPalidrome("été") {
		t.Error(`IsPalindrom("été") = true`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if IsPalidrome(input) {
		t.Error(`IsPalindrom("été") = true`)
	}
}
