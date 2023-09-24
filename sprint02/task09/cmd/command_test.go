package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetErrorOnPush(t *testing.T) {
	q := NewMyQueue(0)
    err := q.push(1)
	assert.Error(t, err)
}

func Test_CanGetErrorOnPop(t *testing.T) {
	q := NewMyQueue(0)
	_, err := q.pop()
	assert.Error(t, err)
}

func Test_CanPop(t *testing.T) {
	q := NewMyQueue(1)
	err := q.push(1)
	assert.Nil(t, err)
	val, _ := q.pop()
	assert.Equal(t, 1, val)
}

func Test_CanPushPushPopPop(t *testing.T) {
	q := NewMyQueue(2)
	q.push(1)
	q.push(2)
	var p int
	p, _ = q.pop()
	assert.Equal(t, 1, p)
	p, _ = q.pop()
	assert.Equal(t, 2, p)
	q.push(1)
	q.push(2)
	p, _ = q.pop()
	assert.Equal(t, 1, p)
	p, _ = q.pop()
	assert.Equal(t, 2, p)
}

func Test_CanGetQueueSize(t *testing.T) {
	q := NewMyQueue(2)
	var p int
	q.push(1)
	p, _ = q.peek()
	assert.Equal(t, 1, p)
	q.push(2)
	p, _ = q.peek()
	assert.Equal(t, 1, p)
	_, _ = q.pop()
	p, _ = q.peek()
	assert.Equal(t, 2, p)
	
}
