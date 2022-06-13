package utils

import (
	"sort"
)

// MaxHeapString returns the max value in a map.
// The map is sorted by value
func MaxHeapString(m map[string]int64) string {
	keys := make([]string, 0) // Create a slice of keys for sorting the map

	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool { // Sort by occurence
		return m[keys[i]] < m[keys[j]]
	})

	return keys[len(keys)-1] // Max
}
