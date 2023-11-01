package main

import (
	"fmt"
	"testing"

	"github.com/rusinov-artem/go-practicum/daily"
	"github.com/stretchr/testify/assert"
)

func Test_CanSortIntegers(t *testing.T) {
	data := []int{5, 4, 3, 10, 1, 2, 10}

	c := &SortableIndex{
		Index: iRange(len(data)),
	}

	c.Less = func(i, j int) bool {
		return data[c.Index[i]] > data[c.Index[j]]
	}

	qSort(c)

	fmt.Println(c)
	for _, v := range c.Index {
		fmt.Println(data[v])
	}
}

func Test_Part(t *testing.T) {
	t.Run("swap", func(t *testing.T) {
		inp := []int{1, 2}
		swap(inp, 0, 1)
		assert.Equal(t, []int{2, 1}, inp)
	})
	t.Run("Can part nothing", func(t *testing.T) {
		inp := []int{}
		assert.Equal(t, 0, part(inp, 0, 0))
	})
	t.Run("Can part single element", func(t *testing.T) {
		inp := []int{1}
		assert.Equal(t, 0, part(inp, 0, 1))
	})
	t.Run("Can part 1,2", func(t *testing.T) {
		inp := []int{1, 2}
		assert.Equal(t, 1, part(inp, 0, 2))
		assert.Equal(t, []int{1, 2}, inp)
	})
	t.Run("Can part 2,1", func(t *testing.T) {
		inp := []int{2, 1}
		assert.Equal(t, 0, part(inp, 0, 2))
		assert.Equal(t, []int{1, 2}, inp)
	})
	t.Run("Can move left 1 time", func(t *testing.T) {
		inp := []int{1, 9, 10, 5}
		assert.Equal(t, 1, part(inp, 0, 4))
		assert.Equal(t, []int{1, 5, 10, 9}, inp)
	})
	t.Run("Can swap left and right 1 time", func(t *testing.T) {
		inp := []int{9, 2, 5}
		assert.Equal(t, 1, part(inp, 0, 3))
		assert.Equal(t, []int{2, 5, 9}, inp)
	})
	t.Run("Can part with same numers", func(t *testing.T) {
		inp := []int{9, 2, 2}
		assert.Equal(t, 1, part(inp, 0, 3))
		assert.Equal(t, []int{2, 2, 9}, inp)
	})
	t.Run("Bug1", func(t *testing.T) {
		inp := []int{2, 3, 1}
		assert.Equal(t, 0, part(inp, 0, 3))
		assert.Equal(t, []int{1, 3, 2}, inp)
	})
	t.Run("Can qsort numers", func(t *testing.T) {
		inp := []int{5, 4, 3, 10, 1, 2, 10}
		qs(inp, 0, 7)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 10, 10}, inp)
	})
}

func Test_Brutoforse(t *testing.T) {
	gen := daily.NewPermGenerator([]int{1, 2, 3,4,5})
	for gen.Next() {
		inp := gen.Get()
		fmt.Println(inp)
		qs(inp, 0, len(inp))
		assert.Equal(t, []int{1, 2, 3,4,5}, inp)
	}
}

