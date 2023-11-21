package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Fib(t *testing.T) {
	// 1 1 2 3 5 8 13
	assert.Equal(t, 1, fib(1))
	assert.Equal(t, 1, fib(2))
	assert.Equal(t, 2, fib(3))
	assert.Equal(t, 3, fib(4))
	assert.Equal(t, 5, fib(5))
	assert.Equal(t, 8, fib(6))
	assert.Equal(t, 13, fib(7))
}

func Test_Factorial(t *testing.T) {
	assert.Equal(t, int64(1), factorial(1))
	assert.Equal(t, int64(2), factorial(2))
	assert.Equal(t, int64(6), factorial(3))
	assert.Equal(t, int64(24), factorial(4))
}

func Test_Katalan(t *testing.T) {
	assert.Equal(t, int64(1), katalan(0))
	assert.Equal(t, int64(1), katalan(1))
	assert.Equal(t, int64(2), katalan(2))
	assert.Equal(t, int64(5), katalan(3))
	assert.Equal(t, int64(14), katalan(4))
	assert.Equal(t, int64(9694845), katalan(15))
}
