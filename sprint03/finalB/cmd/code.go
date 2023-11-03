package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Быстрая сортировка сделана без выделелния дополнительной
// памяти при партицировании данный
// Алгоритмическая сложность O(n*log(n)) в среднем и
// O(n^2) - в худшем случае, если данные уже отсортированны
//
// Дополнительная память выделяется на рекурсивных вызовах
// в среднем O(log(n)) на рекурсинвые вызовы, и в худшем случае O(n)
// на рекурсивные вызовы
// + необходимо хранить все сортируемые объекты в памяти
// Cоответственно оценка по памяти:
// O(log(n) + n) -> O(n) в среднем
// O(n + n) -> O(2n) в худшем случае

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 4*1024*1024), 0)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	users := make([]User, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		users[i].Name = sc.Text()
		sc.Scan()
		users[i].Score, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		users[i].Penalty, _ = strconv.Atoi(sc.Text())
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

	qSort(c)

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

func qSort(c *SortableIndex) {
	qSortImpl(c, 0, len(c.Index))
}

func qSortImpl(c *SortableIndex, begin, end int) {
	if end-begin < 2 {
		return
	}
	p := partition(c, begin, end)
	qSortImpl(c, begin, p)
	qSortImpl(c, p+1, end)
}

func iRange(n int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	return r.Perm(n)
}

func partition(c *SortableIndex, begin, end int) int {
	if end-begin < 2 {
		return begin
	}

	left := begin
	right := end - 2
	mid := end - 1

	for left <= right {
		for !c.Less(mid, left) && left < mid {
			left++
		}

		for c.Less(mid, right) && right > 0 {
			right--
		}

		if left < right {
			c.Swap(left, right)
		}
	}

	c.Swap(mid, left)

	return left
}

func part(inp []int, begin, end int) int {
	if end-begin < 2 {
		return begin
	}

	left := begin
	right := end - 2
	mid := end - 1

	for {
		for inp[left] <= inp[mid] && left < mid {
			left++
		}

		for inp[right] > inp[mid] && right > 0 {
			right--
		}

		if left < right {
			swap(inp, right, left)
		} else {
			break
		}
	}

	swap(inp, mid, left)

	return left
}

func swap(inp []int, i, j int) {
	inp[i], inp[j] = inp[j], inp[i]
}

func qs(inp []int, begin, end int) {
	if end-begin < 2 {
		return
	}
	p := part(inp, begin, end)
	qs(inp, begin, p)
	qs(inp, p+1, end)
}
