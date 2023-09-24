package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 10000), 0)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	size, _ := strconv.Atoi(sc.Text())

	q := NewMyQueue(size)

	for i := 0; i < n; i++ {
		sc.Scan()
		cmd := sc.Text()
		if cmd == "push" {
			sc.Scan()
			val, _ := strconv.Atoi(sc.Text())
			err := q.push(val)
			if err != nil {
				fmt.Println("error")
				continue
			}
		} else if cmd == "pop" {
			val, err := q.pop()
			if err != nil {
				fmt.Println("None")
				continue
			}
			fmt.Println(val)
		} else if cmd == "size" {
			fmt.Println(q.size)
		} else if cmd == "peek" {
		   val, err := q.peek()
			if err != nil {
				fmt.Println("None")
				continue
			}
			fmt.Println(val)
		}
	}

}

func NewMyQueue(size int) *myQueue {
	return &myQueue{
		data: make([]int, size, size),
	}
}

type myQueue struct {
	data []int
	head int
	tail int
	size int
}

func (q *myQueue) push(val int) error {
	if q.size == len(q.data) {
		return errors.New("overflow")
	}

	q.data[q.tail] = val
	q.size++
	q.tail++
	q.tail = q.tail % len(q.data)

	return nil
}

func (q *myQueue) pop() (int, error) {
	if q.size == 0 {
		return 0, errors.New("empty")
	}

	val := q.data[q.head]
	q.head++
	q.head = q.head % len(q.data)
	q.size--
	return val, nil

}

func (q *myQueue) peek() (int, error) {
	if q.size == 0 {
		return 0, errors.New("empty")
	}

	val := q.data[q.head]
	return val, nil
}
