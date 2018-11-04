package luhn

import "strings"

func Valid(cardNumber string) bool {
	cardNumber = strings.Replace(cardNumber, " ", "", -1)

	if len(cardNumber) <= 1 {
		return false
	}
	count := 0

	for index, r := range cardNumber {
		digit := int(r - '0')
		if index % 2 != len(cardNumber) % 2 {
			count += digit
		} else if 2 * digit <= 9 {
			count += 2 * digit
		} else {
			count += 2 * digit - 9
		}
	}
	return count % 10 == 0
}