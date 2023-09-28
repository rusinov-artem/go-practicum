package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanGenerate(t *testing.T) {
	assertGenerate(t, []int{2}, []string{
		"a", "b", "c",
	})
	assertGenerate(t, []int{2,2}, []string{
		"aa", "ab", "ac",
		"ba", "bb", "bc",
		"ca", "cb", "cc",
	})
	assertGenerate(t, []int{9,2}, []string{
		"wa", "wb", "wc",
		"xa", "xb", "xc",
		"ya", "yb", "yc",
		"za", "zb", "zc",
	})
}

func assertGenerate(t *testing.T, inp []int, out []string) {
	act := generate(inp)
	assert.Equal(t, out, act)

}

///////////////////////////////////
// a 
