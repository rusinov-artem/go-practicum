package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanRemoveNode(t *testing.T) {
	assert.True(t, true)

	t.Run("empty tree", func(t *testing.T) {
		root := (*Node)(nil)
		newRoot := remove(root, 10)
		assert.Equal(t, (*Node)(nil), newRoot)
	})

	t.Run("single element", func(t *testing.T) {
		root := &Node{10, nil, nil}
		newRoot := remove(root, 10)
		assert.Equal(t, (*Node)(nil), newRoot)
	})

	t.Run("node has left child", func(t *testing.T) {
		left := &Node{5, nil, nil}
		root := &Node{10, left, nil}
		newRoot := remove(root, 10)
		assert.Equal(t, left, newRoot)
	})

	t.Run("node has left and right children", func(t *testing.T) {
		left := &Node{5, nil, nil}
		right := &Node{15, nil, nil}
		root := &Node{10, left, right}
		newRoot := remove(root, 10)
		assert.Equal(t, right, newRoot)
		assert.Equal(t, 15, newRoot.value)
		assert.Equal(t, 5, newRoot.left.value)
	})

	t.Run("node has right child", func(t *testing.T) {
		right := &Node{15, nil, nil}
		root := &Node{10, nil, right}
		newRoot := remove(root, 10)
		assert.Equal(t, right, newRoot)
		assert.Equal(t, 15, newRoot.value)
		assert.Equal(t, (*Node)(nil), newRoot.left)
		assert.Equal(t, (*Node)(nil), newRoot.right)
		assert.Equal(t, (*Node)(nil), root.left)
		assert.Equal(t, (*Node)(nil), root.right)
	})

	t.Run("can remove right child", func(t *testing.T) {
		left := &Node{5, nil, nil}
		right := &Node{15, nil, nil}
		root := &Node{10, left, right}
		newRoot := remove(root, 15)
		assert.Equal(t, root, newRoot)
		assert.Equal(t, 10, newRoot.value)
		assert.Equal(t, 5, newRoot.left.value)
		assert.Equal(t, (*Node)(nil), newRoot.right)
	})

	t.Run("can remove left child", func(t *testing.T) {
		left := &Node{5, nil, nil}
		right := &Node{15, nil, nil}
		root := &Node{10, left, right}
		newRoot := remove(root, 5)
		assert.Equal(t, root, newRoot)
		assert.Equal(t, 10, newRoot.value)
		assert.Equal(t, 15, newRoot.right.value)
		assert.Equal(t, (*Node)(nil), newRoot.left)
	})
}
