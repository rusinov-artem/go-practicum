package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 10000), 0)
	sc.Scan()
	s := []rune(sc.Text())
	sc.Scan()
	t := []rune(sc.Text())

	r := checkSub(s, t)
	if r {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

}

func checkSub(s, t []rune) bool {
	if len(s) > len(t) {
		return false
	}

	j := 0
	for i := 0; i < len(s) ; i++ {
		for j < len(t) && t[j] != s[i] {
			j++
		}
		if j == len(t) {
			return false
		}
		j++
	}

	return true
}
