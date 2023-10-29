package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestSubstr(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		txt := ""
		res := longestSubstr([]rune(txt))
		assert.Equal(t, 0, res)
	})
	t.Run("string with 1 letter", func(t *testing.T) {
		txt := "a"
		res := longestSubstr([]rune(txt))
		assert.Equal(t, 1, res)
	})
	t.Run("string with 2 different letters", func(t *testing.T) {
		txt := "ab"
		res := longestSubstr([]rune(txt))
		assert.Equal(t, 2, res)
	})
	t.Run("string with 2 different letters", func(t *testing.T) {
		txt := "aa"
		res := longestSubstr([]rune(txt))
		assert.Equal(t, 1, res)
	})
	t.Run("string with 2 different letters", func(t *testing.T) {
		txt := "aaabcddd88990012345"
		res := longestSubstr([]rune(txt))
		assert.Equal(t, 6, res)
	})
	t.Run("corner case1", func(t *testing.T) {
		txt := "abcadefj"
		res := longestSubstr([]rune(txt))
		assert.Equal(t, 7, res)
	})
}
