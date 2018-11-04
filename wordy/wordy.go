package wordy

import (
	"fmt"
	"strconv"
	"strings"
)

var newStrings = map[string]string {
	"What is ": "",
	"?": "",
	"plus": "+",
	"minus": "-",
	"multiplied by": "*",
	"divided by": "/",
}

var precedence = map[string]int {
	"*": 1,
	"/": 1,
	"+": 1,
	"-": 1,
}

func tokens(question string) []string {
	for key, value := range newStrings {
		question = strings.Replace(question, key, value, -1)
	}
	return strings.Fields(question)
}

func reversePolishNotation(tokens []string) (rpn string, err error) {
	defer func() {
		if r := recover(); r != nil {
			rpn, err = "", fmt.Errorf("Invalid tokens")
		}
	}()
	for i := 1; i < len(tokens); i += 2 {
		tokens[i], tokens[i + 1] = tokens[i + 1], tokens[i]
	}
	return strings.Join(tokens, " "), nil
}

func evaluate(rpn string) (int, error) {
	stack := []int{}
	for _, token := range strings.Split(rpn, " ") {
		if number, err := strconv.Atoi(token); err == nil {
			stack = append(stack, number)
		} else {
			if len(stack) < 2 {
				return 0, fmt.Errorf("Invalid question")
			}
			operand1, operand2 := stack[len(stack) - 2], stack[len(stack) - 1]
			stack = stack[:len(stack) - 2]
			stack = append(stack, evaluateMath(token, operand1, operand2))
		}
	}
	if len(stack) == 0 {
		return 0, fmt.Errorf("Invalid question")
	}
	return stack[0], nil
}

func evaluateMath(operator string, operand1, operand2 int) int {
	switch operator {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "*":
		return operand1 * operand2
	case "/":
		return operand1 / operand2
	default:
		return 0
	}
}

func Answer(question string) (int, bool) {
	rpn, err := reversePolishNotation(tokens(question))
	if err != nil {
		return 0, false
	}
	answer, err := evaluate(rpn)
	if err != nil {
		return 0, false
	}
	return answer, true
}