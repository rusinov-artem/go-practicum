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
	childernCount, _ := strconv.Atoi(sc.Text())

	greedy := make([]int, 0, childernCount)
	for i := 0; i < childernCount; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		greedy = append(greedy, v)
	}

	sc.Scan()
	cookiesCount, _ := strconv.Atoi(sc.Text())
	cookiesMap := make([]int, 1001)

	for i := 0; i < cookiesCount; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		cookiesMap[v]++
	}

	game := NewGame(greedy, cookiesMap, cookiesCount)
	fmt.Println(game.happy())

}

type game struct {
	greedy       []int
	cookiesMap   []int
	cookiesCount int
}

func NewGame(g, cm []int, cc int) *game {
	return &game{
		greedy:       g,
		cookiesMap:   cm,
		cookiesCount: cc,
	}
}

func (t *game) happy() int{
	count := 0
	cc := t.cookiesCount

	for _, g := range t.greedy {
		if cc <= 0 {
			return count
		}

		if t.cookiesMap[g] > 0 {
			t.cookiesMap[g]--
			cc--
			count++
			continue
		}

		for i := g+1 ; i < 1001 ; i++ {
			if t.cookiesMap[i] > 0 {
				t.cookiesMap[i]--
				cc--
				count++
				break
			}
		}
	}

	return count
}
