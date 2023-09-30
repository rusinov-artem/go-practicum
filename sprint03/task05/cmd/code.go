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
	buildingCount, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	balance, _ := strconv.Atoi(sc.Text())

	buildingList := make([]int, 0, buildingCount)
	for i := 0; i < buildingCount; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		buildingList = append(buildingList, v)
	}

	g := &game{
		balance:      balance,
		buildingList: buildingList,
	}

	fmt.Println(g.count())

}

type game struct {
	balance      int
	buildingList []int
}

func (t *game) count() int {
	if t.balance == 0 {
		return 0
	}

	if len(t.buildingList) == 0 {
		return 0
	}

	sort.Ints(t.buildingList)

	res := 0
	b := t.balance
	i := 0
	for b > 0 && i < len(t.buildingList) {
		if b >= t.buildingList[i] {
			b -= t.buildingList[i]
			res++
			i++
		} else {
			break
		}
	}
	return res
}
