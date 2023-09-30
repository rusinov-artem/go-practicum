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
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	list := make([]int, 0, n)

	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		list = append(list, v)
	}

	fmt.Println(maxP(list))
}

func maxP(list []int) int {
	if len(list) < 3 {
		return 0
	}

	sort.Ints(list)

	for i := len(list) - 1; i >= 2; i-- {
		if isTriangle(list[i], list[i-1], list[i-2]) {
			return list[i] + list[i-1] + list[i-2]
		}
	}

	return 0
}

func isTriangle(a, b, c int) bool {
	return (a < (b + c)) &&
		(b < (a + c)) &&
		(c < (a + b))
}
