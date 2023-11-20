package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	fmt.Println(katalan(int64(n)))
}

func katalan(n int64) int64 {
	d :=int64(1)
	c := int64(1)
	for i := int64(2) ; i <= n ; i++ {
		d = d * i
		c = c * (2*n-i+2)
		if c %d == 0 {
			c = c / d
			d = 1
		}
	}
	return c / d
}

func factorial(n int64) int64 {
	f := int64(1)
	for i := n; i > 1; i-- {
		f = f * i
	}
	return f
}

func partFactorial(n int64, b int64) int64 {
	f := int64(1)
	for i := n; i > b; i-- {
		f = f * i
	}
	return f
}

func fib(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}

	a := 1
	b := 1

	for i := 2; i < n; i++ {
		t := b
		b = a + b
		a = t
	}

	return b
}
