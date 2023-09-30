package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanCount(t *testing.T) {
	assertCount(t, 0, 0, []int{})
	assertCount(t, 1, 1, []int{10, 20, 1})
	assertCount(t, 2, 21, []int{10, 20, 1})
	assertCount(t, 2, 11, []int{10, 20, 1})
	assertCount(t, 3, 99999, []int{10, 20, 1})
}

func assertCount(t *testing.T, exp, balance int, buildings []int ) {
	g := game{balance, buildings}
	count := g.count()
	assert.Equal(t, exp, count)
}
