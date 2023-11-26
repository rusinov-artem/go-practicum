package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Алгоритм heap sort
// Алгоритмическая сложность O(n*log(n))
// Алгоритм использует дополнительную память на рекурсинвый вызовах, 
// оценка дополнительной памяти O(log(n))
// Корректность работы продемострирована тестами в соседнем файле
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 0, 4*1024*1024), 0)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	users := []User{}
	for i := 0; i < n; i++ {
		sc.Scan()
		name := sc.Text()
		sc.Scan()
		score, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		penalty, _ := strconv.Atoi(sc.Text())
		users = append(users, User{
			Name:    name,
			Score:   score,
			Penalty: penalty,
		})
	}

	c := &SortableIndex{
		Index: iRange(n),
	}

	c.Less = func(i, j int) bool {
		i = c.Index[i]
		j = c.Index[j]
		if users[i].Score > users[j].Score {
			return true
		} else if users[i].Score == users[j].Score {
			if users[i].Penalty < users[j].Penalty {
				return true
			} else if users[i].Penalty == users[j].Penalty {
				return strings.Compare(users[i].Name, users[j].Name) < 0
			}
		}
		return false
	}

	heapSort(c)

	for i := 0; i < n; i++ {
		fmt.Println(users[c.Index[i]].Name)
	}
}

type User struct {
	Name    string
	Score   int
	Penalty int
}

type SortableIndex struct {
	Index []int
	Less  func(i, j int) bool
}

func (c *SortableIndex) Swap(i, j int) {
	c.Index[i], c.Index[j] = c.Index[j], c.Index[i]
}

func iRange(n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = i
	}
	return r
}

func heapSort(sortable *SortableIndex) {
	heapify(sortable)
	popSort(sortable)
}

func popSort(sortable *SortableIndex) {
	d := sortable.Index
	end := len(d)

	for end > 1 {
		end--
		sortable.Swap(0, end)
		sDown(sortable, 0, end)
	}
}

func heapify(sortable *SortableIndex) {
	if len(sortable.Index) < 2 {
		return
	}

	for i := parent(len(sortable.Index) - 1); i >= 0; i-- {
		sDown(sortable, i, len(sortable.Index))
	}
}

func sDown(sortable *SortableIndex, idx, end int) int {
	left := leftChild(idx)
	right := rightChild(idx)

	if left >= end {
		return idx
	}

	maxChildIdx := left
	if right < end && sortable.Less(left, right) {
		maxChildIdx = right
	}

	if sortable.Less(idx, maxChildIdx) {
		sortable.Swap(idx, maxChildIdx)
		return sDown(sortable, maxChildIdx, end)
	}

	return idx
}

func parent(idx int) int {
	return (idx+1)/2 - 1
}

func leftChild(idx int) int {
	return idx*2 + 1
}

func rightChild(idx int) int {
	return idx*2 + 2
}
