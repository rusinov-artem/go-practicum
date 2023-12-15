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
	graph := NewGraph(n)

	sc.Scan()
	l, _ := strconv.Atoi(sc.Text())

	for i := 0; i < l; i++ {
		sc.Scan()
		u, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		graph.AddEdge(u, v)
	}

	graph.Normalize()
	n, vertex := components(graph)
	printComponents(n, vertex)
}

func printComponents(n int, vertex []int) {
	fmt.Println(n)
	for i := 1; i <= n; i++ {
		for j := 0; j < len(vertex); j++ {
			if vertex[j] == i {
				fmt.Printf("%d ", j+1)
			}
		}
		fmt.Println()
	}
}

func components(graph *MyGraph) (int, []int) {
	vertex := make([]int, len(graph.Data))

	componentN := 0
	for i := 0; i < len(vertex); i++ {
		if vertex[i] > 0 {
			continue
		}
		componentN++
		dfsMark(graph, vertex, i, componentN)
	}
	return componentN, vertex
}

func dfsMark(graph *MyGraph, vertex []int, start, component int) {
	if vertex[start] != 0 {
		return
	}

	vertex[start] = component

	for i := 0; i < len(graph.Data[start]); i++ {
		dfsMark(graph, vertex, graph.Data[start][i], component)
	}
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
