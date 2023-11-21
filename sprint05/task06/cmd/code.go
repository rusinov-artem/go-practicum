package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func Solution(root *Node) int {
	return countDepth(root, 0)
}

func countDepth(root *Node, depth int) int {
	if root == nil {
		return depth
	}

	if root != nil {
		depth++
	}

	ldepth := countDepth(root.left, depth)
	rdepth := countDepth(root.right, depth)

	if ldepth > rdepth {
		return ldepth
	}
	return rdepth
}
