package allyourbase

import (
	"fmt"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	err := isValidInput(inputBase, outputBase, inputDigits)
	if err != nil {
		return nil, err
	}
	value := convertToInt(inputDigits, inputBase)
	return convertToDigits(value, outputBase), nil
}

func isValidInput(inputBase, outputBase int, inputDigits []int) error {
	switch {
	case inputBase < 2: return fmt.Errorf("input base must be >= 2")
	case outputBase < 2: return fmt.Errorf("output base must be >= 2")
	case !areValidDigits(inputDigits, inputBase): return fmt.Errorf("all digits must satisfy 0 <= d < input base")
	default: return nil
	}
}

func areValidDigits(digits []int, base int) bool {
	for _, digit := range digits {
		if digit < 0 || base <= digit {
			return false
		}
	}
	return true
}

func convertToDigits(value, base int) []int {
	log := intLog(value, base)
	digits := make([]int, log)

	for value > 0 {
		digits[log - 1] = value % base
		log--
		value /= base
	}

	return digits
}

func convertToInt(digits []int, base int) int {
	value := 0
	for _, digit := range digits {
		value = base * value + digit
	}
	return value
}

func intLog(value, base int) int {
	log := 1
	value /= base
	for value > 0 {
		log++
		value /= base
	}
	return log
}