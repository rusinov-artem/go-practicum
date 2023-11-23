package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{key, nil, nil}
	}

	if root.value <= key {
		if root.right != nil {
			insert(root.right, key)
			return root
		}
		root.right = &Node{key, nil, nil}
	}

	if root.value > key {
		if root.left != nil {
			insert(root.left, key)
			return root
		}
		root.left = &Node{key, nil, nil}
	}

	return root
}
