package wordcount

import (
	"bufio"
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	frequency := Frequency{}
	phrase = parse(phrase)
	scanner := bufio.NewScanner(strings.NewReader(phrase))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		frequency[scanner.Text()]++
	}
	return frequency
}

func parse(phrase string) string {
	phrase = strings.Replace(phrase, "' ", " ", -1)
	phrase = strings.Replace(phrase, " '", " ", -1)
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'' {
			return unicode.ToLower(r)
		}
		return ' '
	}, phrase)
}