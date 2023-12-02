package main

import (
	"bufio"
	"fmt"
	"os"
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

		if len(gr[u]) < n+1 {
			gr[u] = make([]int, n+1)
		}

		gr[u][v] = 1
	}

	for i := 1; i < n+1; i++ {
		if len(gr[i]) < n+1 {
			gr[i] = make([]int, n+1)
		}
		fmt.Printf("%d", gr[i][1])
		for j := 2; j < n+1; j++ {
			fmt.Printf(" %d", gr[i][j])
		}
		fmt.Println("")
	}
}
