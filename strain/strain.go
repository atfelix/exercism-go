package strain

type Ints []int
type Lists [][]int
type Strings []string

func (list Ints) Keep(filter func(int) bool) Ints {
	return list.filter(func(element int) bool {
		return filter(element)
	})
}

func (list Ints) Discard(filter func(int) bool) Ints {
	return list.filter(func(element int) bool {
		return !filter(element)
	})
}

func (list Ints) filter(condition func(int) bool) Ints {
	if list == nil {
		return Ints(nil)
	}
	filteredInts := Ints{}
	for _, element := range list {
		if condition(element) {
			filteredInts = append(filteredInts, element)
		}
	}
	return filteredInts
}

func (list Strings) Keep(filter func(string) bool) Strings {
	if list == nil {
		return Strings(nil)
	}
	filteredStrings := Strings{}
	for _, element := range list {
		if filter(element) {
			filteredStrings = append(filteredStrings, element)
		}
	}
	return filteredStrings
}

func (list Lists) Keep(filter func([]int) bool) Lists {
	if list == nil {
		return Lists(nil)
	}
	filteredLists := Lists{}
	for _, element := range list {
		if filter(element) {
			filteredLists = append(filteredLists, element)
		}
	}
	return filteredLists
}