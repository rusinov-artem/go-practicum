package main

// <templatet>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

// Алгоритм удленияя ключа из BST
// Алгоритмическая сложность O(h) где h - высот дерева
// Алгоритм не расходует допольнительной памяти. Сложность
// по памяти O(1)
// Корректность алгоритма подтвержена тестами в соседнем файле
func remove(node *Node, key int) *Node {
	root := node
	var prev *Node

	for {
		if node == nil && root == nil {
			return nil
		}

		if node == nil {
			return root
		}

		if node.value == key {
			if root == node {
				right := root.right
				left := root.left
				root.left = nil
				root.right = nil

				if left == nil && right == nil {
					return nil
				}

				if left != nil && right != nil {
					root = right
					insert(right, left)
					return root
				}

				if right != nil {
					return right
				}
				if left != nil {
					return left
				}
			}

			if node.left == nil && node.right == nil {
				removeFrom(prev, node)
				return root
			}

			left := node.left
			right := node.right
			node.left = nil
			node.right = nil
			removeFrom(prev, node)
			insert(prev, right)
			insert(prev, left)
			return root
		}

		prev = node
		if node.value < key {
			node = prev.right
		} else {
			node = prev.left
		}
	}
}

func removeFrom(parent, child *Node) {
	if parent.left == child {
		parent.left = nil
	} else {
		parent.right = nil
	}
}

func insert(root, node *Node) {
	if node == nil {
		return
	}

	for {
		if node.value < root.value {
			if root.left == nil {
				root.left = node
				return
			} else {
				root = root.left
			}
		} else {
			if root.right == nil {
				root.right = node
				return
			} else {
				root = root.right
			}
		}
	}
}
