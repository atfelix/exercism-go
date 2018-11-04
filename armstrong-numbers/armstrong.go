package armstrong

func power(base, exponent int) int {
	if exponent == 0 {
		return 1
	}
	return base * power(base, exponent - 1)
}

func lengthOf(n, base int) int {
	count := 0
	for n > 0 {
		count ++
		n /= base
	}
	return count
}

func IsNumber(n int) bool {
	armstrong, length, copy := 0, lengthOf(n, 10), n
	for n > 0 {
		digit := n % 10
		armstrong += power(digit, length)
		n /= 10
	}
	return armstrong == copy
}