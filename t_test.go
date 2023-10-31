package main

import (
	"fmt"
	"testing"
)

func Test_Map(t *testing.T) {
	ar := [10]int{}

	ar2 := ar

	fmt.Println(ar2)
	fmt.Println(ar)

	fmt.Println(ar2 == ar)

}
