package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}
// <template>

func Solution(root *Node) int {
	return maxScore(root)
}

func maxScore(root *Node) int {
	if root == nil {
		return 0
	}

	m := root.value
	iterateTree(root, func(node *Node) {
		v := scorePath(node)
		if m < v {
			m = v
		}
	})

	return m
}

func iterateTree(root *Node, fn func(node *Node)) {
	if root == nil {
		return
	}

	fn(root)

	iterateTree(root.left, fn)
	iterateTree(root.right, fn)
}

func scorePath(root *Node) int {
	if root == nil {
		return 0
	}

	left := scoreBranch(root.left)
	right := scoreBranch(root.right)

	values := []int{root.value}
	if root.left != nil {
		values = append(values, left)
		values = append(values, left + root.value)
	}

	if root.right != nil {
		values = append(values, right)
		values = append(values, right + root.value)
	}

	if root.right != nil && root.left != nil {
		values = append(values, right + left + root.value)
	}

	return max(values)
}

func scoreBranch(root *Node) int {
	if root == nil {
		return 0
	}


	m := root.value
	iterateTreeWithSum(root, func(sum int) {
		if m < sum {
			m = sum
		}	
	})

	return m
}

func iterateTreeWithSum(root *Node, fn func(sum int)) {
	if root == nil {
		return
	}

	iterateTreeWithSomImpl(root, 0, fn)
}

func iterateTreeWithSomImpl(root *Node, start int, fn func(sum int)) {
	if root == nil {
		fn(start)
		return
	}

	fn(start + root.value)
	iterateTreeWithSomImpl(root.left, start+root.value, fn)
	iterateTreeWithSomImpl(root.right, start+root.value, fn)
}


func max(d []int) int {
	m := d[0]
	for i := 1; i < len(d); i++ {
		if m < d[i] {
			m = d[i]
		}
	}
	return m
}
