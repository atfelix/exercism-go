package prime

import (
	"math/rand"
	"sort"
)

func power(base, exponent int64) int64 {
	result, multiplier := int64(1), base
	for exponent > 0 {
		if exponent % 2 == 1 {
			result *= multiplier
		}
		multiplier *= multiplier
		exponent >>= 1
	}
	return result	
}

func powerMod(base, exponent, modulus int64) int64 {
	result, multiplier := int64(1), base
	for exponent > 0 {
		if exponent % 2 == 1 {
			result = (result * multiplier) % modulus
		}
		multiplier = (multiplier * multiplier) % modulus
		exponent >>= 1
	}
	return result
}

func isComposite(base, oddPart, modulus, numberOfTries int64) bool {
	if powerMod(base, oddPart, modulus) == 1 {
		return false
	}
	for i := int64(0); i < numberOfTries; i++ {
		if powerMod(base, power(2, i) * oddPart, modulus) == modulus - 1 {
			return false
		}
	}
	return true
}

func powerOf2(n int64) int64 {
	power := int64(0)
	for n % 2 == 0 {
		power++
		n >>= 1
	}
	return power
}

func millerRabin(n int64) bool {
	if n == 1 {
		return false
	}
	if n == 2 || n == 3 || n == 5 || n == 7 {
		return true
	}
	if n % 2 == 0 {
		return false
	}
	oddPart := n - 1
	powerOfTwo := powerOf2(oddPart)
	oddPart /= powerOfTwo
	for _, base := range []int64{int64(2), int64(3), int64(5), int64(7)} {
		if isComposite(base, oddPart, n, powerOfTwo) {
			return false
		}
	}
	return true
}

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func pollardRho(n int64) int64 {
	if n % 2 == 0 {
		return int64(2)
	}
	b := rand.Intn(200) - 100
	for b == -2 {
		b = rand.Intn(200) - 100
	}
	base := int64(b)
	x := int64(rand.Intn(99)) + 1
	y := x
	d := int64(1)
	for d == 1 {
		x = (powerMod(x, 2, n) + base) % n
		y = (powerMod(y, 2, n) + base) % n
		y = (powerMod(y, 2, n) + base) % n
		d = gcd(abs(x - y), n)
		if d == n {
			break
		}
	}
	return d
}

func rhoFactor(n int64) []int64 {
	if n == 1 {
		return []int64{}
	}
	divisors := []int64{n}
	factors := []int64{}
	count := 0
	for len(divisors) != 0 && count < 20 {
		count++
		divisor := divisors[len(divisors) - 1]
		divisors = divisors[:len(divisors) - 1]
		if millerRabin(divisor) {
			factors = append(factors, divisor)
			continue
		}
		newDivisor := pollardRho(divisor)
		if divisor == newDivisor {
			divisors = append(divisors, divisor)
			continue
		}
		divisors = append(divisors, []int64{divisor / newDivisor, newDivisor}...)
	}
	return factors
}

func Factors(n int64) []int64 {
	factors := rhoFactor(n)
	sort.Slice(factors, func(i, j int) bool {
		return factors[i] < factors[j]
	})
	return factors
}
