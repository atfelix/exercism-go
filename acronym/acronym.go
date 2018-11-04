package acronym

import (
	"unicode"
	"strings"
)

func Abbreviate(s string) string {
	return strings.Map(func(r rune) rune {
		return unicode.ToUpper(r)
	}, parse(s))
}

func parse(s string) string {
	s = strings.Replace(s, "-", " ", -1)
	s = strings.Trim(s, " ")
	t := ""
	for _, ss := range strings.Split(s, " ") {
		t += string(ss[0])
	}
	return t
}