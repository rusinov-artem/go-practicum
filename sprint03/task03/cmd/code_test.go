package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanCheckSubstr(t *testing.T) {
	assert.True(t, checkSubS("", "abc"))
	assert.True(t, checkSubS("a", "abc"))
	assert.True(t, checkSubS("b", "abc"))
	assert.True(t, checkSubS("c", "abc"))
	assert.True(t, checkSubS("ac", "abc"))
	assert.True(t, checkSubS("bc", "abc"))
	assert.True(t, checkSubS("abc", "abc"))
	
	assert.False(t, checkSubS("d", "abc"))
	assert.False(t, checkSubS("abd", "abc"))
	assert.False(t, checkSubS("acd", "abc"))

	assert.False(t, checkSubS("cab", "abc"))
	assert.False(t, checkSubS("aab", "abb"))
}

func checkSubS(s, t string) bool{
	return checkSub([]rune(s), []rune(t))
}


