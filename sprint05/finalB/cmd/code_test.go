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

	t.Run("can remove right child with children", func(t *testing.T) {
		left := &Node{5, nil, nil}
		rl := &Node{12, nil, nil}
		rr := &Node{20, nil, nil}
		right := &Node{15, rl, rr}
		root := &Node{10, left, right}
		//           10
		//        5        15
		//             12      20
		////////////////////////////////
		//           10
		//        5      20
		//             12
		newRoot := remove(root, 15)
		assert.Equal(t, root, newRoot)
		assert.Equal(t, 10, newRoot.value)
		assert.Equal(t, 5, newRoot.left.value)
		assert.Equal(t, 20, newRoot.right.value)
		assert.Equal(t, 12, newRoot.right.left.value)
		assert.Equal(t, (*Node)(nil), newRoot.right.right)
	})

	t.Run("can remove left child with children", func(t *testing.T) {
		ll := &Node{1, nil, nil}
		lr := &Node{7, nil, nil}
		left := &Node{5, ll, lr}
		right := &Node{15, nil, nil}
		root := &Node{10, left, right}
		//           10
		//        5        15
		//    1    7
		////////////////////////////////
		//           10
		//        7      15
		//      1
		newRoot := remove(root, 5)
		assert.Equal(t, root, newRoot)
		assert.Equal(t, 10, newRoot.value)
		assert.Equal(t, 15, newRoot.right.value)
		assert.Equal(t, 7, newRoot.left.value)
		assert.Equal(t, 1, newRoot.left.left.value)
		assert.Equal(t, (*Node)(nil), newRoot.left.right)
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

	t.Run("Complex test #1", func(t *testing.T) {
		node7 := &Node{6, nil, nil}
		node6 := &Node{8, node7, nil}
		node5 := &Node{10, node6, nil}
		node4 := &Node{2, nil, nil}
		node3 := &Node{3, node4, nil}
		node2 := &Node{1, nil, node3}
		node1 := &Node{5, node2, node5}
		newRoot := remove(node1, 10)
		//                 n1 (5)
		//      n2 (1)               n5 (10)
		//                     n6(8)
		//                n7(6)
		assert.Equal(t, node1, newRoot)
	})

	t.Run("remove unknown key", func(t *testing.T) {
		node10 := &Node{932, nil, nil}
		node9 := &Node{912, nil, node10}
		node8 := &Node{822, nil, nil}
		node7 := &Node{810, node8, node9}
		node6 := &Node{701, nil, nil}
		node5 := &Node{702, node6, node7}
		node4 := &Node{266, nil, nil}
		node3 := &Node{191, nil, node4}
		node2 := &Node{298, node3, nil}
		node1 := &Node{668, node2, node5}
		newRoot := remove(node1, 545)
		//                 n1 (5)
		//      n2 (1)               n5 (10)
		//                     n6(8)
		//                n7(6)
		assert.Equal(t, node1, newRoot)
	})
}
