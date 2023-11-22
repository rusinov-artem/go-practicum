package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanCountMaxScore(t *testing.T) {
	assert.True(t, true)
    t.Run("empty tree", func(t *testing.T){
       root := (*Node)(nil)
       ms := maxScore(root)
       assert.Equal(t, 0, ms)
    })

    t.Run("single element", func(t *testing.T){
       root := &Node{5, nil, nil}
       ms := maxScore(root)
       assert.Equal(t, 5, ms)
    })

   t.Run("single element less than 0", func(t *testing.T){
      root := &Node{-1, nil, nil}
      ms := maxScore(root)
      assert.Equal(t, -1, ms)
   })

   t.Run("root with left child", func(t *testing.T){
      root := &Node{5, nil, nil}
      left := &Node{2, nil, nil}

      root.left = left

      ms := maxScore(root)
      assert.Equal(t, 7, ms)
   })

   t.Run("root with two children", func(t *testing.T){
      root := &Node{5, nil, nil}
      left := &Node{2, nil, nil}
      right := &Node{20, nil, nil}

      root.left = left
      root.right = right

      ms := maxScore(root)
      assert.Equal(t, 27, ms)
   })

   t.Run("root with left child < 0", func(t *testing.T){
      root := &Node{5, nil, nil}
      left := &Node{-2, nil, nil}
      right := &Node{20, nil, nil}

      root.left = left
      root.right = right

      ms := maxScore(root)
      assert.Equal(t, 25, ms)
   })

   t.Run("root with right child < 0", func(t *testing.T){
      root := &Node{5, nil, nil}
      left := &Node{2, nil, nil}
      right := &Node{-20, nil, nil}

      root.left = left
      root.right = right

      ms := maxScore(root)
      assert.Equal(t, 7, ms)
   })

   t.Run("root < 0 ", func(t *testing.T){
      root := &Node{-5, nil, nil}
      left := &Node{2, nil, nil}
      right := &Node{20, nil, nil}

      root.left = left
      root.right = right

      ms := maxScore(root)
      assert.Equal(t, 20, ms)
   })

   t.Run("complex test", func(t *testing.T){
      root := &Node{-5, nil, nil}
      rl := &Node{1, nil, nil}
      rr := &Node{7, nil, nil}
      rrl := &Node{2, nil, nil}
      rrr := &Node{3, nil, nil}

      root.left = rl
      root.right = rr
      root.right.left = rrl
      root.right.right = rrr

      ms := maxScore(root)
      assert.Equal(t, 12, ms)
   })

   t.Run("complex test 2", func(t *testing.T){
      r1 := &Node{-9, nil, nil}
      r3 := &Node{-15, nil, nil}
      r4 := &Node{-7, nil, nil}
      r2 := &Node{-20, r3, r4}
      r0 := &Node{-10, r1, r2}

      //          -10
      //      -9           -20 
      //             -15           -7 

      ms := maxScore(r0)
      assert.Equal(t, -7, ms)
   })

   t.Run("complex test 3", func(t *testing.T){
      r3 := &Node{1, nil, nil}
      r2 := &Node{1, nil, nil}
      r1 := &Node{16, r3, nil}
      r0 := &Node{1, r1, r2}
      //         1 
      //       16  1 
      //     1

      ms := maxScore(r0)
      assert.Equal(t, 19, ms)
   })
}
