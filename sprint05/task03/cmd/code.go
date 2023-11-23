package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func Solution(root *Node) bool {
	return isPalindrome(root)
}

func isPalindrome(root *Node) bool {
	if root == nil {
		return true
	}

	if root.left == nil && root.right == nil {
		return true
	}

	itIsPalindrome := true
	iterateLevels(root, func(vals []*int) bool {
		if !checkPalindrome(vals) {
			itIsPalindrome = false
			return false
		}
		return true
	})

	return itIsPalindrome
}

func checkPalindrome(vals []*int) bool {
	if len(vals) < 2 {
		return true
	}

	begin := 0
	end := len(vals)

	for begin < end {
		if vals[begin] != nil && vals[end-1] != nil {
			if *vals[begin] != *vals[end-1] {
				return false
			}
		} else {
			if vals[begin] != vals[end-1] {
				return false
			}
		}
		begin++
		end--
	}

	return true
}

func iterateLevels(root *Node, fn func(vals []*int) bool) {
	if root == nil {
		return
	}

	q := &MyQueue{}
	q.PushBack(NodeWrap{root, 0})
	level := 0
	values := []*int{}
	iterateLevels2(q, func(l int, v *int) bool {
		goNext := true
		if level == l {
			values = append(values, v)
		} else {
			goNext = fn(values)
			level = l
			values = []*int{v}
		}
		return goNext
	})
}

func iterateLevels2(q *MyQueue, fn func(level int, val *int) bool) {
	for !q.Empty() {
		root := q.PopFront().(NodeWrap)
		if root.n == nil {
			fn(root.level, nil)
			continue
		}

		if !fn(root.level, &root.n.value) {
			return
		}

		q.PushBack(NodeWrap{root.n.left, root.level + 1})
		q.PushBack(NodeWrap{root.n.right, root.level + 1})
	}
}

type NodeWrap struct {
	n     *Node
	level int
}

type MyListNode struct {
	data interface{}
	next *MyListNode
}

type MyQueue struct {
	head *MyListNode
	tail *MyListNode
}

func (this *MyQueue) Empty() bool {
	return this.head == nil
}

func (this *MyQueue) PushBack(val interface{}) {
	if this.tail == nil {
		this.tail = &MyListNode{val, nil}
		this.head = this.tail
		return
	}

	this.tail.next = &MyListNode{val, nil}
	this.tail = this.tail.next
}

func (this *MyQueue) PopFront() interface{} {
	v := this.head.data
	if this.head == this.tail {
		this.tail = nil
	}
	this.head = this.head.next
	return v
}
