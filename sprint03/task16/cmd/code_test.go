package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanCalculateBloxForEmptySlice(t *testing.T) {
	act := maxBlox([]int{})
	assert.Equal(t, 0, act)
}

func Test_CanCalculateBloxForSlizeSize1(t *testing.T) {
	act := maxBlox([]int{1})
	assert.Equal(t, 1, act)
}

func Test_CanCalculateBloxForSlizeSize2(t *testing.T) {
    t.Run("0,1", func(t *testing.T) {
		act := maxBlox([]int{0,1})
		assert.Equal(t, 2, act)
	})

	t.Run("1,0", func(t *testing.T) {
		act := maxBlox([]int{1,0})
		assert.Equal(t, 1, act)
	})
}

func Test_CanCalculateBloxForSlizeSize3(t *testing.T) {
	t.Run("0,1,2", func(t *testing.T){
		act := maxBlox([]int{0,1,2})
		assert.Equal(t, 3, act)
	})
	t.Run("0,2,1", func(t *testing.T){
		act := maxBlox([]int{0,2,1})
		assert.Equal(t, 2, act)
	})
	t.Run("1,0,2", func(t *testing.T){
		act := maxBlox([]int{1,0,2})
		assert.Equal(t, 2, act)
	})
	t.Run("1,2,0", func(t *testing.T){
		act := maxBlox([]int{1,2,0})
		assert.Equal(t, 1, act)
	})
	t.Run("2,0,1", func(t *testing.T){
		act := maxBlox([]int{2,0,1})
		assert.Equal(t, 1, act)
	})
	t.Run("2,1,0", func(t *testing.T){
		act := maxBlox([]int{2,1,0})
		assert.Equal(t, 1, act)
	})
}
