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
	sc.Buffer(make([]byte, 0, 5*1024*1024), 0)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	data1 := make([]int, 0, n)
	data2 := make([]int, 0, m)

	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		data1 = append(data1, v)

	}

	for i := 0; i < m; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		data2 = append(data2, v)

	}

	merged := merge(data1, data2)
	fmt.Println(merged)

	mid := len(merged) / 2   
	if len(merged)%2 == 0 {
		fmt.Println((float64(merged[mid])+float64(merged[mid-1])) / 2.0)
	} else {
		fmt.Println(merged[mid])
	}

}

func merge(d1, d2 []int) []int {
	res := make([]int, 0, len(d1)+len(d2))

	d1begin := 0
	d1end := len(d1)
	d2begin := 0
	d2end := len(d2)

	for d1begin < d1end || d2begin < d2end {
		if d1begin >= d1end {
			res = append(res, d2[d2begin])
			d2begin++
		} else if d2begin >= d2end {
			res = append(res, d1[d1begin])
			d1begin++
		} else if d1[d1begin] > d2[d2begin] {
			res = append(res, d2[d2begin])
			d2begin++
		} else {
			res = append(res, d1[d1begin])
			d1begin++
		}
	}

	return res
}
