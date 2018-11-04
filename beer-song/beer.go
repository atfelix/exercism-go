package beer

import (
	"bytes"
	"fmt"
)

func Verse(verse int) (string, error) {
	if 0 <= verse && verse < 100 {
		return firstLineOf(verse) + secondLineOf(verse), nil
	}
	return "", fmt.Errorf("Invalid verse number")
}

func numberOfBottlesIn(verse int, capitalize bool) interface{} {
	if verse == 0 && capitalize{
		return "No more"
	} else if verse == 0 {
		return "no more"
	}
	return verse
}

func pluralOrSingularNounFor(verse int) string {
	if verse == 1 {
		return "bottle"
	}
	return "bottles"
}

func pronounFor(verse int) string {
	if verse == 1 {
		return "it"
	}
	return "one"
}

func firstLineOf(verse int) string {
	noun := pluralOrSingularNounFor(verse)
	return fmt.Sprintf("%v %s of beer on the wall, %v %s of beer.\n", numberOfBottlesIn(verse, true), noun, numberOfBottlesIn(verse, false), noun)
}

func secondLineOf(verse int) string {
	if verse == 0 {
		return fmt.Sprintf("Go to the store and buy some more, 99 bottles of beer on the wall.\n")
	}
	return fmt.Sprintf("Take %s down and pass it around, %v %s of beer on the wall.\n", pronounFor(verse), numberOfBottlesIn(verse - 1, false), pluralOrSingularNounFor(verse - 1))
}

func Verses(upperbound, lowerbound int) (string, error) {
	if upperbound < lowerbound {
		return "", fmt.Errorf("upperbound is less than lowerbound")
	}
	var buffer bytes.Buffer
	for verse := upperbound; verse >= lowerbound; verse-- {
		stringVerse, err := Verse(verse)
		if err != nil {
			return "", err
		}
		buffer.WriteString(stringVerse + "\n")
	}
	return buffer.String(), nil
}

func Song() string {
	s, _ := Verses(99, 0)
	return s
}