package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Алгоритм описан в задаче
// https://contest.yandex.ru/contest/22781/problems/B/
// Корректность решения продемонстрирована
// тестами в фалйе code_test.go
//
// Алгоритмическая сложность линейная O(n) где n - количество
// чисел над которыми должны быть выполнены операции
//
// Пространственная сложность тоже линейная O(n) где n -
// количество чисел над которыми выполняются операции. В худшем
// случае в стеке придется держать все числа

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 10000), 0)
	sc.Split(bufio.ScanWords)
	expr := NewExpression()
	for sc.Scan() {
		token := sc.Text()
		switch token {
		case "+":
			expr.Sum()
		case "-":
			expr.Sub()
		case "*":
			expr.Mult()
		case "/":
			expr.Div()
		default:
			val, _ := strconv.Atoi(sc.Text())
			expr.PushValue(val)
		}
	}
	fmt.Println(expr.Result())
}

type MyStack struct {
	data []int
}

func NewMyStack() *MyStack {
	return &MyStack{
		data: make([]int, 0, 1000),
	}
}

func (t *MyStack) Push(val int) {
	t.data = append(t.data, val)
}

func (t *MyStack) Pop() (int, error) {
	if len(t.data) == 0 {
		return 0, errors.New("empty")
	}

	val := t.data[len(t.data)-1]
	t.data = t.data[:len(t.data)-1]

	return val, nil
}

type Expression struct {
	stack MyStack
}

func NewExpression() *Expression {
	return &Expression{}
}

func (t *Expression) PushValue(val int) {
	t.stack.Push(val)
}

func (t *Expression) Result() int {
	val, _ := t.stack.Pop()
	return val
}

func (t *Expression) Sum() {
	t.update(func(v1, v2 int) int { return v1 + v2 })
}

func (t *Expression) Mult() {
	t.update(func(v1, v2 int) int { return v1 * v2 })
}

func (t *Expression) Sub() {
	t.update(func(v1, v2 int) int { return v1 - v2 })
}

func (t *Expression) Div() {
	t.update(func(v1,v2 int)int {
		res := float64(v1) / float64(v2)
		if res > 0 {
			return int(math.Trunc(res))
		} else {
			return int(math.Floor(res))
		}
	})
}

func (t *Expression) update(op func(v1, v2 int) int) {
	v2, _ := t.stack.Pop()
	v1, _ := t.stack.Pop()
	t.stack.Push(op(v1, v2))
}
