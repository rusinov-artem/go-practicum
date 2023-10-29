package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 3*1024*1024), 0)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	txt := sc.Text()

	fmt.Println(longestSubstr([]rune(txt)))
}

func longestSubstr(txt []rune) int {
	if len(txt) < 2 {
		return len(txt)
	}

	table := make(map[rune]int)

	maxLen := 1
	curLen := 0
	for i := 0; i < len(txt); i++ {
		if oldI, ok := table[txt[i]]; !ok {
			table[txt[i]] = i
			curLen++
		} else {
			if maxLen < curLen {
				maxLen = curLen
			}
			curLen = 1
			i = oldI+1
			table = make(map[rune]int)
			table[txt[i]] = i
		}
	}

	if maxLen < curLen {
		maxLen = curLen
	}

	return maxLen
}
