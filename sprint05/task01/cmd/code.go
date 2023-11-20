package main


// <template>
type Node struct {  
	value  int  
	left   *Node  
	right  *Node  
}
// <template>

func Solution(root *Node) int {
	return findMax(root)
}

func findMax(root *Node) int {
	max := root.value

	if root.left != nil {
		lmax := findMax(root.left)
		if max < lmax {
			max = lmax
		}
	}

	if root.right != nil {
		rmax := findMax(root.right)
		if max < rmax {
			max = rmax
		}
	}


	return max

}
