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
		switch command {
		case "get":
			{
				sc.Scan()
				key, _ := strconv.Atoi(sc.Text())
				val, err := hashMap.Get(key)
				printResult(val, err, sb)
			}
		case "put":
			{
				sc.Scan()
				key, _ := strconv.Atoi(sc.Text())
				sc.Scan()
				val, _ := strconv.Atoi(sc.Text())
				hashMap.Put(key, val)
			}
		case "delete":
			{
				sc.Scan()
				key, _ := strconv.Atoi(sc.Text())
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

func NewElement(key, val int) *Element {
	return &Element{
		Next:  nil,
		Key:   key,
		Value: val,
	}
}

func (this *HashMap) Put(key, val int) {
	k := this.Hash(key)
	if this.storage[k] == nil {
		this.storage[k] = NewElement(key, val)
		return
	}

	current := this.storage[k]
	for {
		if current.Key == key {
			current.Value = val
			return
		}

		if current.Next == nil {
			break
		}

		current = current.Next
	}

	current.Next = NewElement(key, val)
}

func (this *HashMap) Get(key int) (int, error) {
	k := this.Hash(key)
	if this.storage[k] == nil {
		return 0, errors.New("Not found")
	}

	current := this.storage[k]
	for {
		if current.Key == key {
			return current.Value, nil
		}

		if current.Next == nil {
			return 0, errors.New("Not found")
		}

		current = current.Next
	}
}

func (this *HashMap) Delete(key int) (int, error) {
	k := this.Hash(key)
	head := this.storage[k]
	element := this.storage[k]
	if element == nil {
		return 0, errors.New("Not found")
	}

	var prev *Element
	for {
		if element.Key == key {
			if element == head {
				this.storage[k] = element.Next
				return element.Value, nil
			} else {
				prev.Next = element.Next
				return element.Value, nil
			}
		}

		if element.Next == nil {
			return 0, errors.New("Not found")
		}

		prev = element
		element = element.Next
	}

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
