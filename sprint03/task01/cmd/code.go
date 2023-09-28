package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	res := generateBraces2(n)
	for _, v := range res {
		fmt.Println(v)
	}
}

func generateBraces2(n int) []string {
	all := generateBraces(2 * n)
	res := []string{}
	for _, v := range all {
		if checkBraces(v) {
			res = append(res, v)
		}
	}
	return res
}

func generateBraces(n int) []string {
	if n == 0 {
		return []string{}	
	}

	if n == 1 {
		return []string{"(", ")"}
	}

	res := []string{}

	bp := generateBraces(n-1)
	for _, v := range bp {
		res = append(res, "(" + v)
	}
	for _, v := range bp {
		res = append(res, ")" + v)
	}
	
	return res
}

func checkBraces(str string) bool {
	sum := 0
	if len(str) % 2 != 0 {
		return false
	}

	for _, c := range str {
		if c == '(' {
			sum++
		} else if c == ')' {
			sum--
		}
		if sum < 0 {
			return false
		}
	}

	return sum == 0

}
