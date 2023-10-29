package main

import (
	"fmt"
	"testing"
)

func Test_CanCalculateSubhashes(t *testing.T) {
	txt := "abcde"
	// firstI := 2
	// lastI := 3


	a := int64(1000)
	m := int64(1000009)
	fmt.Println(m)
	fmt.Println(hash("abc", a,m))
	fmt.Println(hash("abcd", a,m))
	subHashs := subhash(txt, a, m)
	fmt.Println(subHashs)

	fmt.Println(hash("bc", a,m))
	d :=(subHashs[0].hash * subHashs[2].powA)% m
	fmt.Println("d =",d)
	hash := ((subHashs[2].hash) - (d))%(m) 
	fmt.Println(hash)
	if hash < 0 {
		hash = m + hash
	}
	fmt.Println(hash)
}

