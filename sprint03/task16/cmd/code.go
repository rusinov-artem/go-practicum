package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 1024*1024*3), 0)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	data := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		data[i], _ = strconv.Atoi(sc.Text())
	}

	fmt.Println(maxBlox(data))
}

func maxBlox(d []int) int {
	if len(d) <2 {
		return len(d)
	}

	count :=0
	combod := make([]int,len(d))
	begin := 0
	for i := 0; i<len(d); i++ {
		if begin == i && d[i]==i {
			count++
			begin++
		}
		combod[d[i]] = 1 
		if combo(combod, begin, i+1) {
			count++
			begin=i
		}

	}
	return count
}

func combo(d []int, begin, end int) bool {
	if begin >= end {
		return false
	}
	for i := begin ; i != end ; i++ {
		if d[i]==0 {
			return false
		}
	}	
	return true
}
