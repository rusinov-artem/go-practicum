package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 4*1024*1024), 0)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	g := NewAnogramGroup()
	for i := 0; i < n; i++ {
		sc.Scan()
		g.Add([]rune(sc.Text()), i)
	}
	g.Print()
}

type AnogramGroups struct {
	Maps  map[string][]int
	Index []string
}

func NewAnogramGroup() *AnogramGroups {
	return &AnogramGroups{
		Maps: make(map[string][]int),
	}
}

func (this *AnogramGroups) Add(txt []rune, i int) {
	sort.Slice(txt, func(l, r int) bool {
		return txt[l] < txt[r]
	})

	sTxt := string(txt)
	if v, ok := this.Maps[sTxt]; ok {
		this.Maps[sTxt] = append(v, i)
	} else {
		this.Maps[sTxt] = []int{i}
		this.Index = append(this.Index, sTxt)
	}
}

func (this *AnogramGroups) Print() {
	sb := strings.Builder{}
	for _, s := range this.Index {
		indexies := this.Maps[s]
		sb.WriteString(fmt.Sprintf("%d", indexies[0]))
		for i:=1 ; i<len(indexies); i++ {
			sb.WriteString(fmt.Sprintf(" %d", indexies[i]))
		}
		sb.WriteString(fmt.Sprintf("\n"))
	}
	fmt.Println(sb.String())
}
