package encode

import (
	"strings"
	"fmt"
	"unicode"
)

type pair struct {
	count int
	character byte
}

// Run Length Encoding algorithms

func RunLengthEncode(input string) string {
	if len(input) == 0 {
		return ""
	}
	return encodePairs(countingPairs(input))
}

func countingPairs(input string) []pair {
	pairs := []pair{}

	lastByte := input[0]
	count := 1

	for _, r := range input[1:] {
		if byte(r) == lastByte {
			count++
		} else {
			pairs = append(pairs, pair {
				count: count, 
				character: lastByte,
			})
			count = 1
			lastByte = byte(r)
		}
	}
	pairs = append(pairs, pair {
		count: count,
		character: lastByte,
	})

	return pairs
}

func encodePairs(pairs []pair) string {
	encoding := ""
	for _, p := range pairs {
		count, b := p.count, p.character
		if count == 1 {
			encoding += string(b)
		} else {
			encoding += fmt.Sprintf("%d%c", count, b)
		}
	}
	return encoding
}

// Run Length Decode algorithms

func RunLengthDecode(input string) string {
	if len(input) == 0 {
		return ""
	}
	pairs := decodedPairs(input)
	
	answer := ""
	for _, p := range pairs {
		answer += strings.Repeat(string(p.character), p.count)
	}
	return answer
}

func decodedPairs(input string) []pair {
	if len(input) == 0 {
		return []pair{}
	} 
	pairs := []pair{}

	scanningInteger := false
	runningTotal := 0
	for _, character := range input {
		if unicode.IsDigit(character) && !scanningInteger {
			scanningInteger = true
			runningTotal = int(character - '0')
		} else if unicode.IsDigit(character) {
			runningTotal = 10 * runningTotal + int(character - '0')
		} else if scanningInteger {
			pairs = append(pairs, pair { count: runningTotal, character: byte(character) })
			runningTotal = 0
			scanningInteger = false
		} else {
			pairs = append(pairs, pair { count: 1, character: byte(character) })
		}
	}
	return pairs
}
