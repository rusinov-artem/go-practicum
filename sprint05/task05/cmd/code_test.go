package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ChekTreeIsBST(t *testing.T) {
	t.Run("empty tree is ok", func(t *testing.T) {
		assert.True(t, Solution(nil))
	})

	t.Run("tree without children is ok", func(t *testing.T) {
		root := &Node{0, nil, nil}
		assert.True(t, Solution(root))
	})

	t.Run("if root has 2 children", func(t *testing.T) {
		leftChild := &Node{-1, nil, nil}
		rightChild := &Node{2, nil, nil}
		root := &Node{0, leftChild, rightChild}
		assert.True(t, Solution(root))
	})

	t.Run("if root has 2 children. left is eq to root", func(t *testing.T) {
		leftChild := &Node{0, nil, nil}
		rightChild := &Node{2, nil, nil}
		root := &Node{0, leftChild, rightChild}
		assert.False(t, Solution(root))
	})

	t.Run("if root has 2 children. left is gt  root", func(t *testing.T) {
		leftChild := &Node{2, nil, nil}
		rightChild := &Node{2, nil, nil}
		root := &Node{0, leftChild, rightChild}
		assert.False(t, Solution(root))
	})

	t.Run("if root has 2 children. right is eq root", func(t *testing.T) {
		leftChild := &Node{-1, nil, nil}
		rightChild := &Node{0, nil, nil}
		root := &Node{0, leftChild, rightChild}
		assert.False(t, Solution(root))
	})

	t.Run("if root has 2 children. right is lt root", func(t *testing.T) {
		leftChild := &Node{-1, nil, nil}
		rightChild := &Node{-1, nil, nil}
		root := &Node{0, leftChild, rightChild}
		assert.False(t, Solution(root))
	})

	t.Run("if root has 2 children. and left child has 2 children", func(t *testing.T) {
		llChild := &Node{-6, nil, nil}
		lrChild := &Node{-3, nil, nil}
		leftChild := &Node{-5, llChild, lrChild}
		rightChild := &Node{1, nil, nil}
		root := &Node{0, leftChild, rightChild}
		assert.True(t, Solution(root))
	})

	t.Run("if root has 2 children. and left child has wrong left child(eq)", func(t *testing.T) {
		llChild := &Node{-5, nil, nil}
		lrChild := &Node{-3, nil, nil}
		leftChild := &Node{-5, llChild, lrChild}
		rightChild := &Node{1, nil, nil}
		root := &Node{0, leftChild, rightChild}
		assert.False(t, Solution(root))
	})

	t.Run("if root has 2 children. and left child has wrong left child(gt)", func(t *testing.T) {
		llChild := &Node{5, nil, nil}
		lrChild := &Node{-3, nil, nil}
		leftChild := &Node{-5, llChild, lrChild}
		rightChild := &Node{1, nil, nil}
		root := &Node{0, leftChild, rightChild}
		assert.False(t, Solution(root))
	})

	t.Run("if root has 2 children. and left child has wrong right child(eq)", func(t *testing.T) {
		llChild := &Node{-6, nil, nil}
		lrChild := &Node{-5, nil, nil}
		leftChild := &Node{-5, llChild, lrChild}
		rightChild := &Node{1, nil, nil}
		root := &Node{0, leftChild, rightChild}
		assert.False(t, Solution(root))
	})

	t.Run("if root has 2 children. and left child has wrong right child(lt)", func(t *testing.T) {
		llChild := &Node{-6, nil, nil}
		lrChild := &Node{-6, nil, nil}
		leftChild := &Node{-5, llChild, lrChild}
		rightChild := &Node{1, nil, nil}
		root := &Node{0, leftChild, rightChild}
		assert.False(t, Solution(root))
	})

	t.Run("if root has 2 children. and right child has wrong left child(over interval)", func(t *testing.T) {
		rlChild := &Node{-5, nil, nil}
		rrChild := &Node{20, nil, nil}
		leftChild := &Node{-5, nil, nil}
		rightChild := &Node{10, rlChild, rrChild}
		root := &Node{0, leftChild, rightChild}
		assert.False(t, Solution(root))
	})
}

func Test_Debug(t *testing.T) {
	rlChild := &Node{-5, nil, nil}
	rrChild := &Node{20, nil, nil}
	leftChild := &Node{-5, nil, nil}
	rightChild := &Node{10, rlChild, rrChild}
	root := &Node{0, leftChild, rightChild}
	assert.False(t, Solution(root))
}
