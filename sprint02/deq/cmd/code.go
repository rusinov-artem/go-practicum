package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	var c int
	fmt.Scanf("%d\n%d", &n, &c)
	dq := NewDeq(c)
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 10000), 0)
	sc.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		sc.Scan()
		cmd := sc.Text()
		if cmd == "push_back" {
			sc.Scan()
			val, _ := strconv.Atoi(sc.Text())
			err := dq.PushBack(val)
			if err != nil {
				fmt.Println("error")
			}
		} else if cmd == "push_front" {
			sc.Scan()
			val, _ := strconv.Atoi(sc.Text())
			err := dq.PushFront(val)
			if err != nil {
				fmt.Println("error")
			}
		} else if cmd == "pop_back" {
			val, err := dq.PopBack()
			if err != nil {
				fmt.Println("error")
				continue
			}
			fmt.Println(val)
		} else if cmd == "pop_front" {
			val, err := dq.PopFront()
			if err != nil {
				fmt.Println("error")
				continue
			}
			fmt.Println(val)
		}
	}

}

type Deq struct {
	capacity int
	size     int
	data     []int
	begin    int
	end      int
}

func NewDeq(capacity int) *Deq {
	return &Deq{
		capacity: capacity,
		data:     make([]int, capacity),
	}
}

func (t *Deq) PushBack(v int) error {
	if t.size+1 > t.capacity {
		return errors.New("overflow")
	}
	t.size++
	t.data[t.end] = v
	t.end++
	t.end = t.end % t.capacity
	return nil
}

func (t *Deq) PushFront(v int) error {
	if t.size+1 > t.capacity {
		return errors.New("overflow")
	}
	if t.size == 0 {
		t.begin=0
		t.end = 1
		t.size=1
		t.data[t.begin] = v
		return nil
	}
	t.begin--
	if t.begin < 0 {
		t.begin = t.capacity + t.begin
	}
	t.data[t.begin] = v
	t.size++
	return nil
}

func (t *Deq) PopBack() (int, error) {
	if t.size == 0 {
		return 0, errors.New("empty")
	}
	t.end--
	if t.end < 0 {
		t.end = t.capacity + t.end
	}
	val := t.data[t.end]
	t.size--

	return val, nil
}

func (t *Deq) PopFront() (int, error) {
	if t.size == 0 {
		return 0, errors.New("empty")
	}
	t.size--
	val := t.data[t.begin]
	t.begin++
	t.begin = t.begin % t.capacity
	return val, nil
}