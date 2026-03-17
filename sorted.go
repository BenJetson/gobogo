package bogo

import "cmp"

// IsSorted determines whether or not the input slice is sorted or not.
func IsSorted[T cmp.Ordered](s []T) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}
