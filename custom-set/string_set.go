package stringset

import (
	"fmt"
	"strings"
)

type Set map[string]bool

func New() Set {
	return Set{}
}

func NewFromSlice(slice []string) Set {
	set := New()
	for _, s := range slice {
		set[s] = true
	}
	return set
}

func (set Set) String() string {
	slice := []string{}

	for key := range set {
		slice = append(slice, fmt.Sprintf("%q", key))
	}

	return "{" + strings.Join(slice, ", ") + "}"
}

func (set Set) IsEmpty() bool {
	return len(set) == 0
}

func (set Set) Has(element string) bool {
	return set[element]
}

func Subset(s, t Set) bool {
	return any(s, func(key string) bool {
		return !t[key]
	})
}

func Disjoint(s, t Set) bool {
	return any(s, func(key string) bool {
		return t[key]
	})
}

func any(s Set, failingCondition func(string) bool) bool {
	for key := range s {
		if failingCondition(key) {
			return false
		}
	}
	return true
}

func Equal(s, t Set) bool {
	return Subset(s, t) && Subset(t, s)
}

func (set Set) Add(s string) {
	set[s] = true
}

func Intersection(s, t Set) Set {
	return reduce(Set{}, s, func(key string) bool {
		return s[key] && t[key]
	})
}

func Difference(s, t Set) Set {
	return reduce(Set{}, s, func(key string) bool {
		return !(s[key] && t[key])
	})
}

func Union(s, t Set) Set {
	return reduce(s, t, func(key string) bool {
		return true
	})
}

func reduce(original, s Set, addCondition func(key string) bool) Set {
	set := original
	for key := range s {
		if addCondition(key) {
			set.Add(key)
		}
	}
	return set
}