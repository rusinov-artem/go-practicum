package main

import (
	"fmt"
)

func main() {
	var n uint64
	var k uint64
	fmt.Scanf("%d %d", &n, &k)
	m := pow10(k)
	fmt.Println(mfib(n, m))
}

func mfib(n, m uint64) uint64 {
	if n == 0 || n == 1 {
		return 1
	}

	prev := uint64(1)
	curr := uint64(1)
	next := uint64(0)

	for i := uint64(2); i <= n; i++ {
		next = (prev + curr) % m 
		prev = curr
		curr = next
	}

	return next
}

func pow10(k uint64) uint64 {
	if k == 0 {
		return 1
	}
	if k == 1 {
		return 10
	}
	v := uint64(10)
	for i := uint64(2); i <= k; i++ {
		v = v * 10
	}
	return v
}
