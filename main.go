package main

import (
	"fmt"

	"github.com/rusinov-artem/go-practicum/daily"
)

func main() {
	gen := daily.NewPermGenerator([]int{1, 2, 3, 4, 5, 6, 7,8,9,10,11, 12})
	for gen.Next() {
		fmt.Println(gen.Get())
	}
}
