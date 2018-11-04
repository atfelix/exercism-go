package sublist

import (
	"reflect"
)

type Relation string

func Sublist(listOne, listTwo []int) Relation {
	if len(listOne) == 0 && len(listTwo) == 0 {
		return Relation("equal")
	} else if len(listOne) == 0 {
		return Relation("sublist")
	} else if len(listTwo) == 0 {
		return Relation("superlist")
	} else if reflect.DeepEqual(listOne, listTwo) {
		return Relation("equal")
	} else if len(listOne) == len(listTwo) {
		return Relation("unequal")
	}
	
	strictSublistString := Relation("sublist")

	if len(listOne) > len(listTwo) {
		listOne, listTwo = listTwo, listOne
		strictSublistString = Relation("superlist")
	}

	for i := 0; i + len(listOne) < len(listTwo) + 1; i++ {
		if slicesEqual(listOne, listTwo[i:i+len(listOne)]) {
			return strictSublistString
		}
	}
	return Relation("unequal")
}

func slicesEqual(sliceOne, sliceTwo []int) bool {
	return reflect.DeepEqual(sliceOne, sliceTwo)
}
