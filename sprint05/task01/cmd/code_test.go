package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanFindNode(t *testing.T) {
	t.Run("single node", func(t *testing.T) {
		root := &Node{3, nil, nil}
		max := Solution(root)
		assert.Equal(t, 3, max)
	})

	t.Run("2 level nodes", func(t *testing.T){
		lChild := &Node{10, nil, nil}
		rChild := &Node{999, nil, nil}
		root := &Node{3, lChild, rChild}
		max := Solution(root)
		assert.Equal(t, 999, max)
	})
}
