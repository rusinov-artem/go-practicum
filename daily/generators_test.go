package daily

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanGenerateWords(t *testing.T) {
	t.Run("can generate empty words list", func(t *testing.T) {
		gen := NewWordsGenerator([]int{})
		gen.Next()
		res := gen.Get()
		assert.Empty(t, res)
	})

	t.Run("can generate words from single letter", func(t *testing.T) {
		gen := NewWordsGenerator([]int{1})
		gen.Next()
		var res []int
		res = gen.Get()
		assert.Equal(t, []int{1}, res)
		gen.Next()
		res = gen.Get()
		assert.Equal(t, []int{1}, res)
	})

	t.Run("can generate words from two letters", func(t *testing.T) {
		gen := NewWordsGenerator([]int{1, 2})
		var res []int
		res = gen.Get()
		assert.Equal(t, []int{1, 1}, res)
		gen.Next()
		res = gen.Get()
		assert.Equal(t, []int{1, 2}, res)
		gen.Next()
		res = gen.Get()
		assert.Equal(t, []int{2, 1}, res)
		gen.Next()
		res = gen.Get()
		assert.Equal(t, []int{2, 2}, res)
		gen.Next()
		res = gen.Get()
		assert.Equal(t, []int{1, 1}, res)
	})
}

func Test_CanGeneratePermutations(t *testing.T) {
	t.Run("can generate permutation from empy list", func(t *testing.T) {
		gen := NewPermGenerator([]int{})
		var res []int
		gen.Next()
		res = gen.Get()
		assert.Equal(t, []int{}, res)
	})
	t.Run("can generate permutation from single", func(t *testing.T) {
		gen := NewPermGenerator([]int{1})
		var res []int
		gen.Next()
		res = gen.Get()
		assert.Equal(t, []int{1}, res)
	})
	t.Run("can generate permutation from two elements", func(t *testing.T) {
		gen := NewPermGenerator([]int{1, 2})
		var res []int
		res = gen.Get()
		assert.Equal(t, []int{1, 2}, res)
		gen.Next()
		res = gen.Get()
		assert.Equal(t, []int{2, 1}, res)
	})
	t.Run("can generate permutation from three elements", func(t *testing.T) {
		gen := NewPermGenerator([]int{1, 2, 3})
		var res []int
		res = gen.Get()
		assert.Equal(t, []int{1, 2, 3}, res)
		gen.Next()
		res = gen.Get()
		assert.Equal(t, []int{1, 3, 2}, res)
		gen.Next()
		res = gen.Get()
		assert.Equal(t, []int{2, 1, 3}, res)
		gen.Next()
		res = gen.Get()
		assert.Equal(t, []int{2, 3, 1}, res)
		gen.Next()
		res = gen.Get()
		assert.Equal(t, []int{3, 1, 2}, res)
		gen.Next()
		res = gen.Get()
		assert.Equal(t, []int{3, 2, 1}, res)
		gen.Next()
		res = gen.Get()
		assert.Equal(t, []int{1, 2, 3}, res)
	})

}

func Test_prem(t *testing.T) {
	res := perm([]int{1,2, 3})
	fmt.Println(res)
}

func Test_TestLogPerm(t *testing.T) {
	gen := NewPermGenerator([]int{1,2,3,4,5,6,7})
	for gen.Next() {
		fmt.Println(gen.Get())
	}
}
