package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanCreateGenerator(t *testing.T) {
	g := NewGenerator(1, []uint64{1, 2, 3})
	var act []uint64
	act = g.Next()
	assert.Equal(t, []uint64{1}, act)
	act = g.Next()
	assert.Equal(t, []uint64{2}, act)
	act = g.Next()
	assert.Equal(t, []uint64{3}, act)
	act = g.Next()
	assert.Equal(t, []uint64{1}, act)
}

func Test_CanGenerate2(t *testing.T) {
	g := NewGenerator(2, []uint64{1, 2, 3})
	var act []uint64
	act = g.Next()
	assert.Equal(t, []uint64{1, 1}, act)
	act = g.Next()
	assert.Equal(t, []uint64{1, 2}, act)
	act = g.Next()
	assert.Equal(t, []uint64{1, 3}, act)
	act = g.Next()
	assert.Equal(t, []uint64{2, 1}, act)
}

func Test_CanGoRound(t *testing.T) {
	g := NewGenerator(2, []uint64{1, 2, 3})
	for i := 0; i < 2*2*2+2; i++ {
		v := g.Next()
		fmt.Println(v)
	}
}
