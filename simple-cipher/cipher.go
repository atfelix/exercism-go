package cipher

import (
	"bytes"
	"unicode"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type vigenere struct {
	key string
}

func NewCaesar() Cipher {
	return NewVigenere("d")
}

func NewShift(shift int) Cipher {
	if shift <= 0 {
		return NewVigenere(string('a' + shift + 26))
	}
	return NewVigenere(string('a' + shift))
}

func NewVigenere(key string) Cipher {
	if isValidVigenere(key) {
		return vigenere{key: key}
	}
	return nil
}

func isValidVigenere(key string) bool {
	for _, r := range key {
		if !unicode.IsLetter(r) || !unicode.IsLower(r) {
			return false
		}
	}
	return len(key) != 0 && !allA(key)
}

func allA(key string) bool {
	for _, r := range key {
		if r != 'a' {
			return false
		}
	}
	return true
}

func (cipher vigenere) Encode(input string) string {
	return cipher.shifted(input, 1)
}

func (cipher vigenere) Decode(input string) string {
	return cipher.shifted(input, -1)
}

func (cipher vigenere) shifted(input string, direction rune) string {
	buffer := bytes.Buffer{}

	index := 0
	for _, r := range input {
		if unicode.IsLetter(r) {
			r = unicode.ToLower(r)
			shift := rune(cipher.key[index % len(cipher.key)]) - 'a'
			index++
			buffer.WriteRune(((r - 'a') + direction * shift + 26) % 26 + 'a')
		}
	}

	return buffer.String()	
}
