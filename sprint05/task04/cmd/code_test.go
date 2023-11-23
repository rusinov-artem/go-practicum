package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_(t *testing.T) {
	assert.True(t, true)
	t.Run("nil trees are the same", func(t *testing.T) {
		assert.True(t, isSame(nil, nil))
	})

	t.Run("same tree with one node", func(t *testing.T) {
		root1 := &Node{1, nil, nil}
		root2 := &Node{1, nil, nil}
		assert.True(t, isSame(root1, root2))
	})

	t.Run("not same tree with one node", func(t *testing.T) {
		root1 := &Node{2, nil, nil}
		root2 := &Node{1, nil, nil}
		assert.False(t, isSame(root1, root2))
	})

	t.Run("Same with left child", func(t *testing.T) {
		left1 := &Node{23, nil, nil}
		root1 := &Node{1, left1, nil}
		left2 := &Node{23, nil, nil}
		root2 := &Node{1, left2, nil}
		assert.True(t, isSame(root1, root2))
	})

	t.Run("not Same with left child", func(t *testing.T) {
		left1 := &Node{24, nil, nil}
		root1 := &Node{1, left1, nil}
		left2 := &Node{23, nil, nil}
		root2 := &Node{1, left2, nil}
		assert.False(t, isSame(root1, root2))
	})

	t.Run("not Same with left child 2", func(t *testing.T) {
		left1 := &Node{24, nil, nil}
		root1 := &Node{1, left1, nil}
		right2 := &Node{24, nil, nil}
		root2 := &Node{1, nil, right2}
		assert.False(t, isSame(root1, root2))
	})
}
