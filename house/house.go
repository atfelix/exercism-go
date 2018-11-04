package house

import (
	"strings"
	"fmt"
)

var verbNouns = []struct {
	verb string
	noun string
}{
	{"", ""},
	{"lay in", "the house that Jack built"},
	{"ate", "the malt"},
	{"killed", "the rat"},
	{"worried", "the cat"},
	{"tossed", "the dog"},
	{"milked", "the cow with the crumpled horn"},
	{"kissed", "the maiden all forlorn"},
	{"married", "the man all tattered and torn"},
	{"woke", "the priest all shaven and shorn"},
	{"kept", "the rooster that crowed in the morn"},
	{"belonged to", "the farmer sowing his corn"},
	{"", "the horse and the hound and the horn"},
}

func Verse(index int) string {
	return fmt.Sprintf("This is %v%v.", verbNouns[index].noun, recursiveVerse(index - 1))
}

func recursiveVerse(index int) string {
	if index == 0 {
		return ""
	}
	return fmt.Sprintf("\nthat %v %v%v", verbNouns[index].verb, verbNouns[index].noun, recursiveVerse(index - 1))
}

func Song() string {
	verses := []string{}
	for index := 1; index < len(verbNouns); index++ {
		verses = append(verses, Verse(index))
	}
	return strings.Join(verses, "\n\n")
}