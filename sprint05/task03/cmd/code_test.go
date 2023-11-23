package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_(t *testing.T) {
	assert.True(t, true)

	t.Run("nil is palindrome", func(t *testing.T) {
		assert.True(t, isPalindrome(nil))
	})

	t.Run("single nod is palindrome", func(t *testing.T) {
		root := &Node{0, nil, nil}
		assert.True(t, isPalindrome(root))
	})

	t.Run("root with left child is not a palindrome", func(t *testing.T) {
		left := &Node{0, nil, nil}
		root := &Node{0, left, nil}
		assert.False(t, isPalindrome(root))
	})

	t.Run("root has both children equal", func(t *testing.T) {
		left := &Node{0, nil, nil}
		right := &Node{0, nil, nil}
		root := &Node{0, left, right}
		assert.True(t, isPalindrome(root))
	})

	t.Run("root has both children and they are not equal", func(t *testing.T) {
		left := &Node{1, nil, nil}
		right := &Node{0, nil, nil}
		root := &Node{0, left, right}
		assert.False(t, isPalindrome(root))
	})

	t.Run("root has one grand children", func(t *testing.T) {
		ll := &Node{1, nil, nil}
		left := &Node{0, ll, nil}
		right := &Node{0, nil, nil}
		root := &Node{0, left, right}
		assert.False(t, isPalindrome(root))
	})
}

func Test_CanGetValuesOnLevels(t *testing.T) {
	t.Run("can get level values for nil", func(t *testing.T) {
		var levelValues []*int
		iterateLevels(nil, func(vals []*int) bool {
			levelValues = vals
			return true
		})
		assert.Equal(t, ([]*int)(nil), levelValues)
	})

	t.Run("single node", func(t *testing.T) {
		var levelValues []*int
		root := &Node{1, nil, nil}
		iterateLevels(root, func(v []*int) bool {
			levelValues = v
			return true
		})
		assert.Equal(t, 1, *levelValues[0])
	})

	t.Run("node with children", func(t *testing.T) {
		left := &Node{1, nil, nil}
		right := &Node{2, nil, nil}
		root := &Node{0, left, right}

		levelValues := [][]*int{}
		iterateLevels(root, func(vals []*int) bool {
			levelValues = append(levelValues, vals)
			return true
		})
		assert.Equal(t, 0, *levelValues[0][0])
		assert.Equal(t, 1, *levelValues[1][0])
		assert.Equal(t, 2, *levelValues[1][1])
	})
}

func Test_MyQueu(t *testing.T) {
	t.Run("Can create Empty queue", func(t *testing.T) {
		q := &MyQueue{}
		assert.True(t, q.Empty())
	})

	t.Run("CanPushAndPop", func(t *testing.T) {
		q := &MyQueue{}
		q.PushBack(Node{1, nil, nil})
		assert.False(t, q.Empty())
		v := q.PopFront()
		val := v.(Node)
		assert.Equal(t, 1, val.value)
		assert.True(t, q.Empty())
	})

	t.Run("Can2PushAnd2Pop", func(t *testing.T) {
		q := &MyQueue{}
		q.PushBack(Node{1, nil, nil})
		q.PushBack(Node{2, nil, nil})
		{
			val := q.PopFront().(Node)
			assert.Equal(t, 1, val.value)
		}
		{
			val := q.PopFront().(Node)
			assert.Equal(t, 2, val.value)
		}
	})

	t.Run("Can2PushAndPop", func(t *testing.T) {
		q := &MyQueue{}
		q.PushBack(Node{1, nil, nil})
		{
			val := q.PopFront().(Node)
			assert.Equal(t, 1, val.value)
		}
		q.PushBack(Node{2, nil, nil})
		{
			val := q.PopFront().(Node)
			assert.Equal(t, 2, val.value)
		}
	})
}

func Test_iterateLevels2(t *testing.T) {
	t.Run("Can iterate levels on nil", func(t *testing.T) {
		q := &MyQueue{}
		q.PushBack(NodeWrap{nil, 0})

		var value *int
		var level int
		iterateLevels2(q, func(l int, v *int) bool {
			value = v
			level = l
			return true
		})
		assert.Equal(t, (*int)(nil), value)
		assert.Equal(t, 0, level)
	})

	t.Run("single node", func(t *testing.T) {
		root := &Node{1, nil, nil}
		q := &MyQueue{}
		q.PushBack(NodeWrap{root, 0})
		levelVal := [][]*int{}
		iterateLevels2(q, func(l int, v *int) bool {
			if len(levelVal) <= l {
				levelVal = append(levelVal, []*int{})
			}
			levelVal[l] = append(levelVal[l], v)
			return true
		})
		fmt.Println(levelVal)
	})
}
