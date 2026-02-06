package twopointer

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"hello", false},
		{"", true},
		{"a", true},
		{"ab", false},
		{"madam", true},
		{"step on no pets", true},
	}

	for _, test := range tests {
		result := IsPalindrome(test.input)
		if result != test.expected {
			t.Errorf("IsPalindrome(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
