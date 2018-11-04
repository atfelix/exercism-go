package triangle

import "math"

type Kind int

const (
    NaT Kind = iota
    Equ
    Iso
    Sca
)

func KindFromSides(a, b, c float64) Kind {
	a, b, c = orderedTriple(a, b, c)
	if !isTriange(a, b, c) {
		return NaT
	} else if isEquilateral(a, b, c) {
		return Equ
	} else if isIsoceles(a, b, c) {
		return Iso
	} else {
		return Sca
	}
}

func orderedTriple(a, b, c float64) (float64, float64, float64) {
	return min(a, b, c), middle(a, b, c), max(a, b, c)
}

func isTriange(a, b, c float64) bool {
	a, b, c = orderedTriple(a, b, c)
	return 0 < a && a + b >= c
}

func isEquilateral(a, b, c float64) bool {
	return a == b && b == c
}

func isIsoceles(a, b, c float64) bool {
	return a == b && b < c || a < b && b == c
}

func max(a, b, c float64) float64 {
	return math.Max(a, math.Max(b, c))
}

func min(a, b, c float64) float64 {
	return math.Min(a, math.Min(b, c))
}

func middle(a, b, c float64) float64 {
	return a + b + c - max(a, b, c) - min(a, b, c)
}