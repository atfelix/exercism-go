package twelve

import (
	"fmt"
)

const (
	start = "On the %s day of Christmas my true love gave to me"
)

var (
	days = []string{
		"first",
		"second",
		"third",
		"fourth",
		"fifth",
		"sixth",
		"seventh",
		"eighth",
		"ninth",
		"tenth",
		"eleventh",
		"twelfth",
	}
	gifts = []string {
		"a Partridge in a Pear Tree",
		"two Turtle Doves",
		"three French Hens",
		"four Calling Birds",
		"five Gold Rings",
		"six Geese-a-Laying",
		"seven Swans-a-Swimming",
		"eight Maids-a-Milking",
		"nine Ladies Dancing",
		"ten Lords-a-Leaping",
		"eleven Pipers Piping",
		"twelve Drummers Drumming",
	}
)

func Verse(input int) string {
	s := ", "
	if input > 1 {
		s += "and "
	}
	s = s + gifts[0]
	for index := 1; index < input; index ++ {
		s = ", " + gifts[index] + s
	}
	return fmt.Sprintf(start, days[input - 1]) + s + "."
}

func Song() string {
	s := ""
	for index := range days {
		s += Verse(index + 1) + "\n"
	}
	return s
}
