package piglatin

import (
	"strings"
)

type wordScanner struct {
	word string
	cursor int
}

func newWordScanner(word string) *wordScanner {
	return &wordScanner{word: word, cursor: 0}
}

func (scanner *wordScanner) hasNext() bool {
	return scanner.cursor < len(scanner.word)
}

func (scanner *wordScanner) hasPrevious() bool {
	return scanner.cursor > 0
}

func (scanner *wordScanner) scan() string {
	if scanner.hasNext() {
		scanner.cursor++
		return string(scanner.word[scanner.cursor - 1])
	}
	return ""
}

func (scanner *wordScanner) lookback() {
	if scanner.hasPrevious() {
		scanner.cursor--
	}
}

func (scanner *wordScanner) currentToken() string {
	return string(scanner.word[scanner.cursor])
}

type pigLatinTokenizer struct {
	*wordScanner
	tokens []token
}

func newPigLatinTokenizer(word string) *pigLatinTokenizer {
	scanner := newWordScanner(word)
	return &pigLatinTokenizer{wordScanner: scanner, tokens: []token{}}
}

type token struct {
	value string
	isVowel bool
}

var tokens = []token {
	token{"str", false},
	token{"sch", false},
	token{"thr", false},
	token{"th", false},
	token{"tr", false},
	token{"qu", false},
	token{"ch", false},
	token{"sh", false},
	token{"rh", false},
	token{"st", false},
	token{"a", true},
	token{"b", false},
	token{"c", false},
	token{"d", false},
	token{"e", true},
	token{"f", false},
	token{"g", false},
	token{"h", false},
	token{"i", true},
	token{"j", false},
	token{"k", false},
	token{"l", false},
	token{"m", false},
	token{"n", false},
	token{"o", true},
	token{"p", false},
	token{"q", false},
	token{"r", false},
	token{"s", false},
	token{"t", false},
	token{"u", true},
	token{"v", false},
	token{"w", false},
	token{"x", false},
	token{"y", false},
	token{"z", false},
}

func (tokenizer *pigLatinTokenizer) populateTokens() {
	currentString := ""
	for tokenizer.hasNext() {
		currentString += tokenizer.scan()
		filteredResults := filterTokensBy(currentString)
		if len(filteredResults) == 1 && filteredResults[0].value == currentString {
			tokenizer.tokens = append(tokenizer.tokens, filteredResults[0])
			tokenizer.cursor += len(filteredResults[0].value) - len(currentString)
			currentString = ""
		} else if len(filteredResults) == 0 {
			currentString = string(currentString[:len(currentString) - 1])
			tokenizer.lookback()
			tokenizer.addExactTokenMatch(currentString)
			currentString = ""
		}
	}
	if len(currentString) != 0 {
		tokenizer.addExactTokenMatch(currentString)
	}
	tokenizer.cleanUp()
}

func (tokenizer *pigLatinTokenizer) cleanUp() {
	if tokenizer.tokens[0].value == "x" || tokenizer.tokens[0].value == "y" {
		if len(tokenizer.tokens) > 1 {
			tokenizer.tokens[0].isVowel = !tokenizer.tokens[1].isVowel
		}
	}
	if len(tokenizer.tokens) > 1 && !tokenizer.tokens[0].isVowel && tokenizer.tokens[1].value == "qu" {
		tokenizer.tokens = append(tokenizer.tokens[:1], tokenizer.tokens[2:]...)
		tokenizer.tokens[0] = token{value: tokenizer.tokens[0].value + "qu", isVowel: false}
	}
}

func filterTokensBy(prefix string) []token {
	results := []token{}
	for _, t := range tokens {
		if strings.HasPrefix(t.value, prefix) {
			results = append(results, t)
		}
	}
	return results
}

func (tokenizer *pigLatinTokenizer) addExactTokenMatch(tokenString string) {
	for _, t := range filterTokensBy(tokenString) {
		if t.value == tokenString {
			tokenizer.tokens = append(tokenizer.tokens, t)
		}
	}
}

func pigLatinWord(word string) string {
	tokenizer := newPigLatinTokenizer(word)
	tokenizer.populateTokens()
	pigLatinWord := ""
	ending := "ay"
	start := 0
	if !tokenizer.tokens[0].isVowel {
		ending = tokenizer.tokens[0].value + ending
		start = 1
	}
	for _, t := range tokenizer.tokens[start:] {
		pigLatinWord += t.value
	}
	pigLatinWord += ending
	return pigLatinWord
}

func Sentence(phrase string) string {
	words := strings.Split(phrase, " ")
	pigLatinWords := []string{}
	for _, word := range words {
		pigLatinWords = append(pigLatinWords, pigLatinWord(word))
	}
	return strings.Join(pigLatinWords, " ")
}