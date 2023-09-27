package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_N(t *testing.T) {
	assert.Equal(t, uint64(1), mfib(0, 8))
	assert.Equal(t, uint64(1), mfib(1, 8))
	assert.Equal(t, uint64(2), mfib(2, 9))
	assert.Equal(t, uint64(3), mfib(3, 8))
	assert.Equal(t, uint64(5), mfib(4, 8))
}

func Test_pow10(t *testing.T) {

	assert.Equal(t, uint64(10), pow10(uint64(1)))
	assert.Equal(t, uint64(100), pow10(uint64(2)))
	assert.Equal(t, uint64(1000), pow10(uint64(3)))
    
}
