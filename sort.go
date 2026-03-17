package bogo

import (
	"cmp"
)

// Sort accepts a slice of any ordered type and performs an in-place sort of the
// data using the bogosort algorithm.
//
// If you call this function, good luck.
func Sort[T cmp.Ordered](s []T) {
	for !IsSorted(s) {
		Shuffle(s)
	}
}
