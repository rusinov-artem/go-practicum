package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanMerge(t *testing.T) {
	a := []int{1, 4, 9, 2, 10, 11}
	b := merge(a, 0, 3, 6)
	expected := []int{1, 2, 4, 9, 10, 11}
	assert.Equal(t, expected, b)
}

func Test_CanMerge2(t *testing.T) {
	a := []int{1, 4}
	b := merge(a, 0, 1, 2)
	expected := []int{1, 4}
	assert.Equal(t, expected, b)
}

func Test_CanMerge3(t *testing.T) {
	a := []int{4, 1}
	b := merge(a, 0, 1, 2)
	expected := []int{1, 4}
	assert.Equal(t, expected, b)
}

func Test_CanMerge4(t *testing.T) {
	a := []int{1, 3, 0, 9}
	b := merge(a, 2, 3, 4)
	expected := []int{0,  9}
	assert.Equal(t, expected, b)
}

func Test_CanSort(t *testing.T) {
	c := []int{1, 4, 2, 10, 1, 2}
	merge_sort(c, 0, 6)
	expected := []int{1, 1, 2, 2, 4, 10}
	assert.Equal(t, expected, c)
}
