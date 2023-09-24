package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 0, 10000), 0)

	a := &StackMax{
		data: make([]int, 0, 100),
	}

	for i := 0; i < n; i++ {
		sc.Scan()
		command := sc.Text()
		if command == "pop" {
			a.pop()
		} else if command == "push" {
			sc.Scan()
			val, _ := strconv.Atoi(sc.Text())
			a.push(val)
		} else if command == "get_max" {
			a.get_max()
		}
	}
}

type StackMax struct {
	data  []int
	max	[]int
}

func (a *StackMax) pop() {
	if len(a.data) == 0 {
		fmt.Println("error")
		return
	}

	if a.max[len(a.max)-1] == a.data[len(a.data)-1] {
		a.max = a.max[:len(a.max)-1]
	}

	a.data = a.data[:len(a.data)-1]
}

func (a *StackMax) push(v int) {
	a.data = append(a.data, v)
	if len(a.data) == 1 {
		a.max = append(a.max, v)
		return
	}

	if a.max[len(a.max)-1] <= v {
		a.max = append(a.max, v)
	}
}

func (a *StackMax) get_max() {
	if len(a.data) == 0 {
		fmt.Println("None")
		return
	}

	fmt.Println(a.max[len(a.max)-1])
}
