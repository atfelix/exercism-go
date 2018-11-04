package pangram

import (
	"unicode"
)

func IsPangram(s string) bool {
	m := map[rune]bool{}
	for _, char := range s {
		m[unicode.ToLower(char)] = true
	}
	for _, char := range "abcdefghijklmnopqrstuvwxyz" {
		if !m[char] {
			return false
		}
	}
	return true
}