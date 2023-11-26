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
		sc.Scan() // число вершин
		u, _ := strconv.Atoi(sc.Text())
		sc.Scan() // число вершин
		v, _ := strconv.Atoi(sc.Text())

		gr[u] = append(gr[u], v)
	}

	for i := 1; i < n+1; i++ {
		sort.Slice(gr[i], func(i, j int) bool {
			return i < j
		})
		fmt.Printf("%d", len(gr[i]))
		for j := 0; j < len(gr[i]); j++ {
			fmt.Printf(" %d", gr[i][j])
		}
		fmt.Println("")
	}
}
