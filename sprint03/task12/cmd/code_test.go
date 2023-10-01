package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LowerBound(t *testing.T) {
	testCases := []struct{
		Name string
		Expected int
		Key int
		Data []int
	}{
		{
			"Return -1 if empty",
			-1, 3,
			[]int{},
		},
		{
			"Can find first one",
			0, 3,
			[]int{3},
		},
		{
			"Can find first one",
			0, 3,
			[]int{30},
		},
		{
			"Can find second one",
			1, 3,
			[]int{-3, 3},
		},
		{
			"Can find second one",
			1, 3,
			[]int{-3, 30},
		},
		{
			"Can find third one",
			2, 3,
			[]int{-3, 0, 30},
		},
		{
			"Can find third one",
			2, 3,
			[]int{-3, 0, 3},
		},
		{
			"Can find middle one",
			1, 3,
			[]int{-3, 3, 30},
		},
		{
			"Can find middle one",
			1, 3,
			[]int{-3, 31, 39},
		},
		{
			"Can find",
			2, 3,
			[]int{-3,0,5, 31, 39},
		},
		{
			"Can find",
			2, 3,
			[]int{-3,2,3, 31, 39},
		},
		{
			"Can find",
			2, 3,
			[]int{1,2,4,4,6,8},
		},
		{
			"Can find on 3rd place",
			2,4,
			[]int{1, 1, 4, 4, 4, 4},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			act := lowerBound(testCase.Data, testCase.Key)
			assert.Equal(t, testCase.Expected, act)
		})
	}
}

