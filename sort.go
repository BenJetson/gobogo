package bogo

import (
	"cmp"
)

func Sort[T cmp.Ordered](s []T) {
	for !IsSorted(s) {
		Shuffle(s)
	}
}
