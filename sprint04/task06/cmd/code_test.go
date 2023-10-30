package main

import (
	"fmt"
	"testing"
)

func Test_AnogramGroups(t *testing.T) {
	g := NewAnogramGroup()
	g.Add([]rune("abs"), 1)
	fmt.Println(g.Maps)
	g.Print()
}
