package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanCountDepth(t *testing.T) {
	assert.True(t, true)
	t.Run("empty tree", func(t *testing.T) {
		root := (*Node)(nil)
		depth := countDepth(root, 0)
		assert.Equal(t, 0, depth)
	})

	t.Run("single node", func(t *testing.T) {
		root := &Node{1, nil, nil}
		depth := countDepth(root, 0)
		assert.Equal(t, 1, depth)
	})

	t.Run("root with left child", func(t *testing.T) {
		root := &Node{1, nil, nil}
		left := &Node{4, nil, nil}
		root.left = left
		depth := countDepth(root, 0)
		assert.Equal(t, 2, depth)
	})

	t.Run("root with right child", func(t *testing.T) {
		root := &Node{1, nil, nil}
		right := &Node{4, nil, nil}
		root.right = right
		depth := countDepth(root, 0)
		assert.Equal(t, 2, depth)
	})

	t.Run("root with 2 child, and one grandson on the left lefft", func(t *testing.T) {
		root := &Node{1, nil, nil}
		right := &Node{4, nil, nil}
		left := &Node{40, nil, nil}
		grandson := &Node{400, nil, nil}

		root.left = left
		root.right = right
		root.left.left = grandson

		root.right = right
		depth := countDepth(root, 0)
		assert.Equal(t, 3, depth)
	})

	t.Run("root with 2 child, and one grandson on the right right", func(t *testing.T) {
		root := &Node{1, nil, nil}
		right := &Node{4, nil, nil}
		left := &Node{40, nil, nil}
		grandson := &Node{400, nil, nil}

		root.left = left
		root.right = right
		root.right.right = grandson

		root.right = right
		depth := countDepth(root, 0)
		assert.Equal(t, 3, depth)
	})
}
