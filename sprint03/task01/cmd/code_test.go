package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanGenerate(t *testing.T) {
	assertBraces2(t, 0, []string{})
	assertBraces2(t, 1, []string{"()"})
	assertBraces2(t, 2, []string{"(())", "()()"})
	assertBraces2(t, 3, []string{
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()",
	})
}


func Test_CanCheckBraces(t *testing.T) {
	assert.Equal(t, true, checkBraces(""))
	assert.Equal(t, true, checkBraces("()"))
	assert.Equal(t, true, checkBraces("(()())"))
	assert.Equal(t, true, checkBraces("((()()))"))
	assert.Equal(t, false, checkBraces(")()"))
	assert.Equal(t, false, checkBraces("())()"))
}

func assertBraces2(t *testing.T, n int, exp []string) {
	res := generateBraces2(n)
	assert.Equal(t, exp, res)
}

func TestGenerateB(t *testing.T) {
	assertBraces(t, 0, []string{})
	assertBraces(t, 1, []string{
		"(", ")",
	})
	assertBraces(t, 2, []string{
		"((", "()", ")(", "))",
	})
}

func assertBraces(t *testing.T, n int, exp []string) {
	res := generateBraces(n)
	assert.Equal(t, exp, res)
}
