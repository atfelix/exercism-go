package palindrome

import (
	"errors"
)

type Product struct {
	Product int
	Factorizations [][2]int
}

func Products(min, max int) (Product, Product, error) {
	if min > max {
		return Product{}, Product{}, errors.New("fmin > fmax...")
	}
	_min, _max, factorizations, err := palindromeFactorizations(min, max)
	if err != nil {
		return Product{}, Product{}, err
	}
	return Product{_min, factorizations[_min]}, Product{_max, factorizations[_max]}, nil
}

func palindromeFactorizations(min, max int) (int, int, map[int][][2]int, error) {
	factorizations := map[int][][2]int{}
	_max, _min := min * min, max * max
	if _min < _max {
		_min, _max = _max, _min
	}
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			n := a * b
			if !isPalindrome(n) {
				continue
			}
			if n < _min {
				_min = n
			}
			if n > _max {
				_max = n
			}
			if _, ok := factorizations[n]; !ok {
				factorizations[n] = [][2]int{}
			}
			factorizations[n] = append(factorizations[n], [2]int{a, b})
		}
	}
	if len(factorizations) == 0 {
		return -1, -1, nil, errors.New("no palindromes...")
	}
	return _min, _max, factorizations, nil
}

func isPalindrome(n int) bool {
	reverse := 0
	if n < 0 {
		n = -n
	}
	copy := n
	for copy != 0 {
		reverse = reverse * 10 + copy % 10
		copy /= 10
	}
	return n == reverse
}

func divisorPairs(n, min, max int) [][2]int {
	pairs := [][2]int{}
	for a := min; a <= max && a <= n / a; a++ {
		if n % a == 0 && min <= n / a && n / a <= max {
			pairs = append(pairs, [2]int{a, n / a})
		}
	}
	return pairs
}