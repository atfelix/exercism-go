package romannumerals

import (
	"errors"
)

var (
	ones = [...]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens = [...]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hundreds = [...]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thousands = [...]string{"", "M", "MM", "MMM"}
)

func ToRomanNumeral(arabic int) (string, error) {
	if !isValid(arabic) {
		return "", errors.New("Invalid year")
	}
	return thousands[arabic / 1000] + hundreds[(arabic / 100) % 10] + tens[(arabic / 10) % 10] + ones[arabic % 10], nil
}

func isValid(arabic int) bool {
	return 0 < arabic && arabic <= 3000
}