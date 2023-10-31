package daily

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanFindElement(t *testing.T) {
	inp := []int{1, 3, 5, 8, 19, 23, 99, 100}
	idx := search(inp, 5)
	assert.Equal(t, 2, idx)
}

func Test_TestBinarySearch(t *testing.T) {
	t.Run("search in empty slice return -1", func(t *testing.T) {
		inp := []int{}
		assert.Equal(t, -1, search(inp, 4))
	})

	t.Run("if no element found return -1", func(t *testing.T) {
		inp := []int{1}
		assert.Equal(t, -1, search(inp, 4))
	})

	t.Run("can find element in slice with single element", func(t *testing.T) {
		inp := []int{1}
		assert.Equal(t, 0, search(inp, 1))
	})

	t.Run("can find element in slice with two elements", func(t *testing.T) {
		inp := []int{1, 2}
		assert.Equal(t, 1, search(inp, 2))
	})

	t.Run("can look at right subarray", func(t *testing.T) {
		inp := []int{1, 2, 3}
		assert.Equal(t, 2, search(inp, 3))
	})

	t.Run("can look at left subarray", func(t *testing.T) {
		inp := []int{1, 2, 3}
		assert.Equal(t, 0, search(inp, 1))
	})

	t.Run("can find left element", func(t *testing.T) {
		inp := []int{1, 2, 3, 5, 19, 20, 33, 44, 90, 100}
		assert.Equal(t, 0, search(inp, 1))
	})

	t.Run("can find right element", func(t *testing.T) {
		inp := []int{1, 2, 3, 5, 19, 20, 33, 44, 90, 100}
		assert.Equal(t, 9, search(inp, 100))
	})

	t.Run("can find middle element", func(t *testing.T) {
		inp := []int{1, 2, 3, 5, 19, 20, 33, 44, 90, 100}
		assert.Equal(t, 5, search(inp, 20))
	})
}

func Test_CanFindLowerBound(t *testing.T) {
	t.Run("can find lower bound", func(t *testing.T) {
		inp := []int{1, 2, 3, 5, 19, 20, 33, 44, 90, 100}
		for i := 0; i < len(inp); i++ {
			assert.Equal(t, i, lowerBound(inp, inp[i]))
		}
	})
	t.Run("can find lower bound", func(t *testing.T) {
		inp := []int{1, 2, 3, 5, 5, 5, 5, 5, 5, 19, 20, 33, 44, 90, 100}
		assert.Equal(t, 3, lowerBound(inp, 5))
	})

	t.Run("can find lower bound", func(t *testing.T) {
		inp := []int{1, 2, 3, 19, 20, 33, 44, 90, 100}
		assert.Equal(t, 3, lowerBound(inp, 19))
	})
}

func Test_CanFindUpperBound(t *testing.T) {
	t.Run("can find upper bound", func(t *testing.T) {
		inp := []int{1, 2, 3, 5, 19, 20, 33, 44, 90, 100}
		for i := 0; i < len(inp); i++ {
			assert.Equal(t, i+1, upperBound(inp, inp[i]))
		}
	})
	t.Run("can find upper bound 5555", func(t *testing.T) {
		inp := []int{1, 2, 3, 5, 5, 5, 5, 5, 5, 19, 20, 33, 44, 90, 100}
		assert.Equal(t, 9, upperBound(inp, 5))
	})
}
