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

func Test_CanBuildHeap(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		data := []int{}
		buildHeap(data)
		assert.Equal(t, []int{}, data)
	})

	t.Run("single element", func(t *testing.T) {
		data := []int{1}
		buildHeap(data)
		assert.Equal(t, []int{1}, data)
	})
	t.Run("two elements", func(t *testing.T) {
		data := []int{1, 2}
		buildHeap(data)
		assert.Equal(t, []int{2, 1}, data)
	})
	t.Run("three elements", func(t *testing.T) {
		data := []int{1, 2, 3}
		buildHeap(data)
		assert.Equal(t, []int{3, 2, 1}, data)
	})
	t.Run("4 elements", func(t *testing.T) {
		data := []int{1, 2, 3, 4}
		buildHeap(data)
		assert.Equal(t, []int{4, 2, 3, 1}, data)
	})
}

func Test_PopSort(t *testing.T) {
   t.Run("empty", func(t *testing.T){
      heap := []int{} 
      popSort(heap)
      assert.Equal(t, []int{}, heap)
   })
   t.Run("1", func(t *testing.T){
      heap := []int{1} 
      popSort(heap)
      assert.Equal(t, []int{1}, heap)
   })
   t.Run("2", func(t *testing.T){
      heap := []int{2,1} 
      popSort(heap)
      assert.Equal(t, []int{1,2}, heap)
   })
   t.Run("3", func(t *testing.T){
      heap := []int{3,2,1} 
      popSort(heap)
      assert.Equal(t, []int{1,2,3}, heap)
   })
   t.Run("4", func(t *testing.T){
      heap := []int{4, 2, 3, 1} 
      popSort(heap)
      assert.Equal(t, []int{1,2,3,4}, heap)
   })
}
