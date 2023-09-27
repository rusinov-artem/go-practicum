package main

import "fmt"

func main() {
	var n uint64
	fmt.Scanf("%d", &n)
	fmt.Println(mfib(n))
}

func mfib(n uint64) uint64 {
	if n == 0 || n == 1 {
		return 1
	}

	prev := uint64(1)
	curr := uint64(1)
	next := uint64(0)

	for i := uint64(2); i <= n ; i++ {
		next = prev + curr
		prev = curr
		curr = next
	}


	return next
}
