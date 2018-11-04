package prime

func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}
	primes := []int{2, 3, 5, 7}
	for i := 11; len(primes) < n; i += 2 {
		if isPrime(i, primes) {
			primes = append(primes, i)
		}
	}
	return primes[n - 1], true
}

func isPrime(n int, primes []int) bool {
	for _, p := range primes {
		if n % p == 0 {
			return false
		} else if p * p > n {
			return true
		}
	}
	return false
}
