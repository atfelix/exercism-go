package atbash

import (
	"strings"
	"unicode"
)

var (
	forwards = "abcdefghijklmnopqrstuvwxyz"
	backwards = "zyxwvutsrqponmlkjihgfedcba"
)

func Atbash(s string) string {
	return strings.Join(atbashSlice(parse(s)), " ")
}

func atbashSlice(s string) []string {
	slice := []string{}
	for index := 0; index < len(s); index += 5 {
		endIndex := index + 5
		if endIndex > len(s) {
			endIndex = len(s)
		}
		slice = append(slice, atbashString(s[index:endIndex]))
	}
	return slice
}

func atbashString(s string) string {
	slice := []rune(s)
	for currentIndex, r := range slice {
		if index := firstIndex(forwards, unicode.ToLower(r)); index != -1 {
			slice[currentIndex] = rune(backwards[index])
		}
	}
	return string(slice)
}

func parse(s string) string {
	t := ""
	for _, r := range s {
		if !(unicode.IsPunct(r) || unicode.IsSpace(r)) {
			t += string(unicode.ToLower(r))
		}
	}
	return t
}

func firstIndex(s string, r rune) int {
	for index, rr := range s {
		if rr == r {
			return index
		}
	}
	return -1
}