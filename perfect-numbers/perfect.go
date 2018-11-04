package perfect

import (
	"errors"
)

type Classification int

var ErrOnlyPositive = errors.New("Only positive number are allowed")

const (
	ClassificationPerfect Classification = iota
	ClassificationAbundant
	ClassificationDeficient
)

func Classify(n int64) (Classification, error) {
	sumOfDivisors, err := divisorSum(n)
	if err != nil {
		return ClassificationAbundant, err
	}
	if sumOfDivisors == 2 * n {
		return ClassificationPerfect, nil
	} else if sumOfDivisors < 2 * n {
		return ClassificationDeficient, nil
	}
	return ClassificationAbundant, nil
}

func divisorSum(n int64) (int64, error) {
	if n <= 0 {
		return -1, ErrOnlyPositive
	}
	sumOfDivisors := int64(0)
	d := int64(1)
	for ; d * d < n; d++ {
		if n % d == 0 {
			sumOfDivisors += d + n / d
		}
	}
	if d * d == n {
		sumOfDivisors += d
	}
	return sumOfDivisors, nil
}