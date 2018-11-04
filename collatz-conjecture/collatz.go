package collatzconjecture

import (
	"fmt"
)

func CollatzConjecture(input int) (int, error) {
	switch {
	case input <= 0:
		return -1, fmt.Errorf("Invalid input:  input must be larger than 0")
	default:
		return collatz(input), nil
	}
}

func collatz(input int) int {
	switch {
		case input <= 0:
			panic("Invalid input:  input must be larger than 0")
		case input == 1:
			return 0
		case input % 2 == 1:
			return 1 + collatz(3 * input + 1)
		default:
			return 1 + collatz(input / 2)
	}
}