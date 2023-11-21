package main

import "fmt"

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func printRange(root *Node, left int, right int) {
	iterateRange(root, left, right, func(val int) {
		fmt.Printf("\n%d", val)
	})
}

type Action func(val int)

func iterateRange(root *Node, left, right int, fn Action) {
	if root == nil {
		return
	}

	if root.value >= left {
		iterateRange(root.left, left, right, fn)
	}

	if root.value >= left && root.value <= right {
		fn(root.value)
	}

	if root.value <= right {
		iterateRange(root.right, left, right, fn)
	}
}
