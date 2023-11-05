package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Хешмап.
// Алгоритмическая сложность всех операций в среднем
// O(1) При наполнении большим количеством элементов
// количество колизий возрастает и сложность превращается
// в O(n)
//
// Для хранения элементов расходуется пропорциональное
// количество памяти, поэтому
// сложность по памяти O(n)
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 4*1024*1024), 0)
	sc.Split(bufio.ScanWords)

	hashMap := NewHashMap()
	sb := &strings.Builder{}

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		sc.Scan()
		command := sc.Text()
		sc.Scan()
		key, _ := strconv.Atoi(sc.Text())
		switch command {
		case "get":
			{
				val, err := hashMap.Get(key)
				printResult(val, err, sb)
			}
		case "put":
			{
				sc.Scan()
				val, _ := strconv.Atoi(sc.Text())
				hashMap.Put(key, val)
			}
		case "delete":
			{
				val, err := hashMap.Delete(key)
				printResult(val, err, sb)
			}
		}
	}
	fmt.Print(sb.String())
}

func printResult(val int, err error, sb *strings.Builder) {
	if err != nil {
		sb.WriteString("None\n")
	} else {
		sb.WriteString(fmt.Sprintf("%d\n", val))
	}
}

type HashMap struct {
	storage []*Element
	Hash    func(int) int
}

type Element struct {
	Next  *Element
	Key   int
	Value int
}

func (this Element) hasKey(key int) bool {
	return this.Key == key
}

func NewElement(key, val int) *Element {
	return &Element{
		Next:  nil,
		Key:   key,
		Value: val,
	}
}

func scan(head *Element, fn func(*Element) bool) *Element {
	current := head
	for {
		if fn(current) {
			return current
		}

		if current.Next == nil {
			return nil
		}

		current = current.Next
	}
}

func (this *HashMap) Put(key, val int) {
	k := this.Hash(key)
	if this.storage[k] == nil {
		this.storage[k] = NewElement(key, val)
		return
	}

	var last *Element
	found := scan(this.storage[k], func(el *Element) bool {
		last = el
		if el.Key == key {
			el.Value = val
			return true
		}
		return false
	})

	if found == nil {
		last.Next = NewElement(key, val)
	}
}

func (this *HashMap) Get(key int) (int, error) {
	k := this.Hash(key)
	if this.storage[k] == nil {
		return 0, errors.New("Not found")
	}

	found := scan(this.storage[k], func(el *Element) bool {
		if el.Key == key {
			return true
		}
		return false
	})

	if found != nil {
		return found.Value, nil
	}

	return 0, errors.New("Not found")
}

func (this *HashMap) Delete(key int) (int, error) {
	k := this.Hash(key)
	if this.storage[k] == nil {
		return 0, errors.New("Not found")
	}

	prev := this.storage[k]
	found := scan(this.storage[k], func(el *Element) bool {
		if el.Key == key {
			return true
		}
		prev = el
		return false
	})

	if found != nil {
		if found == this.storage[k] && found.Next == nil {
			this.storage[k] = nil
		} else {
			prev.Next = found.Next
		}
		return found.Value, nil
	}

	return 0, errors.New("Not found")
}

func NewHashMap() *HashMap {
	h := &HashMap{
		storage: make([]*Element, 10000),
	}

	h.Hash = func(key int) int {
		r := key % len(h.storage)
		if r < 0 {
			return -r
		}
		return r
	}

	return h
}
