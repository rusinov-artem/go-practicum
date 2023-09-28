package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var keyboard = []string{
	"", "",
	"abc", "def",
	"ghi", "jkl", "mno",
	"pqrs", "tuv", "wxyz",
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()
	inp := []int{}
	for _, c := range s {
		i, _ := strconv.Atoi(string(c))
		inp = append(inp, i)
	}

	res := generate(inp)

	if len(res) > 0 {
		fmt.Printf("%s", res[0])
	}

	for i := 1; i < len(res); i++ {
		fmt.Printf(" %s", res[i])
	}
}

func generate(inp []int) []string {
	if len(inp) == 1 {
		res := []string{}
		for _, v := range keyboard[inp[0]] {
			res = append(res, string(v))
		}
		return res
	}

	p := generate(inp[1:])
	res := []string{}
	for _, c := range keyboard[inp[0]] {
		for _, s := range p {
			res = append(res, string(c)+s)
		}
	}
	return res
}
