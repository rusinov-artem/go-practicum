package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 0, 4*1024*1024), 0)
	sc.Scan() // число вершин
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan() // число ребер
	m, _ := strconv.Atoi(sc.Text())

	gr := make([][]int, n+1)

	for i := 0; i < m; i++ {
		sc.Scan()
		u, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())

		gr[u] = append(gr[u], v)
		gr[v] = append(gr[v], u)
	}

	for i := 0; i < n+1; i++ {
		sort.Slice(gr[i], func(a, b int) bool {
			return gr[i][a] < gr[i][b]
		})
	}

	sc.Scan()
	startVertex, _ := strconv.Atoi(sc.Text())

	bfs(gr, startVertex, func(v int) {
		fmt.Printf("%d ", v)
	})
}

func bfs(gr [][]int, startVertex int, fn func(v int)) {
	q := &MyQueue{}
	m := make([]bool, len(gr))
	q.PushBack(startVertex)
	m[startVertex] = true

	for !q.Empty() {
		v := q.PopFront().(int)
		fn(v)
		for _, nv := range gr[v] {
			if !m[nv] {
				q.PushBack(nv)
				m[nv] = true
			}
		}
	}
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
