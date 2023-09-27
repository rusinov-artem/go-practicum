package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanCreateDeq(t *testing.T) {
	d := NewDeq(0)
	assert.Equal(t, 0, d.size)
	assert.Equal(t, 0, d.capacity)
	assert.NotNil(t, d)
}

func Test_CanPushBackValue(t *testing.T) {
	d := NewDeq(0)
	err := d.PushBack(10)
	assert.Error(t, err)
}

func Test_PushBackIncreaseSize(t *testing.T) {
	d := NewDeq(1)
	assert.Equal(t, 1, d.capacity)
	assert.Equal(t, 0, d.size)
	err := d.PushBack(10)
	assert.Equal(t, 1, d.size)
	assert.Nil(t, err)
}

func Test_CanPushFront(t *testing.T) {
	d := NewDeq(0)
	err := d.PushFront(10)
	assert.Error(t, err)
}

func Test_CanPopBack(t *testing.T) {
	d := NewDeq(0)
	_, err := d.PopBack()
	assert.Error(t, err)
}

func Test_CanPopBackAfterPushBack(t *testing.T) {
	d := NewDeq(1)
	d.PushBack(10)
	v, err := d.PopBack()
	assert.Equal(t, 10, v)
	assert.Nil(t, err)
}

func Test_CanPopBackAfter2PushBack(t *testing.T) {
	d := NewDeq(2)
	d.PushBack(10)
	d.PushBack(11)
	v, err := d.PopBack()
	assert.Equal(t, 11, v)
	assert.Nil(t, err)
}

func Test_Can2PopBackAfter2PushBack(t *testing.T) {
	d := NewDeq(2)
	d.PushBack(10)
	d.PushBack(11)
	_, _ = d.PopBack()
	v, err := d.PopBack()
	assert.Equal(t, 10, v)
	assert.Nil(t, err)
}

func Test_CanPopBackAfterPushFront(t *testing.T) {
	d := NewDeq(1)
	d.PushFront(10)
	v, err := d.PopBack()
	assert.Equal(t, 10, v)
	assert.Nil(t, err)
    
}

func Test_CanPopBackAfter2PushFront(t *testing.T) {
	d := NewDeq(2)
	d.PushFront(10)
	d.PushFront(11)
	_, _ = d.PopBack()
	v, err := d.PopBack()
	assert.Equal(t, 11, v)
	assert.Nil(t, err)
}

func Test_CanPopFront(t *testing.T) {
	d := NewDeq(0)
	_, err := d.PopFront()
	assert.Error(t, err)
}

func Test_CanPopFrontAfterPushBack(t *testing.T) {
	d := NewDeq(1)
	d.PushBack(10)
	v, err := d.PopFront()
	assert.Nil(t, err)
	assert.Equal(t, 10, v)
}

func Test_CanPopFrontAfterPushFront(t *testing.T) {
	d := NewDeq(1)
	d.PushFront(10)
	v, err := d.PopFront()
	assert.Nil(t, err)
	assert.Equal(t, 10, v)
}

func Test_Can2PopFrontAfter2PushBack(t *testing.T) {
	d := NewDeq(2)
	d.PushBack(10)
	d.PushBack(11)
	d.PopFront()
	v, err := d.PopFront()
	assert.Nil(t, err)
	assert.Equal(t, 11, v)
}

func Test_Can2PopFrontAfter2PushFront(t *testing.T) {
	d := NewDeq(2)
	d.PushFront(10)
	d.PushFront(11)
	d.PopFront()
	v, err := d.PopFront()
	assert.Nil(t, err)
	assert.Equal(t, 10, v)
}

func Test_CirclaWithPushBackPopFront(t *testing.T) {
	d := NewDeq(2)
	d.PushBack(10)
	d.PopFront()
	d.PushBack(11)
	d.PopFront()
	d.PushBack(12)
	v, err := d.PopFront()
	assert.Equal(t, 12, v)
	assert.Nil(t, err)
	
}

func Test_CirclaWithPushFrontPopBack(t *testing.T) {
	var v int
	var err error
	d := NewDeq(2)
	d.PushFront(10)
	v, err = d.PopBack()
	assert.Equal(t, 10, v)
	d.PushFront(11)
	v, err = d.PopBack()
	assert.Equal(t, 11, v)
	d.PushFront(12)
	v, err = d.PopBack()
	assert.Equal(t, 12, v)
	assert.Nil(t, err)
}
func Test_CirclaWithPushFrontPopBack_2(t *testing.T) {
	var v int
	var err error
	d := NewDeq(2)
	d.PushFront(10)
	d.PushFront(11)
	v, err = d.PopBack()
	assert.Equal(t, 10, v)
	d.PushFront(12)
	v, err = d.PopBack()
	assert.Equal(t, 11, v)
	d.PushFront(13)
	v, err = d.PopBack()
	assert.Equal(t, 12, v)
	assert.Nil(t, err)
}

func Test_PushBackPushFrontPopBackPopFront(t *testing.T) {
	var v int
	var err error
	d := NewDeq(2)
	d.PushBack(11)
	d.PushFront(10)
	v, err = d.PopBack()
	assert.Equal(t, 11, v)
	assert.Nil(t, err)
	v, err = d.PopFront()
	assert.Equal(t, 10, v)
	assert.Nil(t, err)
}

func Test_PushBackPushFrontPopBackPopFront2(t *testing.T) {
	d := NewDeq(2)
	d.PushFront(10)
	d.PopFront()
	d.PushFront(11)
	d.PopFront()
	d.PushFront(12)
	v, err := d.PopBack()
	assert.Equal(t, 12, v)
	assert.Nil(t, err)
}
