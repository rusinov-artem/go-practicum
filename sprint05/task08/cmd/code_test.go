package main

 import (
     "testing"

    "github.com/stretchr/testify/assert"
 )

 func Test_(t *testing.T) {
    assert.True(t, true) 

    t.Run("if nil than 0", func(t *testing.T){
       r := sumBranch(nil)
       assert.Equal(t, 0, r)
    })

   t.Run("single node", func(t *testing.T){
      root := &Node{1, nil, nil}
      r := sumBranch(root)
      assert.Equal(t, 1, r)
   })

   t.Run("single node", func(t *testing.T){
      root := &Node{1, nil, nil}
      r := sumBranch(root)
      assert.Equal(t, 1, r)
   })

   t.Run("single node with left child", func(t *testing.T){
      left := &Node{3, nil, nil}
      root := &Node{1, left, nil}
      r := sumBranch(root)
      assert.Equal(t, 13, r)
   })

   t.Run("complex 1", func(t *testing.T){
      left := &Node{3, nil, nil}
      right := &Node{5, nil, nil}
      root := &Node{1, left, right}
      r := sumBranch(root)
      assert.Equal(t, 28, r)
   })

   t.Run("complex 2", func(t *testing.T){
      rl := &Node{2, nil, nil}
      rr := &Node{1, nil, nil}
      left := &Node{2, nil, nil}
      right := &Node{3, rl, rr}
      root := &Node{1, left, right}
      r := sumBranch(root)
      assert.Equal(t, 275, r)
   })
 }

 func Test_CanIterateLeafes(t *testing.T) {
    t.Run("nil", func(t *testing.T){
       iterateLeafes(nil,0, nil) 
    }) 

   t.Run("single node", func(t *testing.T){
      root := &Node{1, nil, nil}
      leafs := []int{}
      iterateLeafes(root,0, func(val int) {
         leafs = append(leafs, val)
      }) 
      assert.Equal(t, []int{1}, leafs)
   }) 

   t.Run("single node with left child", func(t *testing.T){
      left := &Node{2, nil, nil}
      root := &Node{1, left, nil}
      leafs := []int{}
      iterateLeafes(root,0, func(val int) {
         leafs = append(leafs, val)
      }) 
      assert.Equal(t, []int{12}, leafs)
   }) 

   t.Run("single node with left and right children", func(t *testing.T){
      left := &Node{2, nil, nil}
      right := &Node{3, nil, nil}
      root := &Node{1, left, right}
      leafs := []int{}
      iterateLeafes(root,0, func(val int) {
         leafs = append(leafs, val)
      }) 
      assert.Equal(t, []int{12, 13}, leafs)
   }) 

   t.Run("single node with left left and right children", func(t *testing.T){
      leftLeft := &Node{4, nil, nil}
      left := &Node{2, leftLeft, nil}
      right := &Node{3, nil, nil}
      root := &Node{1, left, right}
      leafs := []int{}
      iterateLeafes(root,0, func(val int) {
         leafs = append(leafs, val)
      }) 
      assert.Equal(t, []int{124, 13}, leafs)
   }) 
 }
