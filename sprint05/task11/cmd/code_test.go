package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanPrintRange(t *testing.T) {
	assert.True(t, true)
	t.Run("no error if empty", func(t *testing.T) {
		node := (*Node)(nil)
		values := make([]int, 0)
		iterateRange(node, 0, 100, func(val int) {
			values = append(values, val)
		})
		assert.Empty(t, values)
	})

	t.Run("single node in range", func(t *testing.T) {
		node := &Node{10, nil, nil}
		values := make([]int, 0)
		iterateRange(node, 0, 100, func(val int) {
			values = append(values, val)
		})
		assert.Equal(t, 1, len(values))
	})

	t.Run("single node value eq left", func(t *testing.T) {
		node := &Node{0, nil, nil}
		values := make([]int, 0)
		iterateRange(node, 0, 100, func(val int) {
			values = append(values, val)
		})
		assert.Equal(t, 1, len(values))
	})

	t.Run("single node value eq right", func(t *testing.T) {
		node := &Node{100, nil, nil}
		values := make([]int, 0)
		iterateRange(node, 0, 100, func(val int) {
			values = append(values, val)
		})
		assert.Equal(t, 1, len(values))
	})

	t.Run("single node out of range", func(t *testing.T) {
		node := &Node{1000, nil, nil}
		values := make([]int, 0)
		iterateRange(node, 0, 100, func(val int) {
			values = append(values, val)
		})
		assert.Equal(t, 0, len(values))
	})

	t.Run("has node with 2 children", func(t *testing.T) {
		root := &Node{50, nil, nil}
		left := &Node{10, nil, nil}
		right := &Node{100, nil, nil}

		root.left = left
		root.right = right

		values := make([]int, 0)
		iterateRange(root, 0, 100, func(val int) {
			values = append(values, val)
		})

		assert.Equal(t, []int{10, 50, 100}, values)
	})

	t.Run("has node with 2 children. And one, out of range", func(t *testing.T) {
		root := &Node{50, nil, nil}
		left := &Node{10, nil, nil}
		right := &Node{1000, nil, nil}

		root.left = left
		root.right = right

		values := make([]int, 0)
		iterateRange(root, 0, 100, func(val int) {
			values = append(values, val)
		})

		assert.Equal(t, []int{10, 50}, values)
	})
}

func Test_Integration(t *testing.T) {
	root := &Node{50, nil, nil}
	left := &Node{10, nil, nil}
	right := &Node{1000, nil, nil}

	root.left = left
	root.right = right

	printRange(root, 0, 100)
}
