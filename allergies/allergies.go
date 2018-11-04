package allergies

import (
	"sort"
)

var allergyScores = map[string]uint {
	"eggs": 1 << 0,
	"peanuts": 1 << 1,
	"shellfish": 1 << 2,
	"strawberries": 1 << 3,
	"tomatoes": 1 << 4,
	"chocolate": 1 << 5,
	"pollen": 1 << 6,
	"cats": 1 << 7,
}

func AllergicTo(score uint, substance string) bool {
	return allergyScores[substance] & score != 0
}

func Allergies(score uint) []string {
	allergies := []string{}
	for substance, _ := range allergyScores {
		if AllergicTo(score, substance) {
			allergies = append(allergies, substance)
		}
	}
	sort.Slice(allergies, func(i, j int) bool {
		return allergyScores[allergies[i]] < allergyScores[allergies[j]]
	})
	return allergies
}