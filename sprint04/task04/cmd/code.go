package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 4*1024*1024), 0)
	sc.Split(bufio.ScanLines)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	out := make([]string, n)
	m := make(map[string]struct{})
	oi :=0
	for i := 0; i < n; i++ {
		sc.Scan()
		t := sc.Text()
		if _, ok := m[t]; !ok {
			m[t] = struct{}{}
			out[oi] = t
			oi++
		}
	}

	for i := 0; i < len(m); i++ {
		fmt.Println(out[i])
	}
}
