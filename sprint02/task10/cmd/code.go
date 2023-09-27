package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 10000), 0)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	q := ListQueue{}

	for i := 0; i < n; i++ {
		sc.Scan()
		cmd := sc.Text()
		if cmd == "get" {
			v, err := q.get()
			if err != nil {
				fmt.Println("error")
			} else {
				fmt.Println(v)
			}
		} else if cmd == "put" {
			sc.Scan()
			v, _ := strconv.Atoi(sc.Text())
			q.put(v)
		} else if cmd == "size" {
			s := q.size
			fmt.Println(s)
		}
	}
}

type ListQueue struct {
	head *ListQueueNode
	tail *ListQueueNode
	size int
}

type ListQueueNode struct {
	data int
	next *ListQueueNode
}

func (t *ListQueue) put(v int) {
	node := &ListQueueNode{
		data: v,
		next: nil,
	}

	if t.head == nil {
		t.head = node
		t.tail = node
	} else {
		t.tail.next = node
		t.tail = node
	}

	t.size++
}

func (t *ListQueue) get() (int, error) {

	if t.head == nil {
		return 0, errors.New("empty queue")
	}

	v := t.head.data
	t.size--
	t.head = t.head.next

	return v, nil
}
