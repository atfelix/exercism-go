package anagram

import (
	"strings"
	"unicode"
)

func Detect(subject string, candidates []string) []string {
	anagrams := []string{}
	for _, candidate := range candidates {
		if isAnagram(subject, candidate) {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}

func isAnagram(s, t string) bool {
	return strings.ToLower(s) != strings.ToLower(t) && len(s) == len(t) && mapsEqual(counter(s), counter(t))
}

func mapsEqual(mapOne, mapTwo map[rune]int) bool {
	for key, value := range mapOne {
		mapTwo[key] -= value
		if mapTwo[key] != 0 {
			return false
		}
		delete(mapTwo, key)
	}
	return len(mapTwo) == 0
}

func counter(s string) map[rune]int {
	_counter := map[rune]int{}
	for _, r := range s {
		_counter[unicode.ToLower(r)]++
	}
	return _counter
}