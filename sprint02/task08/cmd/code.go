package main

import (
	"bufio"
	"fmt"
	"os"
)

// {(})
// ({[]})
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 10000), 0)
	sc.Scan()
	line := sc.Text()
	ok := CheckBraces(line)

	if ok {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func CheckBraces(s string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s) % 2 != 0 {
		return false
	}

	st := make([]rune, 0, len(s))
	for _, r := range s {
		if isCloseBrace(r) && len(st) == 0 {
			return false
		}

		if len(st) == 0 {
			st = append(st, r)
			continue
		}

		if isCloseBrace(r) {
			if isMatchBraces(st[len(st)-1], r) {
				st = st[:len(st)-1]
				continue
			}
		}

		st = append(st, r)

	}

	return len(st) == 0 
}

func isMatchBraces(left, right rune) bool {
	if left == '{' && right == '}' {
		return true
	}

	if left == '(' && right == ')' {
		return true
	}

	if left == '[' && right == ']' {
		return true
	}

	return false
}

func isCloseBrace(r rune) bool {
	return r == '}' || r == ')' || r == ']'
}
