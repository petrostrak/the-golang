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

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false},
		{"desserts", false},
	}
	for _, test := range tests {
		if got := IsPalidrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}
