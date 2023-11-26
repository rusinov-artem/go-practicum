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
// Алгоритм расходует дополнительную память на 
// рекурсинвых вызовах, поэтом сложность по памяти тоже 
// O(h)
// Корректность алгоритма подтвержена тестами в соседнем файле
func remove(node *Node, key int) *Node {
	if node == nil {
		return nil
	}

	if node.value == key {
		if node.left == nil && node.right == nil {
			return nil
		}

		if node.left != nil && node.right != nil {
			left := node.left
			right := node.right
			node.left = nil
			node.right = nil
			insert(right, left)
			return right
		}

		if node.left != nil {
			left := node.left
			node.left = nil
			return left
		} else {
			right := node.right
			node.right = nil
			return right
		}
	}

	if key < node.value {
		node.left = remove(node.left, key)
		return node
	} else {
		node.right = remove(node.right, key)
		return node
	}
}

func insert(root, node *Node) {
	if node.value < root.value {
		if root.left == nil {
			root.left = node
		} else {
			insert(root.left, node)
		}
	} else {
		if root.right == nil {
			root.right = node
		} else {
			insert(root.right, node)
		}
	}
}
