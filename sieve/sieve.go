package sieve

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}
	primes := []int{2}
	for potential := 3; potential <= limit; potential += 2 {
		for _, p := range primes {
			if p * p > potential {
				primes = append(primes, potential)
				break
			} else if potential % p == 0 {
				break
			}
		}
	}
	return primes
}