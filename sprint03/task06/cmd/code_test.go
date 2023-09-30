package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxP(t *testing.T) {
	assertMaxP(t, 0, []int{0, 2, 3})
	assertMaxP(t, 8, []int{3, 2, 3})
	assertMaxP(t, 8, []int{6,3,3,2})
	assertMaxP(t, 20, []int{5,3,7,2,8,3})
}

func assertMaxP(t *testing.T, exp int, list []int) {
	act := maxP(list)
	assert.Equal(t, exp, act)
}

func Test_CheckIsTriangle(t *testing.T) {
    assert.False(t, isTriangle(0,0,0))
	assert.True(t, isTriangle(3,2,3))
}
