package main

// <template>
type Node struct {
	value int
	left *Node
	right *Node
}
// <template>

func Solution(root *Node) bool {
	return solImpl(root, nil, nil)
}

func solImpl(root *Node, min *int, max *int) bool {
	if root == nil {
		return true
	}

	if root.left != nil {
		if root.left.value >= root.value {
			return false
		}

		if min != nil {
			if root.left.value < *min {
				return false
			}
		}

		if max != nil {
			if root.left.value > *max {
				return false
			}
		}
	}

	if root.right != nil {
		if root.right.value <= root.value {
			return false
		}

		if min != nil {
			if root.right.value < *min {
				return false
			}
		}

		if max != nil {
			if root.right.value > *max {
				return false
			}
		}
	}

	return solImpl(root.left, min, &root.value) && solImpl(root.right, &root.value, max)
}

