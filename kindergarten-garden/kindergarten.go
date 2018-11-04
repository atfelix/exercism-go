package kindergarten

import (
	"sort"
	"unicode"
	"fmt"
	"strings"
)

var plantsByLetter = map[rune]string {
	'R': "radishes",
	'C': "clover",
	'G': "grass",
	'V': "violets",
}

type Garden struct {
	plants map[string][]string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	if !isValid(diagram, children) {
		return nil, fmt.Errorf("Invalid diagram:  There must be 2 rows and 2 columns for each child.")
	}
	garden := &Garden{ plants: map[string][]string{} }

	sortedChildren := make([]string, len(children))
	copy(sortedChildren, children)
	sort.Slice(sortedChildren, func(i, j int) bool {
		return sortedChildren[i] < sortedChildren[j]
	})

	for _, row := range strings.Split(diagram, "\n")[1:] {
		for index, r := range row {
			child := sortedChildren[index / 2]
			garden.plants[child] = append(garden.plants[child], plantsByLetter[r])
		}
	}

	return garden, nil
}

func (garden *Garden) Plants(child string) ([]string, bool) {
	plants, ok := garden.plants[child]
	return plants, ok
}

func isValid(diagram string, children []string) bool  {
	return areThereDuplicateChildren(children) &&
		   isDiagramValid(diagram, children)
}

func areThereDuplicateChildren(children []string) bool {
	setOfChildren := map[string]struct{}{}
	for _, child := range children {
		if _, ok := setOfChildren[child]; ok {
			return false
		}
		setOfChildren[child] = struct{}{}
	}
	return true
}

func isDiagramValid(diagram string, children []string) bool {
	rows := strings.Split(diagram, "\n")[1:]
	return len(rows) == 2 && 
		   len(rows[0]) == len(rows[1]) && 
		   len(rows[0]) == 2 * len(children) &&
		   isRowValid(rows[0]) &&
		   isRowValid(rows[1])
}

func isRowValid(row string) bool {
	for _, r := range row {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}