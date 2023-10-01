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
	sc.Buffer(make([]byte, 0, 6*1024*1024), 0)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	data := make([]segment, 0, n)

	for i := 0; i < n; i++ {
		sc.Scan()
		op, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		cl, _ := strconv.Atoi(sc.Text())
		data = append(data, segment{op, cl})
	}

	sort.Slice(data, func(l, r int) bool {
		if data[l].open == data[r].open {
			return data[l].close > data[r].close
		}
		return data[l].open < data[r].open
	})

	res := []segment{}
	current := (*segment)(nil)
	for i := 0; i < n; i++ {
		if current == nil {
			current = &segment{
				data[i].open, data[i].close,
			}
			continue
		}

		if data[i].open <= current.close {
			if data[i].close > current.close {
				current.close = data[i].close
			}
			continue
		}

		if data[i].open > current.close {
			res = append(res, segment{
				current.open, current.close,
			})
			current = &segment{
				data[i].open, data[i].close,
			}
		}
	}
	if current != nil {
		res = append(res, segment{
			current.open, current.close,
		})
	}

	for _, v := range res {
		fmt.Printf("%d %d\n", v.open, v.close)
	}
}

type segment struct {
	open  int
	close int
}
