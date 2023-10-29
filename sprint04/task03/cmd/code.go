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
	sc.Buffer(make([]byte, 0, 3*1024*1024), 0)
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	txt := sc.Text()

	subH := subhash(txt, int64(a), int64(m))

	sc.Scan()
	commandsCount, _ := strconv.Atoi(sc.Text())

	for i := 0; i < commandsCount; i++ {
		sc.Scan()
		firstI, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		lastI, _ := strconv.Atoi(sc.Text())

		if lastI - firstI == 0 {
			hash := uint64(txt[firstI-1]) % uint64(m)
			fmt.Println(hash)
			continue
		}

		if firstI == 1 {
			fmt.Println(subH[lastI-1].hash)
			continue
		}

		hash := (subH[lastI-1].hash - (subH[firstI-2].hash*subH[lastI-firstI+1].powA)% int64(m)) % int64(m)
		if hash < 0 {
			hash = int64(m) + hash
		}
		fmt.Println(hash)
	}
}

func hash(s string, a, m int64) int64 {
	hash := int64(s[0]) % m
	for i := 1; i < len(s); i++ {
		hash = (hash*a + int64(s[i])%m) % m
	}
	return hash
}

type SubHash struct {
	hash int64
	powA int64
}

func subhash(s string, a, m int64) []SubHash {
	sub := make([]SubHash, len(s))
	sub[0].hash = int64(s[0]) % m
	sub[0].powA = 1
	for i := 1; i < len(s); i++ {
		sub[i].hash = (sub[i-1].hash * a + int64(s[i]))%m
		sub[i].powA = (sub[i-1].powA * a) % m
	}
	return sub
}
