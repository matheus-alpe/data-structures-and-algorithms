package twopointer

import (
	"unicode"
)

func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if !isAlphanumericByte(s[left]) {
			left++
			continue
		}

		if !isAlphanumericByte(s[right]) {
			right--
			continue
		}

		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--

	}

	return true
}

func isAlphanumericByte(c byte) bool {
	return unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c))
}

func toLower(c byte) rune {
	return unicode.ToLower(rune(c))
}
