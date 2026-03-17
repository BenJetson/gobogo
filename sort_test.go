package bogo

import (
	"cmp"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testSort[T cmp.Ordered](t *testing.T, s []T) {
	c := make([]T, len(s))
	copy(c, s)
	slices.Sort(s)
	Sort(c)
	assert.Equal(t, s, c)
}

func TestSortPreSortedInts(t *testing.T) {
	preSortedInts := [][]int{
		{},
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 5, 6, 7},
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
	}
	for _, s := range preSortedInts {
		testSort(t, s)
	}
}

func TestSortPreSortedStrings(t *testing.T) {
	preSortedStrings := [][]string{
		{},
		{"a"},
		{"a", "b"},
		{"a", "b", "c"},
		{"a", "b", "c", "d"},
		{"a", "b", "c", "d", "e"},
		{"a", "b", "c", "d", "e", "f"},
		{"a", "b", "c", "d", "e", "f", "g"},
		{"a", "b", "c", "d", "e", "f", "g", "h"},
		{"a", "b", "c", "d", "e", "f", "g", "h", "i"},
		{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
		{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"},
	}
	for _, s := range preSortedStrings {
		testSort(t, s)

	}
}

func TestSortInts1Through2(t *testing.T) {
	testSort(t, []int{2, 1})
}

func TestSortInts1Through3(t *testing.T) {
	testSort(t, []int{3, 2, 1})
}

func TestSortInts1Through4(t *testing.T) {
	testSort(t, []int{4, 3, 2, 1})
}

func TestSortInts1Through5(t *testing.T) {
	testSort(t, []int{5, 4, 3, 2, 1})
}

func TestSortInts1Through6(t *testing.T) {
	testSort(t, []int{6, 5, 4, 3, 2, 1})
}

func TestSortInts1Through7(t *testing.T) {
	testSort(t, []int{7, 6, 5, 4, 3, 2, 1})
}

func TestSortInts1Through8(t *testing.T) {
	testSort(t, []int{8, 7, 6, 5, 4, 3, 2, 1})
}

func TestSortInts1Through9(t *testing.T) {
	testSort(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1})
}

func TestSortInts1Through10(t *testing.T) {
	testSort(t, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
}

func TestSortInts1Through11(t *testing.T) {
	testSort(t, []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
}

func TestSortLettersAThroughB(t *testing.T) {
	testSort(t, []string{"b", "a"})
}

func TestSortLettersAThroughC(t *testing.T) {
	testSort(t, []string{"c", "b", "a"})
}

func TestSortLettersAThroughD(t *testing.T) {
	testSort(t, []string{"d", "c", "b", "a"})
}

func TestSortLettersAThroughE(t *testing.T) {
	testSort(t, []string{"e", "d", "c", "b", "a"})
}

func TestSortLettersAThroughF(t *testing.T) {
	testSort(t, []string{"f", "e", "d", "c", "b", "a"})
}

func TestSortLettersAThroughG(t *testing.T) {
	testSort(t, []string{"g", "f", "e", "d", "c", "b", "a"})
}

func TestSortLettersAThroughH(t *testing.T) {
	testSort(t, []string{"h", "g", "f", "e", "d", "c", "b", "a"})
}

func TestSortLettersAThroughI(t *testing.T) {
	testSort(t, []string{"i", "h", "g", "f", "e", "d", "c", "b", "a"})
}

func TestSortLettersAThroughJ(t *testing.T) {
	testSort(t, []string{"j", "i", "h", "g", "f", "e", "d", "c", "b", "a"})
}

func TestSortLettersAThroughK(t *testing.T) {
	testSort(t, []string{"k", "j", "i", "h", "g", "f", "e", "d", "c", "b", "a"})
}
