package daily

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Partition(t *testing.T) {
	t.Run("Empty arr", func(t *testing.T) {
		inp := []int{}
		p := partition(inp, 0, len(inp))
		assert.Equal(t, 0, p)
	})

	t.Run("single elemnt", func(t *testing.T) {
		inp := []int{1}
		p := partition(inp, 0, len(inp))
		assert.Equal(t, 0, p)
	})

	t.Run("[1,2]", func(t *testing.T) {
		inp := []int{1, 2}
		p := partition(inp, 0, len(inp))
		assert.Equal(t, 1, p)
	})

	t.Run("[2,1]", func(t *testing.T) {
		inp := []int{2, 1}
		p := partition(inp, 0, len(inp))
		assert.Equal(t, 0, p)
		assert.Equal(t, []int{1, 2}, inp)
	})

	t.Run("[2,2]", func(t *testing.T) {
		inp := []int{2, 2}
		p := partition(inp, 0, len(inp))
		assert.Equal(t, 1, p)
		assert.Equal(t, []int{2, 2}, inp)
	})

	t.Run("[3,5,2,7,4]", func(t *testing.T) {
		inp := []int{3, 5, 2, 7, 4}
		p := partition(inp, 0, len(inp))
		assert.Equal(t, 2, p)
		assert.Equal(t, []int{3, 2, 4, 7, 5}, inp)
	})

	t.Run("[3,5,2,7,1]", func(t *testing.T) {
		inp := []int{3, 5, 2, 7, 1}
		p := partition(inp, 0, len(inp))
		assert.Equal(t, 0, p)
		assert.Equal(t, []int{1, 5, 2, 7, 3}, inp)
	})

	t.Run("[1,2,1,1]", func(t *testing.T) {
		inp := []int{1, 2, 1, 1}
		p := partition(inp, 0, len(inp))
		assert.Equal(t, 2, p)
		assert.Equal(t, []int{1, 1, 1, 2}, inp)
	})
}

func Test_CanSort(t *testing.T) {
	inp := []int{5, 4, 3, 10, 1, 2, 10}
	qSort(inp)
	assert.True(t, isSorted(inp))
}

func Test_Ultimate(t *testing.T) {
	g := NewWordsGenerator([]int{1, 2, 3, 4, 5})
	for g.Next() {
		inp := g.Get()
		t.Run(fmt.Sprintf("%v", inp), func(t *testing.T) {
			qSort(inp)
			assert.True(t, isSorted(inp))
		})
	}
}

func isSorted(d []int) bool {
	for i := 1; i < len(d); i++ {
		if d[i-1] > d[i] {
			return false
		}
	}
	return true
}
