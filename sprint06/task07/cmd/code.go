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

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	g := NewGraph(n)

	sc.Scan()
	l, _ := strconv.Atoi(sc.Text())
	for i := 0; i < l; i++ {
		sc.Scan()
		u, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		g.AddEdge(u, v)
	}

	sc.Scan()
	from, _ := strconv.Atoi(sc.Text())

	g.Normalize()
	fmt.Println(maxPathLen(g, from))
}

func maxPathLen(g *MyGraph, from int) int {
	from--

	maxL := 0
	q := &MyQueue{}

	vertex := make([]int, len(g.Data))
	for i := 0; i < len(vertex); i++ {
		vertex[i] = -1
	}

	q.PushBack(from)
	vertex[from] = 0
	for !q.Empty() {
		v := q.PopFront().(int)
		for i := 0; i < len(g.Data[v]); i++ {
			u := g.Data[v][i]
			if vertex[u] != -1 {
				continue
			}
			vertex[u] = vertex[v] + 1
			if maxL < vertex[u] {
				maxL = vertex[u]
			}
			q.PushBack(u)
		}
	}
	return maxL
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

type MyGraph struct {
	Data [][]int
}

func (this *MyGraph) Normalize() {
	for i := 0; i < len(this.Data); i++ {
		sort.Slice(this.Data[i], func(l, r int) bool {
			return this.Data[i][l] < this.Data[i][r]
		})
	}
}

func (this *MyGraph) Print() {
	for i := 0; i < len(this.Data); i++ {
		fmt.Printf("%d => {", i+1)
		if len(this.Data[i]) > 0 {
			fmt.Printf("%d", this.Data[i][0]+1)
		}
		for j := 1; j < len(this.Data[i]); j++ {
			fmt.Printf(" %d", this.Data[i][j]+1)
		}
		fmt.Print("}\n")
	}
}

func (this *MyGraph) AddEdge(u, v int) {
	u = u - 1
	v = v - 1
	if len(this.Data[u]) == 0 {
		this.Data[u] = []int{v}
	} else {
		this.Data[u] = append(this.Data[u], v)
	}

	if len(this.Data[v]) == 0 {
		this.Data[v] = []int{u}
	} else {
		this.Data[v] = append(this.Data[v], u)
	}
}

func NewGraph(n int) *MyGraph {
	return &MyGraph{
		Data: make([][]int, n),
	}
}
