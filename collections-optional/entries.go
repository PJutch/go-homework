package collections

import (
	"cmp"
	"slices"
)

type KeyValue struct {
	key   int
	value string
}

func Entries(map_ map[int]string) []KeyValue {
	entries := make([]KeyValue, 0, len(map_))
	for key, value := range map_ {
		entries = append(entries, KeyValue{key, value})
	}

	slices.SortFunc(entries, func(lhs, rhs KeyValue) int {
		return cmp.Compare(lhs.key, rhs.key)
	})

	return entries
}
