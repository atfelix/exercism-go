package railfence

import (
	"sort"
)

func Encode(input string, rails int) string {
	output := []byte{}
	for index := 0; index < len(input); index++ {
		output = append(output, ' ')
	}
	currentIndex := 0
	for railIndex := 0; railIndex < rails; railIndex++ {
		for _, mod := range moduliFor(railIndex, rails, len(input)) {
			output[currentIndex] = input[mod]
			currentIndex++
		}
	}
	return string(output)
}

func Decode(input string, rails int) string {
	output := []byte{}
	for index := 0; index < len(input); index++ {
		output = append(output, ' ')
	}
	currentIndex := 0
	for railIndex := 0; railIndex < rails; railIndex++ {
		for _, mod := range moduliFor(railIndex, rails, len(input)) {
			output[mod] = input[currentIndex]
			currentIndex++
		}
	}
	return string(output)
}

func moduliFor(rail, outOf, length int) []int {
	if rail == 0 || rail == outOf - 1 {
		return indicesFor(rail, 2 * outOf - 2, length)
	}
	slice := indicesFor(rail, 2 * outOf - 2, length)
	slice = append(slice, indicesFor(2 * outOf - 2 - rail, 2 * outOf - 2, length)...)
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice
}

func indicesFor(start, mod, length int) []int {
	slice := []int{}
	for index := start; index < length; index += mod {
		slice = append(slice, index)
	}
	return slice
}