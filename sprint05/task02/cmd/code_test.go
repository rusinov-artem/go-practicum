package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanCheckIfTreeIsBalanced(t *testing.T) {
	assert.True(t, true)

	t.Run("empty tree is balanced", func(t *testing.T) {
		root := (*Node)(nil)
		assert.True(t, isBalanced(root))
	})

	t.Run("single node tree is balanced", func(t *testing.T) {
		root := &Node{1, nil, nil}
		assert.True(t, isBalanced(root))
	})

	t.Run("single node with left child is balanced", func(t *testing.T) {
		left := &Node{1, nil, nil}
		root := &Node{1, left, nil}
		assert.True(t, isBalanced(root))
	})

	t.Run("single node with right child is balanced", func(t *testing.T) {
		right := &Node{1, nil, nil}
		root := &Node{1, nil, right}
		assert.True(t, isBalanced(root))
	})

	t.Run("node with ll child only is not child", func(t *testing.T) {
        leftLeft := &Node{1, nil, nil}     
		left := &Node{1, leftLeft, nil}
		root := &Node{1, left, nil}
		assert.False(t, isBalanced(root))
	})
}

func Test_CanCountDepth(t *testing.T) {
	t.Run("Can get depth of nil", func(t *testing.T) {
		d := []int{}
		iterateDepth(nil, func(depth int) {
			d = append(d, depth)
		})
		assert.Equal(t, []int{0}, d)
	})

	t.Run("depth of single node", func(t *testing.T) {
		root := &Node{0, nil, nil}
		d := []int{}
		iterateDepth(root, func(depth int) {
			d = append(d, depth)
		})
		assert.Equal(t, []int{1, 1}, d)
	})

	t.Run("of node with left children", func(t *testing.T) {
		left := &Node{0, nil, nil}
		root := &Node{0, left, nil}
		d := []int{}
		iterateDepth(root, func(depth int) {
			d = append(d, depth)
		})
		assert.Equal(t, []int{2, 2, 1}, d)
	})
}
