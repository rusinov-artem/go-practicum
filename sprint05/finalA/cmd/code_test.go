package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParentChild(t *testing.T) {
	assert.True(t, true)

	assert.Equal(t, 0, parent(1))
	assert.Equal(t, 0, parent(2))
	assert.Equal(t, 1, parent(3))
	assert.Equal(t, 1, parent(4))
	assert.Equal(t, 2, parent(5))
	assert.Equal(t, 2, parent(6))

	assert.Equal(t, 1, leftChild(0))
	assert.Equal(t, 3, leftChild(1))
	assert.Equal(t, 5, leftChild(2))

	assert.Equal(t, 2, rightChild(0))
	assert.Equal(t, 4, rightChild(1))
	assert.Equal(t, 6, rightChild(2))
}

func newSortableForInts(data []int) *SortableIndex {
	sortable := &SortableIndex{Index: iRange(len(data))}
	sortable.Less = func(i, j int) bool {
		i = sortable.Index[i]
		j = sortable.Index[j]
		return data[i] < data[j]
	}
	return sortable
}

func extract(sortable *SortableIndex, data []int) []int {
	result := make([]int, len(data))
	for i := 0; i < len(data); i++ {
		result[i] = data[sortable.Index[i]]
	}
	return result
}

func Test_CanBuildHeap(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		data := []int{}
		sortable := newSortableForInts(data)
		heapify(sortable)
		assert.Equal(t, []int{}, extract(sortable, data))
	})

	t.Run("single element", func(t *testing.T) {
		data := []int{1}
		sortable := newSortableForInts(data)
		heapify(sortable)
		assert.Equal(t, []int{1}, extract(sortable, data))
	})
	t.Run("two elements", func(t *testing.T) {
		data := []int{1, 2}
		sortable := newSortableForInts(data)
		heapify(sortable)
		assert.Equal(t, []int{2, 1}, extract(sortable, data))
	})
	t.Run("three elements", func(t *testing.T) {
		data := []int{1, 2, 3}
		sortable := newSortableForInts(data)
		heapify(sortable)
		assert.Equal(t, []int{3, 2, 1}, extract(sortable, data))
	})
	t.Run("4 elements", func(t *testing.T) {
		data := []int{1, 2, 3, 4}
		sortable := newSortableForInts(data)
		heapify(sortable)
		assert.Equal(t, []int{4, 2, 3, 1}, extract(sortable, data))
	})
}

func Test_PopSort(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		heap := []int{}
		sortable := newSortableForInts(heap)
		popSort(sortable)
		assert.Equal(t, []int{}, extract(sortable, heap))
	})
	t.Run("1", func(t *testing.T) {
		heap := []int{1}
		sortable := newSortableForInts(heap)
		popSort(sortable)
		assert.Equal(t, []int{1}, extract(sortable, heap))
	})
	t.Run("2", func(t *testing.T) {
		heap := []int{2, 1}
		sortable := newSortableForInts(heap)
		popSort(sortable)
		assert.Equal(t, []int{1, 2}, extract(sortable, heap))
	})
	t.Run("3", func(t *testing.T) {
		heap := []int{3, 2, 1}
		sortable := newSortableForInts(heap)
		popSort(sortable)
		assert.Equal(t, []int{1, 2, 3}, extract(sortable, heap))
	})
	t.Run("4", func(t *testing.T) {
		heap := []int{4, 2, 3, 1}
		sortable := newSortableForInts(heap)
		popSort(sortable)
		assert.Equal(t, []int{1, 2, 3, 4}, extract(sortable, heap))
	})
}
