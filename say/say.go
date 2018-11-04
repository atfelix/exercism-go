package say

import (
	"strings"
)

var (
	lessThanTwenty = map[int64]string {
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
	}
	tens = map[int64]string {
		2: "twenty",
		3: "thirty",
		4: "forty",
		5: "fifty",
		6: "sixty",
		7: "seventy",
		8: "eighty",
		9: "ninety",
	}
)

func Say(input int64) (string, bool) {
	if input < 0 || input >= 1000000000000 {
		return "", false
	}
	if input == 0 {
		return "zero", true
	}
	words := []string{}
	words = add(words, input, "billion", billionFrom)
	words = add(words, input, "million", millionFrom)
	words = add(words, input, "thousand", thousandFrom)
	words = add(words, input, "", hundredFrom)
	return strings.Join(words, " "), true
}

func add(words []string, input int64, descriptor string, groupBy func(int64) int64) []string {
	group, ok := threeDigitNumber(groupBy(input))
	words = addHelper(words, group, descriptor, ok)
	return words
}

func addHelper(words, newWords []string, descriptor string, ok bool) []string {
	if ok {
		words = append(words, newWords...)
		if descriptor != "" {
		words = append(words, descriptor)
		}
	}
	return words
}

func billionFrom(number int64) int64 {
	return number % 1000000000000 / 1000000000
}

func millionFrom(number int64) int64 {
	return number % 1000000000 / 1000000
}

func thousandFrom(number int64) int64 {
	return number % 1000000 / 1000
}

func hundredFrom(number int64) int64 {
	return number % 1000
}

func threeDigitNumber(input int64) ([]string, bool) {
	if input == 0 {
		return []string{}, false
	}
	if input < 100 {
		return twoDigitNumber(input)
	}
	words := []string{lessThanTwenty[input / 100], "hundred"}
	if twoDigit, ok := twoDigitNumber(input % 100); ok {
		words = append(words, twoDigit...)
	}
	return words, true
}

func twoDigitNumber(input int64) ([]string, bool) {
	if input == 0 {
		return []string{}, false
	}
	if input < 20 {
		return []string{lessThanTwenty[input]}, true
	}
	tensPart := tens[input / 10]
	onesPart := lessThanTwenty[input % 10]
	connector := "-"
	if onesPart == "" {
		connector = ""
	}
	return []string{tensPart + connector + onesPart}, true
}