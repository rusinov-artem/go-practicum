package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Дек сделан с помощью массива и 2х индексов
// индекс headIndex указывает на первый элемент
// индекс end указывает на элемент за последним
// при добавлении в начало двигается индекс begin
// при добавление в конец двигается индекс end
//
// Алгоритмическая сложность всех операций O(1)
// Время выполнения операций над стеком не зависит от 
// количества элементов содержащихся в стеке
//
// Пространственная сложность составляет O(n)
// Количество памяти необходимой для хранения элементов стека
// прямо пропорционально количеству элементов в стеке

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
			handlePushError(dq.PushBack(val))
		} else if cmd == "push_front" {
			sc.Scan()
			val, _ := strconv.Atoi(sc.Text())
			handlePushError(dq.PushFront(val))
		} else if cmd == "pop_back" {
			handlePopResult(dq.PopBack())
		} else if cmd == "pop_front" {
			handlePopResult(dq.PopFront())
		}
	}
}

func handlePushError(err error) {
	if err != nil {
		fmt.Println("error")
	}
}

func handlePopResult(val int, err error) {
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(val)
}

type Deq struct {
	capacity  int
	size      int
	data      []int
	begin int
	end int
}

func NewDeq(capacity int) *Deq {
	return &Deq{
		capacity: capacity,
		data:     make([]int, capacity),
	}
}

func inc(v, mod int) int{
	v++
	return v %mod
}

func dec(v, mod int) int {
	v--
	if v < 0 {
		v = mod+v
	}
	return v
}

func (t *Deq) PushBack(v int) error {
	if t.size+1 > t.capacity {
		return errors.New("overflow")
	}
	t.size++
	t.data[t.end] = v
	t.end = inc(t.end, t.capacity)
	return nil
}

func (t *Deq) PopFront() (int, error) {
	if t.size == 0 {
		return 0, errors.New("empty")
	}
	t.size--
	val := t.data[t.begin]
	t.begin = inc(t.begin, t.capacity)
	return val, nil
}

func (t *Deq) PopBack() (int, error) {
	if t.size == 0 {
		return 0, errors.New("empty")
	}
	t.end = dec(t.end, t.capacity)
	val := t.data[t.end]
	t.size--

	return val, nil
}

func (t *Deq) PushFront(v int) error {
	if t.size+1 > t.capacity {
		return errors.New("overflow")
	}
	t.begin = dec(t.begin, t.capacity)
	t.data[t.begin] = v
	t.size++
	return nil
}
