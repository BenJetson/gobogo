package bogo

import (
	"cmp"
	"log/slog"
	"math/rand/v2"
)

// Shuffle performs an in-place randomization of the given slice, swapping
// somewhere between 0.5 and 1.5 times the number of elements in the slice.
func Shuffle[T cmp.Ordered](s []T) {
	slog.Debug("shuffling")
	count := len(s)
	swapCount := (count / 2) + rand.IntN(count)
	for n := 0; n < swapCount; n++ {
		i := rand.IntN(count)
		j := rand.IntN(count)

		temp := s[j]
		s[j] = s[i]
		s[i] = temp
	}
}
