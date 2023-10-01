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
	sc.Buffer(make([]byte, 0, 6*1024*1024), 0)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	confs := make([]conf, 100001)

	for i := range confs {
		confs[i].id = i
	}

	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		confs[v].count++
	}

	sort.Slice(confs, func(l, r int) bool {
		if confs[l].count == confs[r].count {
			return confs[l].id < confs[r].id
		}
		return confs[l].count > confs[r].count
	})

	sc.Scan()
	k, _ := strconv.Atoi(sc.Text())

	if k > 0 {
		fmt.Print(confs[0].id)
	}

	for i := 1; i < k && i < len(confs); i++ {
		if confs[i].count == 0 {
			break
		}
		fmt.Printf(" %d", confs[i].id)
	}

	fmt.Println("")
}

type conf struct {
	id    int
	count int
}
