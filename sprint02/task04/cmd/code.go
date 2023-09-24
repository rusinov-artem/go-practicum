package main

// <template>
type ListNode struct {
	data string
	next *ListNode
}

// <template>

func Solution(head *ListNode, elem string) int {
	current := head
	i := 0
	for current != nil {
		if elem == current.data {
			return i
		}
		current = current.next
		i++
	}
	return -1
}

// func test() {
// 	node3 := ListNode{"node3", nil}
// 	node2 := ListNode{"node2", &node3}
// 	node1 := ListNode{"node1", &node2}
// 	node0 := ListNode{"node0", &node1}
// 	idx := Solution(&node0, "node9")
// 	// result is : idx == 2
// 	fmt.Println(idx)
// }

// func main() {
// 	test()
// }
