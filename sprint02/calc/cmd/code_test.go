package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanCreateExpression(t *testing.T) {
	expr := NewExpression()
	assert.NotNil(t, expr)
}

func Test_CanPushValues(t *testing.T) {
	expr := NewExpression()
	expr.PushValue(3)
}

func Test_CanSumValues(t *testing.T) {
	expr := NewExpression()
	expr.PushValue(5)
	expr.PushValue(4)
	expr.Sum()
	assert.Equal(t, 9, expr.Result())
}

func Test_CanMultiplyValues(t *testing.T) {
	expr := NewExpression()
	expr.PushValue(5)
	expr.PushValue(4)
	expr.Mult()
	assert.Equal(t, 20, expr.Result())
}

func Test_CanDivide(t *testing.T) {
	expr := NewExpression()
	expr.PushValue(5)
	expr.PushValue(4)
	expr.Div()
	assert.Equal(t, 1, expr.Result())
}

func Test_CanDivideNegative(t *testing.T) {
	expr := NewExpression()
	expr.PushValue(-1)
	expr.PushValue(3)
	expr.Div()
	assert.Equal(t, -1, expr.Result())
}

func Test_DivisionCorrect(t *testing.T){
	// 2 1 2 / *
	expr := NewExpression()
	expr.PushValue(2)
	expr.PushValue(1)
	expr.PushValue(2)
	expr.Div()
	expr.Mult()
	assert.Equal(t, 0, expr.Result())
}

func Test_CanSubtract(t *testing.T) {
	expr := NewExpression()
	expr.PushValue(6)
	expr.PushValue(4)
	expr.Sub()
	assert.Equal(t, 2, expr.Result())
}

func Test_CanUseMultipleOpperands(t *testing.T) {
	// 10 2 4 * -
	expr := NewExpression()
	expr.PushValue(10)
	expr.PushValue(2)
	expr.PushValue(4)
	expr.Mult()
	expr.Sub()
	assert.Equal(t, 2, expr.Result())
}

func Test_StackWorking(t *testing.T) {
	var v int
	var err error
	s := NewMyStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	v, err = s.Pop()
	assert.Equal(t, 3, v)
	v, err = s.Pop()
	assert.Equal(t, 2, v)
	v, err = s.Pop()
	assert.Equal(t, 1, v)
	v, err = s.Pop()
	assert.Error(t, err)
}
