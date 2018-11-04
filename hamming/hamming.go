package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Lengths of DNA strands don't match")
	}

	count := 0

	for index := 0; index < len(a); index++ {
		if a[index] != b[index] {
			count++
		}
	}
	return count, nil
}
