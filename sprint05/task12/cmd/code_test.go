package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanFindParentAndChildInHeap(t *testing.T) {
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
	t.Run("If there is empty heap", func(t *testing.T) {
		r := down([]int{0}, 0)
		assert.Equal(t, 0, r)
	})

	t.Run("If there is single element", func(t *testing.T) {
		r := down([]int{0, 1}, 1)
		assert.Equal(t, 1, r)
	})

	t.Run("If there is 2 elements", func(t *testing.T) {
		heap := []int{0, 1, 2}
		r := down(heap, 1)
		assert.Equal(t, 2, r)
		assert.Equal(t, []int{0, 2, 1}, heap)
	})

	t.Run("If there is 3 elements", func(t *testing.T) {
		heap := []int{0, 1, 2, 0}
		r := down(heap, 1)
		assert.Equal(t, 2, r)
		assert.Equal(t, []int{0, 2, 1, 0}, heap)
	})

	t.Run("If there is 4 elements", func(t *testing.T) {
		heap := []int{0, 1, 3, 2, 0}
		r := down(heap, 1)
		assert.Equal(t, 2, r)
		assert.Equal(t, []int{0, 3, 1, 2, 0}, heap)
	})

	t.Run("If there is 5 elements", func(t *testing.T) {
		heap := []int{0, 1, 4, 3, 2, 0}
		r := down(heap, 1)
		assert.Equal(t, 4, r)
		assert.Equal(t, []int{0, 4, 2, 3, 1, 0}, heap)
	})
	t.Run("complex 1", func(t *testing.T) {
		heap := []int{0, 12, 1, 8, 3, 4, 7}
		r := down(heap, 2)
		assert.Equal(t, 5, r)
		assert.Equal(t, []int{0, 12, 4, 8, 3, 1, 7}, heap)
	})
	t.Run("complex 2", func(t *testing.T) {
		heap := []int{0, 12, 1, 8, 3, 4, 7}
		r := siftDown(heap, 2)
		assert.Equal(t, 5, r)
		assert.Equal(t, []int{0, 12, 4, 8, 3, 1, 7}, heap)
	})
}
