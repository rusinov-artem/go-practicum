package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanFindInBrokenArray(t *testing.T) {
	t.Run("Empty input", func(t *testing.T) {
		inp := []int{}
		idx := brokenSearch(inp, 5)
		assert.Equal(t, -1, idx)
	})

	t.Run("Single element", func(t *testing.T) {
		inp := []int{5}
		idx := brokenSearch(inp, 5)
		assert.Equal(t, 0, idx)
	})

	t.Run("2 elements. Can find second", func(t *testing.T) {
		inp := []int{1, 5}
		idx := brokenSearch(inp, 5)
		assert.Equal(t, 1, idx)
	})

	t.Run("2 elements. Can find first", func(t *testing.T) {
		inp := []int{1, 5}
		idx := brokenSearch(inp, 1)
		assert.Equal(t, 0, idx)
	})

	t.Run("BrokenArr mid < k", func(t *testing.T) {
		inp := []int{4, 5, 6, 0, 1, 2, 3}
		idx := brokenSearch(inp, 2)
		assert.Equal(t, 5, idx)
	})

	t.Run("BrokenArr a[mid] < k and a[mid] > a[end]", func(t *testing.T) {
		inp := []int{4, 5, 6, 0, 1, 2, 3}
		idx := brokenSearch(inp, 6)
		assert.Equal(t, 2, idx)
	})

	t.Run("BrokenArr a[mid] < k && a[mid] > a[end] #2", func(t *testing.T) {
		inp := []int{4, 5, 0, 1, 2, 3}
		idx := brokenSearch(inp, 4)
		assert.Equal(t, 0, idx)
	})

	t.Run("BrokenArr a[mid] == k", func(t *testing.T) {
		inp := []int{4, 5, 6, 7, 1, 2, 3}
		idx := brokenSearch(inp, 7)
		assert.Equal(t, 3, idx)
	})

	t.Run("a[mid] > k && a[begin] < k", func(t *testing.T) {
		inp := []int{4, 5, 10, 11, 22, 33}
		idx := brokenSearch(inp, 4)
		assert.Equal(t, 0, idx)
	})

	t.Run("a[mid] > k && a[begin] > k", func(t *testing.T) {
		inp := []int{4, 5, 10, 11, 0, 1}
		idx := brokenSearch(inp, 1)
		assert.Equal(t, 5, idx)
	})

	t.Run("bug 1", func(t *testing.T) {
		inp := []int{9, 0, 1, 2, 3, 4}
		idx := brokenSearch(inp, 9)
		assert.Equal(t, 0, idx)
	})

	t.Run("bug 2", func(t *testing.T) {
		inp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		idx := brokenSearch(inp, 7)
		assert.Equal(t, 6, idx)
	})
}

func Test_UltimateTest(t *testing.T) {
	inp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for j := 0; j < 10; j++ {
		for i := 0; i < 10; i++ {
			t.Run(fmt.Sprintf("%v | %d", inp, i), func(t *testing.T) {
				assert.Equal(t, lineSearch(inp, i), brokenSearch(inp, i))
			})
		}
		inp = append(inp[1:], inp[0])
	}
}

func lineSearch(inp []int, k int) int {
	for i := 0; i < len(inp); i++ {
		if inp[i] == k {
			return i
		}
	}
	return -1
}
