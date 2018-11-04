package yacht

import (
	"sort"
	"fmt"
)

type dice [5]int

// helper methods

func all(cup dice, f func(int, int) bool) bool {
	for index, element := range cup {
		if !f(index, element) {
			return false
		}
	}
	return true
}

func any(cup dice, f func(int, int) bool) bool {
	return !all(cup, func (x, y int) bool {
		return !f(x, y)
	})
}

func reduce(cup dice, initial int, f func(int, int) int) int {
	acc := initial
	for _, element := range cup {
		acc = f(acc, element)
	}
	return acc
}

func filteredSum(cup dice, weight int) int {
	sum := 0
	for _, element := range cup {
		if element == weight {
			sum += weight
		}
	}
	return sum
}

func sum(cup dice) int {
	sum := 0
	for _, element := range cup {
		sum += element
	}
	return sum
}

func counter(cup dice) map[int]int {
	counter := map[int]int{}
	for _, die := range cup {
		counter[die]++
	}
	return counter
}

// dice predicates

func trueAlways(cup dice) bool {
	return true
}

func isYacht(cup dice) bool {
	return all(cup, func(ignored, die int) bool {
		return die == cup[0]
	})
}

func isFullHouse(cup dice) bool {
	counter := counter(cup)
	return len(counter) == 2 && all(cup, func(ignored, die int) bool {
		return counter[die] > 1
	})
}

func isFourOfAKind(cup dice) bool {
	counter := counter(cup)
	return any(cup, func(ignored, die int) bool {
		return counter[die] > 3
	})
}

func isLittleStraight(cup dice) bool {
	sort.Ints(cup[:])
	return cup == [...]int{1,2,3,4,5}
}

func isBigStraight(cup dice) bool {
	sort.Ints(cup[:])
	return cup == [...]int{2,3,4,5,6}
}

// dice scores

func fourOfAKindScore(cup dice) int {
	return sum(cup) - reduce(cup, 0, func(x, y int) int {
		return x ^ y
	})
}

func yachtScore(cup dice) int {
	return 50
}

func straightScore(cup dice) int {
	return 30
}

// calculator functions

type yachtCalculator struct {
	score func(dice) int
	predicate func(dice) bool
}

func newCategory(category string) yachtCalculator {
	switch category {
	case "yacht":
		return yachtCalculator {
			predicate: isYacht,
			score: yachtScore,
		}
	case "ones":
		return filteredSumYachtCalculator(1)
	case "twos":
		return filteredSumYachtCalculator(2)
	case "threes":
		return filteredSumYachtCalculator(3)
	case "fours":
		return filteredSumYachtCalculator(4)
	case "fives":
		return filteredSumYachtCalculator(5)
	case "sixes":
		return filteredSumYachtCalculator(6)
	case "full house":
		return yachtCalculator {
			predicate: isFullHouse,
			score: sum,
		}
	case "four of a kind":
		return yachtCalculator {
			predicate: isFourOfAKind,
			score: fourOfAKindScore,
		}
	case "little straight":
		return yachtCalculator {
			predicate: isLittleStraight,
			score: straightScore,
		}
	case "big straight":
		return yachtCalculator {
			predicate: isBigStraight,
			score: straightScore,
		}
	case "choice":
		return yachtCalculator {
			predicate: trueAlways,
			score: sum,
		}
	default:
		panic(fmt.Sprintf("Invalid category: %s", category))
	}
}

func (calculator yachtCalculator) scoreFor(cup []int) int {
	if len(cup) != 5 {
		return 0
	}

	var array dice
	copy(array[:5], cup)

	if !calculator.predicate(array) {
		return 0
	}
	return calculator.score(array)
}

func filteredSumYachtCalculator(weight int) yachtCalculator {
	return yachtCalculator {
		predicate: func (cup dice) bool { return true },
		score: func(cup dice) int { return filteredSum(cup, weight) },
	}
}

func Score(cup []int, category string) int {
	return newCategory(category).scoreFor(cup)
}