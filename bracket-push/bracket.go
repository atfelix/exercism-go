package brackets

var openBrackets = map[rune]bool {
	rune('{'): true,
	rune('['): true,
	rune('('): true,
}

var closeBrackets = map[rune]rune {
	rune('}'): rune('{'),
	rune(']'): rune('['),
	rune(')'): rune('('),
}

func Bracket(input string) (bool, error) {
	stack := []rune{}

	for _, b := range input {
		if openBrackets[b] {
			stack = append(stack, b)
		} else if rr, ok := closeBrackets[b]; ok && len(stack) == 0 {
			return false, nil
		} else if ok && stack[len(stack) - 1] != rr {
			return false, nil
		} else if ok {
			stack = stack[:len(stack) - 1]
		}
	}
	return len(stack) == 0, nil
}