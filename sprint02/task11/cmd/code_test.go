package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_N(t *testing.T) {
	assert.Equal(t, uint64(1), mfib(0))
	assert.Equal(t, uint64(1), mfib(1))
	assert.Equal(t, uint64(2), mfib(2))
	assert.Equal(t, uint64(3), mfib(3))
	assert.Equal(t, uint64(5), mfib(4))
}
