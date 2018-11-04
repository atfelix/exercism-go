package cryptosquare

import (
	"bytes"
	"strings"
	"unicode"
)

func Encode(s string) string {
	s = parsedString(s)
	slice := []string{}
	_, m := rectangleSize(s)
	for column := 0; column < m; column++ {
		slice = append(slice, stringFor(column, s))
	}
	return strings.Join(slice, " ")
}

func parsedString(s string) string {
	bytesArray := []byte{}
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			bytesArray = append(bytesArray, byte(unicode.ToLower(r)))
		}
	}
	return string(bytes.ToLower(bytesArray))
}

func rectangleSize(s string) (int, int) {
	length := len(s)
	root := squareRoot(length)
	if root * (root - 1) < length {
		return root, root
	}
	return root - 1, root
}

func stringFor(column int, s string) string {
	m, n := rectangleSize(s)
	bytesArray := []byte{}
	t := ""
	for row := 0; row < m; row++ {
		if n * row + column < len(s) {
			bytesArray = append(bytesArray,s[n * row + column])
			t += string(s[n * row + column])
		} else {
			bytesArray = append(bytesArray, byte(' '))
			t += " "
		}
	}
	return string(bytes.ToLower(bytesArray))
	return t
}

func squareRoot(n int) int {
	root := 1
	for root * root < n {
		root++
	}
	return root
}