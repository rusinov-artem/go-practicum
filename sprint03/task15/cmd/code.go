package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	k, d := fetchData()
	r := solution2(d, k)
	fmt.Println(r)
}

func fetchData() (int, []int) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 3*1024*1024), 0)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	d := make([]int, 0, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		d = append(d, v)
	}

	sc.Scan()
	k, _ := strconv.Atoi(sc.Text())
	return k, d
}

func solution2(d []int, k int) int {
	sort.Ints(d)

	low := 0
	high := d[len(d)-1] - d[0]
	for low < high {
		mi := (low + high) / 2
		count := before(d, mi)
		if count >= k {
			high = mi
		} else {
			low = mi + 1
		}
	}
	return low
}

func before(d []int, dist int) int {
	left := 0
	count := 0
	for right := 0; right < len(d); right++ {
		for d[right]-d[left] > dist {
			left++
		}
		count += right - left
	}
	return count
}
