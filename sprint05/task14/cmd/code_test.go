package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_(t *testing.T) {
	assert.True(t, true)

	t.Run("ExampleTest", func(t *testing.T) {
		//                  5(6)
		//         2(2)             10(3)
		//             3(1)     8(1)       11 (1)
		node1 := &Node{3, nil, nil, 1}
		node2 := &Node{2, nil, node1, 2}
		node3 := &Node{8, nil, nil, 1}
		node4 := &Node{11, nil, nil, 1}
		node5 := &Node{10, node3, node4, 3}
		node6 := &Node{5, node2, node5, 6}
		left, right := split(node6, 4)
		assert.Equal(t, 4, size(left))
		assert.Equal(t, 2, size(right))
	})

	t.Run("Can split1 for 3 node balanced bst", func(t *testing.T) {
		left := &Node{1, nil, nil, 1}
		right := &Node{3, nil, nil, 1}
		root := &Node{2, left, right, 3}

		leftTree, rightTree := split(root, 1)
		assert.Equal(t, left, leftTree)
		assert.Equal(t, 1, leftTree.size)
		assert.Equal(t, root, rightTree)
		assert.Equal(t, 2, rightTree.size)
	})

	t.Run("Can split2 for 3 node balanced bst", func(t *testing.T) {
		left := &Node{1, nil, nil, 1}
		right := &Node{3, nil, nil, 1}
		root := &Node{2, left, right, 3}

		leftTree, rightTree := split(root, 2)
		assert.Equal(t, root, leftTree)
		assert.Equal(t, 2, leftTree.size)
		assert.Equal(t, right, rightTree)
		assert.Equal(t, 1, rightTree.size)
	})

	t.Run("Can split3 for 3 node balanced bst", func(t *testing.T) {
		left := &Node{1, nil, nil, 1}
		right := &Node{3, nil, nil, 1}
		root := &Node{2, left, right, 3}

		leftTree, rightTree := split(root, 3)
		assert.Equal(t, root, leftTree)
		assert.Equal(t, (*Node)(nil), rightTree)
	})

	t.Run("Can split1 node with left child", func(t *testing.T) {
		right := &Node{3, nil, nil, 1}
		root := &Node{2, nil, right, 2}
		leftTree, rightTree := split(root, 1)
		assert.Equal(t, root, leftTree)
		assert.Equal(t, 1, leftTree.size)
		assert.Equal(t, right, rightTree)
		assert.Equal(t, 1, rightTree.size)
	})

	t.Run("Can split2 node with left child", func(t *testing.T) {
		right := &Node{3, nil, nil, 1}
		root := &Node{2, nil, right, 2}
		leftTree, rightTree := split(root, 2)
		assert.Equal(t, root, leftTree)
		assert.Equal(t, 2, leftTree.size)
		assert.Equal(t, (*Node)(nil), rightTree)
		assert.Equal(t, 0, size(rightTree))
	})

	t.Run("Can split1 node with 2left child", func(t *testing.T) {
		leftLeft := &Node{1, nil, nil, 1}
		left := &Node{2, leftLeft, nil, 2}
		root := &Node{3, left, nil, 3}
		leftTree, rightTree := split(root, 1)
		assert.Equal(t, leftLeft, leftTree)
		assert.Equal(t, 1, size(leftTree))
		assert.Equal(t, root, rightTree)
		assert.Equal(t, 2, size(rightTree))
	})

	t.Run("Can split1 node with left and leftRight child", func(t *testing.T) {
		leftRight := &Node{25, nil, nil, 1}
		left := &Node{20, nil, leftRight, 2}
		root := &Node{30, left, nil, 3}
		leftTree, rightTree := split(root, 1)
		assert.Equal(t, left, leftTree)
		assert.Equal(t, 1, size(leftTree))
		assert.Equal(t, root, rightTree)
		assert.Equal(t, 2, size(rightTree))

		assert.Equal(t, 30, root.value)
		assert.Equal(t, 25, root.left.value)
	})

	t.Run("Can split2 node with left and leftRight child", func(t *testing.T) {
		leftRight := &Node{25, nil, nil, 1}
		left := &Node{20, nil, leftRight, 2}
		root := &Node{30, left, nil, 3}
		leftTree, rightTree := split(root, 2)
		assert.Equal(t, left, leftTree)
		assert.Equal(t, 2, size(leftTree))
		assert.Equal(t, root, rightTree)
		assert.Equal(t, 1, size(rightTree))

		assert.Equal(t, 20, leftTree.value)
		assert.Equal(t, 25, leftTree.right.value)

		assert.Equal(t, 30, rightTree.value)
	})
}
