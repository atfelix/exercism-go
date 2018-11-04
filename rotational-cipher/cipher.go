package rotationalcipher

import (
	"unicode"
	"bytes"
)

func RotationalCipher(text string, shift int) string {
	runeShift := rune(shift % 26)
	buffer := bytes.Buffer{}
	for _, r := range text {
		switch {
		case !unicode.IsLetter(r):
			buffer.WriteRune(r)
		case unicode.IsUpper(r):
			buffer.WriteRune((r - 'A' + runeShift) % 26 + 'A')
		case unicode.IsLower(r):
			buffer.WriteRune((r - 'a' + runeShift) % 26 + 'a')
		}
	}
	return buffer.String()
}