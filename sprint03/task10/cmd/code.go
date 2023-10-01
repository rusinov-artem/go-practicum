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
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 0, 5*1024*1024), 0)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	data := make([]int, 0, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		data = append(data, v)
	}
	bubleSort(data)
}

func bubleSort(data []int) {
	printedCount := 0
	for j := len(data) - 1; j > 0; j-- {
		sorted := true
		for i := 0; i < j; i++ {
			if data[i] > data[i+1] {
				sorted = false
				data[i], data[i+1] = data[i+1], data[i]
			}
		}
		if sorted {
			break
		}
		printData(data)
		printedCount++
	}
	if printedCount == 0 {
		printData(data)
	}
}

func printData(data []int) {
	sb := strings.Builder{}
	if len(data) > 0 {
		sb.WriteString(fmt.Sprintf("%d", data[0]))
	}
	for i := 1; i < len(data); i++ {
		sb.WriteString(fmt.Sprintf(" %d", data[i]))
	}
	fmt.Printf("%s\n", sb.String())

}
