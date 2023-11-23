package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func Solution(root *Node) int {
	return sumBranch(root)
}

func sumBranch(root *Node) int {
	if root == nil {
		return 0
	}
    sum := 0
    iterateLeafes(root, 0, func(val int) {
      sum += val
    })
	return sum
}

func iterateLeafes(root *Node, val int, fn func(val int)) {
	if root == nil {
		return
	}

	if root.left == nil && root.right == nil {
		fn(val*10 + root.value)
		return
	}

	iterateLeafes(root.left, val*10+root.value, fn)
	iterateLeafes(root.right, val*10+root.value, fn)
}
