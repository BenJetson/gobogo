package bogo

import (
	"cmp"
	"log/slog"
	"math/rand/v2"
)

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
