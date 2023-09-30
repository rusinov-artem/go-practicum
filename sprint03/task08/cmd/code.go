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

	words := make([]string, 0, 10)

	for i := 0; i < n; i++ {
		sc.Scan()
		v := sc.Text()
		words = append(words, v)
	}

	sort.Slice(words, func(l, r int) bool {
		return words[l]+words[r] > words[r]+words[l]
	})

	for i := 0; i < len(words); i++ {
		fmt.Print(words[i])
	}
	fmt.Println("")

}
