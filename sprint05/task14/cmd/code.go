package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
	size  int
}

// <template>

func split(node *Node, k int) (*Node, *Node) {
	if node == nil {
		return nil, nil
	}

	if k == 0 {
		return nil, node
	}

	if node.size == k {
		return node, nil
	}

	if size(node.left) >= k {
		node.size -= k
		left, right := split(node.left, k)
		node.left = right
		return left, node
	}

	if size(node.left)+1 == k {
		node.size -= size(node.right)
		right := node.right
		node.right = nil
		return node, right
	}

	if size(node.left)+1 < k {
		left := node
		left.size -= size(node.right)
		l, r := split(node.right, k-size(left))
		left.right = l
		left.size += size(l)
		return left, r
	}

	return nil, nil
}

func size(n *Node) int {
	if n == nil {
		return 0
	}
	return n.size
}
