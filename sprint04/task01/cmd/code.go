package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 4*1024*1024), 0)
	sc.Split(bufio.ScanLines)

	sc.Scan()
	a, _ := strconv.ParseInt(sc.Text(), 10, 64)
	am := uint64(a)

	sc.Scan()
	m, _ := strconv.ParseInt(sc.Text(), 10, 64)
	mm := uint64(m)

	sc.Scan()
	txt := sc.Text()

	fmt.Println(hash(txt, am, mm))
}

func hash(s string, a, m uint64) uint64 {
	hash := uint64(s[0]) % m
	for i := 1; i < len(s); i++ {
		hash = (hash*a + uint64(s[i])) % m
	}
	return hash
}
