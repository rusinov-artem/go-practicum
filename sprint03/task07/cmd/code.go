package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 10000), 10000000)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	r0 := 0
	r1 := 0
	r2 := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		if v == 0 {
			r0++
		} else if v == 1 {
			r1++
		} else if v == 2 {
			r2++
		}
	}

	sb := strings.Builder{}
	if r0 > 0 {
		sb.WriteString("0")
	}
	for i := 1; i < r0; i++ {
		sb.WriteString(" 0")
	}

	if r0 > 0 && r1 > 0 {
		sb.WriteString(" 1")
	}
	if r0 == 0 && r1 > 0 {
		sb.WriteString("1")
	}
	for i := 1; i < r1; i++ {
		sb.WriteString(" 1")
	}

	if (r0 > 0 || r1 > 0) && r2 > 0 {
		sb.WriteString(" 2")
	}
	if (r0 == 0 && r1 == 0) && r2 > 0 {
		sb.WriteString("2")
	}
	for i := 1; i < r2; i++ {
		sb.WriteString(" 2")
	}

	fmt.Println(sb.String())
}
