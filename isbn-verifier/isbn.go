package isbn

import (
	"fmt"
	"strconv"
	"unicode"
	"bytes"
)

func IsValidISBN(isbn string) bool {
	isbn = filteredISBN(isbn)
	if !isValidFormat(isbn) {
		return false
	}

	sum := 0
	for index, multiplier := 0, 10; index < len(isbn); index, multiplier = index + 1, multiplier - 1 {
		value, err := convertByteToInt(isbn[index])
		if err != nil {
			return false
		}
		sum += multiplier * value
	}
	return sum % 11 == 0
}

func filteredISBN(isbn string) string {
	buffer := bytes.Buffer{}
	for _, r := range isbn {
		if unicode.IsDigit(r) || r == 'X' {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

func isValidFormat(isbn string) bool {
	if !isValidLength(isbn) {
		return false
	}
	length := len(isbn)
	lastRune := rune(isbn[length - 1])
	for _, r := range isbn[:length - 1] {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return unicode.IsDigit(lastRune) || lastRune == 'X'
}

func convertByteToInt(b byte) (int, error) {
	switch {
	case b == 'X':
		return 10, nil
	case unicode.IsDigit(rune(b)):
		value, err := strconv.Atoi(string(b))
		if err != nil {
			return -1, fmt.Errorf("There was an error reading the following supposed digit: %q", b)
		}
		return value, nil
	default:
		return -1, fmt.Errorf("Invalid byte sent: %q, expecting digit or 'X'", b)
	}
}

func isValidLength(isbn string) bool {
	return len(isbn) == 10
}
