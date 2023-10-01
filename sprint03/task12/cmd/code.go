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
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	data := make([]int, 0, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		data = append(data, v)
	}

	sc.Scan()
	k, _ := strconv.Atoi(sc.Text())

	ub1 := lowerBound(data, k)
	ub2 := lowerBound(data, 2*k)
	if ub1 >= 0 {
		ub1++
	}
	if ub2>=0{
		ub2++
	}

	fmt.Printf("%d %d", ub1, ub2)
}

func lowerBound(data []int, k int) int {
	if len(data) == 0 {
		return -1
	}
	if data[0] >= k {
		return 0
	}

	begin := 0
	end := len(data)
	mid := len(data) / 2

	for end-begin > 1 {
		if data[mid] >= k {
			end = mid
		} else {
			begin = mid
		}
		mid = begin + (end-begin)/2
	}

	if data[begin] >= k {
		return begin
	} else if begin+1 < len(data) && data[begin+1] >= k {
		return begin + 1
	}

	return -1
}
