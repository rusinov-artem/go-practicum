package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_N(t *testing.T) {
	lq := &ListQueue{}
	assert.Equal(t, 0, lq.size)
	lq.put(3)
	assert.Equal(t, 1, lq.size)
	v, err := lq.get()
	assert.Equal(t, 0, lq.size)
	assert.Nil(t, err) 
	assert.Equal(t, 3, v)
	v, err = lq.get()
	assert.Error(t, err)
	assert.Equal(t, 0, v)
	assert.Equal(t, 0, lq.size)
}

func Test_QueuCanHandleMultipleElemens(t *testing.T) {
	lq := &ListQueue{}
	lq.put(1)
	assert.Equal(t, 1, lq.size)
	lq.put(2)
	assert.Equal(t, 2, lq.size)
	lq.put(3)
	assert.Equal(t, 3, lq.size)
	{
		v, err := lq.get()
		assert.Equal(t, 2, lq.size)
		assert.Nil(t, err) 
		assert.Equal(t, 1, v)
	}
	{
		v, err := lq.get()
		assert.Equal(t, 1, lq.size)
		assert.Nil(t, err) 
		assert.Equal(t, 2, v)
	}
	{
		v, err := lq.get()
		assert.Equal(t, 0, lq.size)
		assert.Nil(t, err) 
		assert.Equal(t, 3, v)
	}
}

func Test_CanMixOperations(t *testing.T) {
	lq := &ListQueue{}
	lq.put(1)
	assert.Equal(t, 1, lq.size)
	lq.put(2)
	{
		v, err := lq.get()
		assert.Equal(t, 1, lq.size)
		assert.Nil(t, err) 
		assert.Equal(t, 1, v)
	}
	lq.put(3)
	{
		v, err := lq.get()
		assert.Equal(t, 1, lq.size)
		assert.Nil(t, err) 
		assert.Equal(t, 2, v)
	}
	{
		v, err := lq.get()
		assert.Equal(t, 0, lq.size)
		assert.Nil(t, err) 
		assert.Equal(t, 3, v)
	}
    
}
