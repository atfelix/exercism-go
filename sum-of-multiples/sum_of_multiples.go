package summultiples

func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	for index, value := range sieve(limit, divisors...) {
		if value {
			sum += index
		}
	}
	return sum
}

func sieve(limit int, divisors ...int) []bool {
	cache := make([]bool, limit)
	for _, divisor := range divisors {
		for n := divisor; n < limit; n += divisor {
			cache[n] = true
		}
	}
	return cache
}