package bogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSorted(t *testing.T) {
	assert.True(t, IsSorted([]int{}))
	assert.True(t, IsSorted([]int{1}))
	assert.True(t, IsSorted([]int{1, 2}))
	assert.True(t, IsSorted([]int{1, 2, 3}))
	assert.True(t, IsSorted([]int{1, 2, 3, 4}))
	assert.True(t, IsSorted([]int{1, 2, 3, 4, 5}))
	assert.False(t, IsSorted([]int{1, 2, 9, 3, 4, 5}))

	assert.True(t, IsSorted([]string{}))
	assert.True(t, IsSorted([]string{"a"}))
	assert.True(t, IsSorted([]string{"a", "b"}))
	assert.True(t, IsSorted([]string{"a", "b", "c"}))
	assert.True(t, IsSorted([]string{"a", "b", "c", "d"}))
	assert.True(t, IsSorted([]string{"a", "b", "c", "d", "e"}))
	assert.False(t, IsSorted([]string{"a", "b", "z", "c", "d", "e"}))
}
