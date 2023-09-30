package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanCreateGame(t *testing.T) {
	game := NewGame(
		[]int{10},
		mkg([]int{}),
		0,
	)
	
	assert.Equal(t, 0, game.happy())
}

func Test_CanFindSingleCookie(t *testing.T) {
	game := NewGame(
		[]int{10},
		mkg([]int{10}),
		1,
	)
	
	assert.Equal(t, 1, game.happy())
}

func Test_CanFindSingleCookieNext(t *testing.T) {
	game := NewGame(
		[]int{10},
		mkg([]int{100}),
		1,
	)
	
	assert.Equal(t, 1, game.happy())
}

func Test_CanFindSingleCookieNextEdge(t *testing.T) {
	game := NewGame(
		[]int{10},
		mkg([]int{1000}),
		1,
	)
	
	assert.Equal(t, 1, game.happy())
}

func Test_MoreAssertions(t *testing.T) {
	assertHappy(t, 2, 
		[]int{10, 25}, 
		[]int{1000, 1000},
	)
	assertHappy(t, 1, 
		[]int{10, 25}, 
		[]int{1000 },
	)
	assertHappy(t, 1, 
		[]int{10, 25}, 
		[]int{1000, 1 },
	)
	assertHappy(t, 0, 
		[]int{10, 25}, 
		[]int{1, 1 },
	)
	assertHappy(t, 0, 
		[]int{10, 25}, 
		[]int{ },
	)
	assertHappy(t, 2, 
		[]int{10, 25}, 
		[]int{1000, 1000, 1000 },
	)
	assertHappy(t, 0, 
		[]int{}, 
		[]int{1000, 1000, 1000 },
	)
}

func assertHappy(t *testing.T, exp int, g, inp []int) {
	game := NewGame(g, mkg(inp), len(inp))
	act := game.happy()
	assert.Equal(t, exp, act)
}

func mkg(inp []int) []int {
	res := make([]int, 1001)
	for _, v := range inp {
		res[v]++
	}
	return res
}
