package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_(t *testing.T) {
	assert.True(t, true)
	t.Run("nil", func(t *testing.T) {
		newHead := insert(nil, 123)
		assert.Equal(t, 123, newHead.value)
	})

	t.Run("single node", func(t *testing.T) {
		root := &Node{55, nil, nil}
		newHead := insert(root, 55)
		assert.Equal(t, root, newHead)
		assert.Equal(t, 55, newHead.value)
		assert.Equal(t, 55, newHead.right.value)
	})

	t.Run("single node with right child", func(t *testing.T) {
		right := &Node{55, nil, nil}
		root := &Node{55, nil, right}
		newHead := insert(root, 55)
		assert.Equal(t, root, newHead)
		assert.Equal(t, 55, newHead.value)
		assert.Equal(t, 55, newHead.right.value)
		assert.Equal(t, 55, newHead.right.right.value)
	})

   t.Run("single node with right child. Insert to the left", func(t *testing.T) {
       right := &Node{55, nil, nil}
       root := &Node{55, nil, right}
       newHead := insert(root, 3)
       assert.Equal(t, root, newHead)
       assert.Equal(t, 55, newHead.value)
       assert.Equal(t, 3, newHead.left.value)
   })
}
