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

	dfs(gr, startVertex, func(v int) {
		fmt.Printf("%d ", v)
	})
}

func dfs(gr [][]int, vertex int, fn func(vertex int)) {
	m := make([]bool, len(gr))
	dfsImpl(gr, m, vertex, fn)
}

func dfsImpl(gr [][]int, m []bool, vertex int, fn func(vertex int)) {
	if m[vertex] {
		return
	}

	fn(vertex)

	m[vertex] = true
	for i := 0; i < len(gr[vertex]); i++ {
		dfsImpl(gr, m, gr[vertex][i], fn)
	}
}
