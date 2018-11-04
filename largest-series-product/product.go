package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(s string, length int) (int64, error) {
	if length > len(s) {
		return -1, errors.New("span exceeds string length")
	}
	if length < 0 {
		return -1, errors.New("length must be non-negative")
	}
	product := int64(-1)
	for start := 0; start + length <= len(s); start++ {
		currentProduct := int64(1)
		for _, r := range s[start:start + length] {
			digit, err := strconv.ParseInt(string(r), 10, 64)
			if err != nil {
				return -1, err
			}
			currentProduct *= digit
		}
		if currentProduct > product {
			product = currentProduct
		}
	}
	return product, nil
}