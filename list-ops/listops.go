package listops

type unaryFunc func(int) int
type binFunc func(int, int) int
type predFunc func(int) bool
type genericFunction func(int) interface{}
type result interface{}
type closure func(result, int) result

type IntList []int

func (list IntList) foldl(initial result, reduce closure) result {
	if list == nil || len(list) == 0 {
		return initial
	}
	return list[1:].foldl(reduce(initial, list[0]), reduce)
}

func (list IntList) Foldl(function binFunc, initial int) int {
	binaryClosure := func(r result, i int) result {
		return function(r.(int), i)
	}
	return list.foldl(initial, binaryClosure).(int)
}

func (list IntList) foldr(initial result, reduce closure) result {
	if list == nil || len(list) == 0 {
		return initial
	}
	return list[:len(list) - 1].foldr(reduce(initial, list[len(list) - 1]), reduce)
}

func (list IntList) Foldr(function binFunc, initial int) int {
	binaryClosure := func(r result, i int) result {
		return function(i, r.(int))
	}
	return list.Reverse().foldl(initial, binaryClosure).(int)
}

func (list IntList) Reverse() IntList {
	return list.foldr(IntList([]int{}), func(r result, i int) result {
		r = append(r.(IntList), i)
		return r
	}).(IntList)
}

func (list IntList) Filter(predicate predFunc) IntList {
	return list.foldl(IntList([]int{}), func(r result, i int) result {
		if predicate(i) {
			r = append(r.(IntList), i)
		}
		return r
	}).(IntList)
}

func (list IntList) Length() int {
	return list.Foldl(func (x, y int) int {
		return x + 1
	}, 0)
}

func (list IntList) Map(function unaryFunc) IntList {
	return list.foldl(IntList([]int{}), func(r result, i int) result {
		r = append(r.(IntList), function(i))
		return r
	}).(IntList)
}

func (list IntList) Append(otherList IntList) IntList {
	return otherList.foldl(list, func(r result, i int) result {
		r = append(r.(IntList), i)
		return r
	}).(IntList)
}

func (list IntList) Concat(listOfLists []IntList) IntList {
	for _, otherList := range listOfLists {
		list = list.Append(otherList)
	}
	return list
}
