package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_(t *testing.T) {
	assert.True(t, true)

	assert.Equal(t, 1, parent(2))
	assert.Equal(t, 1, parent(3))

	assert.Equal(t, 2, leftChild(1))
	assert.Equal(t, 4, leftChild(2))
	assert.Equal(t, 6, leftChild(3))

	assert.Equal(t, 3, rightChild(1))
	assert.Equal(t, 5, rightChild(2))
	assert.Equal(t, 7, rightChild(3))
}

func Test_CanShiftDown(t *testing.T) {
	t.Run("empty heap", func(t *testing.T) {
		heap := []int{0}
		r := up(heap, 0)
		assert.Equal(t, 0, r)
	})
	t.Run("single element heap", func(t *testing.T) {
		heap := []int{0, 1}
		r := up(heap, 1)
		assert.Equal(t, 1, r)
		assert.Equal(t, []int{0, 1}, heap)
	})
	t.Run("two elements heap", func(t *testing.T) {
		heap := []int{0, 1, 2}
		r := up(heap, 2)
		assert.Equal(t, 1, r)
		assert.Equal(t, []int{0, 2, 1}, heap)
	})
	t.Run("three elements heap", func(t *testing.T) {
		heap := []int{0, 5, 4, 30}
		r := up(heap, 3)
		assert.Equal(t, 1, r)
		assert.Equal(t, []int{0, 30, 4, 5}, heap)
	})
	t.Run("three elements heap 2", func(t *testing.T) {
		heap := []int{0, 6, 4, 5}
		r := up(heap, 3)
		assert.Equal(t, 3, r)
		assert.Equal(t, []int{0, 6, 4, 5}, heap)
	})
	t.Run("four elements heap", func(t *testing.T) {
		heap := []int{0, 6, 4, 5, 10}
		r := up(heap, 4)
		assert.Equal(t, 1, r)
		assert.Equal(t, []int{0, 10, 6, 5, 4}, heap)
	})
	t.Run("complex 1", func(t *testing.T) {
		heap := []int{0, 12, 6, 8, 3, 15, 7}
		r := up(heap, 5)
		assert.Equal(t, 1, r)
		assert.Equal(t, []int{0, 15, 12, 8, 3, 6, 7}, heap)
	})
}
