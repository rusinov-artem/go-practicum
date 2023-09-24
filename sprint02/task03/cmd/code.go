package main

import "fmt"

// <template>
type ListNode struct {
	data string
	next *ListNode
}

// <template>

func Solution(head *ListNode, idx int) *ListNode {
	nHead := head
	var prev *ListNode
	current := head
	for i := 0; i < idx; i++ {
		if current != nil {
			prev = current
			current = current.next
		} else {
			break
		}
	}

	if prev == nil {
		nHead = head.next
		return nHead
	}

	if current == nil {
		return nHead
	}

	prev.next = current.next

	return nHead
}

func test() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	newHead := Solution(&node0, 0)
	// result is : node0 -> node2 -> node3
	printList(newHead)
}

func printList(node *ListNode) {
	current := node
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

// func main() {
// 	test()
// }
