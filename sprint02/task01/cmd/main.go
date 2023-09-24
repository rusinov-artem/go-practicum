package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var nI int
	var nJ int
	fmt.Scanf("%d\n%d", &nI, &nJ)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte,0, 10000), 0)

	matrix := make([][]int, nJ)
	for j := 0; j < nJ; j++ {
		matrix[j] = make([]int, nI)
	} 

	for i := 0; i < nI; i++ {
		for j := 0; j < nJ; j++ {
			sc.Scan()
			val, _ := strconv.Atoi(sc.Text())
			matrix[j][i] = val
		}
	}

	builder := strings.Builder{}
	for j:=0; j < nJ; j++ {
		builder.WriteString(strconv.Itoa(matrix[j][0]))
		for i:=1; i < nI; i++ {
			builder.WriteString(" " + strconv.Itoa(matrix[j][i]))
		}
		builder.WriteString("\n")
	}
	fmt.Println(builder.String())
}
