package phonenumber

import (
	"fmt"
	"strings"
	"unicode"
)

type phoneNumber struct {
	countryCode, areaCode, exchangeCode, subscriberNumber string
}

func new(input string) (*phoneNumber, error) {
	if !isValid(input) {
		return nil, fmt.Errorf("Invalid input:  %s, %s", input, parse(input))
	}
	parsed := parse(input)
	var number *phoneNumber
	if len(parsed) == 10 {
		number = lengthTenPhoneNumber(parsed)
	} else {
		number = lengthElevenPhoneNumber(parsed)
	}

	if !number.isValid() {
		return nil, fmt.Errorf("Invalid phone number: %s", input)
	}
	return number, nil
}

func isValid(input string) bool {
	parsed := parse(input)
	return isValidLength(parsed) && isValidRuneSet(parsed)
}

var removalStrings = []string {
	"(",
	")",
	"-",
	"+",
	".",
	" ",
}

func parse(input string) string {
	for _, s := range removalStrings {
		input = strings.Replace(input, s, "", -1)
	}
	return input
}

func isValidLength(s string) bool {
	return 10 <= len(s) && len(s) <= 11
}

func isValidRuneSet(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func lengthTenPhoneNumber(input string) *phoneNumber {
	number := &phoneNumber {
		countryCode: "1",
	}
	number.areaCode = input[:3]
	number.exchangeCode = input[3:6]
	number.subscriberNumber = input[6:]
	return number
}

func lengthElevenPhoneNumber(input string) *phoneNumber {
	number := &phoneNumber{}
	number.countryCode = input[0:1]
	number.areaCode = input[1:4]
	number.exchangeCode = input[4:7]
	number.subscriberNumber = input[7:]
	return number
}

func (number *phoneNumber) isValid() bool {
	return number.countryCode == "1" && isValidAreaCode(number.areaCode) && isValidExchangeCode(number.exchangeCode)
}

func isValidAreaCode(input string) bool {
	return helper(input)
}

func isValidExchangeCode(input string) bool {
	return helper(input)
}

func helper(input string) bool {
	return input[0] != '0' && input[0] != '1'
}

func (number *phoneNumber) number() string {
	return fmt.Sprintf("%s%s%s", number.areaCode, number.exchangeCode, number.subscriberNumber)
}

func (number *phoneNumber) format() string {
	return fmt.Sprintf("(%s) %s-%s", number.areaCode, number.exchangeCode, number.subscriberNumber)
}

func (number *phoneNumber) areaCodeString() string {
	return fmt.Sprintf("%s", number.areaCode)
}

func Number(input string) (string, error) {
	return phoneHelper(input, (*phoneNumber).number)
}

func AreaCode(input string) (string, error) {
	return phoneHelper(input, (*phoneNumber).areaCodeString)
}

func Format(input string) (string, error) {
	return phoneHelper(input, (*phoneNumber).format)
}

func phoneHelper(input string, f func(*phoneNumber) string) (string, error) {
	number, err := new(input)
	if err != nil {
		return "", err
	}
	return f(number), nil
}
