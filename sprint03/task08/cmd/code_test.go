package main

import (
	"fmt"
	"sort"
	"testing"
)

func Test_CanSortStrings(t *testing.T) {
	d := []string{"6", "68", "69", "66", "99", "98", "9"}

	sort.Slice(d, func(l, r int) bool {
		return d[l] + d[r] > d[r] + d[l] 
	})

	fmt.Println(d)
}
