package main

import "math"

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}
// <template>

func Solution(root *Node) bool {
	return isBalanced(root)
}

func isBalanced(root *Node) bool {
	maxDepth := 0
	minDepth := math.MaxInt
	iterateDepth(root, func(depth int) {
		if maxDepth < depth {
			maxDepth = depth
		}
		if depth < minDepth {
			minDepth = depth
		}
	})

	return maxDepth-minDepth <= 1
}

func iterateDepth(root *Node, fn func(depth int)) {
	iterateDepthImpl(root, 0, fn)
}

func iterateDepthImpl(root *Node, d int, fn func(depth int)) {
	if root == nil {
		fn(d)
		return
	}

	d++

	iterateDepthImpl(root.left, d, fn)
	iterateDepthImpl(root.right, d, fn)
}
