package main

// <template>
type ListNode struct {
	data string
	next *ListNode
	prev *ListNode
}

// <template>

func Solution(head *ListNode) *ListNode {
	var nHead *ListNode
	current := head
	for current != nil {
		nHead = current
		next := current.next
		current.next = current.prev
		current.prev = next
		current = next
	}
	return nHead
}

// func test() {
// 	node3 := ListNode{"node3", nil, nil}
// 	node2 := ListNode{"node2", &node3, nil}
// 	node1 := ListNode{"node1", &node2, nil}
// 	node0 := ListNode{"node0", &node1, nil}
// 	node3.prev = &node2
// 	node2.prev = &node1
// 	node1.prev = &node0
// 	newHead := Solution(&node0)
// 	printList(newHead)
// }

// func printList(node *ListNode) {
// 	current := node
// 	for current != nil {
// 		fmt.Println(current.data)
// 		current = current.next
// 	}
// }

// func main() {
// 	test()
// }
