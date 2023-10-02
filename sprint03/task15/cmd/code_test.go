package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanFoundKth0(t *testing.T) {
	inp := []int{9, 100}
	act := solution2(inp, 0)
	assert.Equal(t, 0, act)
}

func Test_CanFoundKth03(t *testing.T) {
	inp := []int{9, 100, 25}
	act := solution2(inp, 0)
	assert.Equal(t, 0, act)
}

func Test_CanFoundKth13(t *testing.T) {
	inp := []int{9, 100, 25}
	act := solution2(inp, 1)
	assert.Equal(t, 16, act)
}

func Test_CanFoundKth23(t *testing.T) {
	inp := []int{9, 100, 25}
	act := solution2(inp, 2)
	assert.Equal(t, 75, act)
}

func Test_CanFoundKth33(t *testing.T) {
	inp := []int{9, 100, 25}
	act := solution2(inp, 3)
	assert.Equal(t, 91, act)
}

func Test_Practicum1(t *testing.T) {
	inp := []int{2, 3, 4}
	var act int
	act = solution2(inp, 1)
	assert.Equal(t, 1, act)
	act = solution2(inp, 2)
	assert.Equal(t, 1, act)
	act = solution2(inp, 3)
	assert.Equal(t, 2, act)
}

func Test_Practicum2(t *testing.T) {
	inp := []int{9, 1000, 250}
	var act int
	act = solution2(inp, 0)
	assert.Equal(t, 0, act)
	act = solution2(inp, 1)
	assert.Equal(t, 241, act)
	act = solution2(inp, 2)
	assert.Equal(t, 750, act)
	act = solution2(inp, 3)
	assert.Equal(t, 991, act)
}
