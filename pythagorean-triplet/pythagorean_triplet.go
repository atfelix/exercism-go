package pythagorean

type Triplet [3]int

func Range(min, max int) []Triplet {
	slice := []Triplet{}
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			sumOfSquares := a * a + b * b
			root := squareRoot(sumOfSquares)
			if min <= root && root <= max && root * root == sumOfSquares {
				slice = append(slice, Triplet{a, b, root})
			}
		}
	}
	return slice
}

func isSquare(n int) bool {
	root := squareRoot(n)
	return root * root == n
}

func squareRoot(n int) int {
	root := 1
	for root * root < n {
		root++
	}
	return root
}

func Sum(p int) []Triplet {
	slice := []Triplet{}
	for a := 1; a <= p; a++ {
		for b := a; b <= p - a; b++ {
			if a * a + b * b == (p - a - b) * (p - a - b) {
				slice = append(slice, Triplet{a, b, p - a - b})
			}
		}
	}
	return slice
}