package etl

import (
	"strings"
)

func Transform(dictionary map[int][]string) map[string]int {
	otherDictionary := map[string]int{}
	for key, list := range dictionary {
		for _, value := range list {
			otherDictionary[strings.ToLower(value)] = key
		}
	}
	return otherDictionary
}