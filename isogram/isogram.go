package isogram

import "strings"

func IsIsogram(phrase string) bool {
	phrase = strings.Replace(phrase, " ", "", -1)
	phrase = strings.Replace(phrase, "-", "", -1)
	phrase = strings.ToLower(phrase)
	setOfLetters := map[rune]bool{}

	for _, r := range phrase {
		if setOfLetters[r] {
			return false
		}
		setOfLetters[r] = true
	}

	return true
}