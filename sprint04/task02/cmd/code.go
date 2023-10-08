package main

import (
	"fmt"
	"strings"
)

func main() {
	a := uint64(1000)
	m := uint64(123987123)

	vals := []uint64{}
	for i := uint64('a'); i <= uint64('z'); i++ {
		vals = append(vals, i)
	}

	h := make(map[uint64]string)
	generator := NewGenerator(6, vals)
	for {
		v := generator.Next()
		str := strings.Builder{}
		for _, r := range v {
			str.WriteRune(rune(r))
		}
		ns := str.String()
		nh := hash(ns, a, m)
		// fmt.Println(ns, nh)
		if other, ok := h[nh]; ok {
			fmt.Println(other)
			fmt.Println(ns)
			break
		}
		h[nh] = ns
	}
}

func hash(s string, a, m uint64) uint64 {
	hash := uint64(s[0]) % m
	for i := 1; i < len(s); i++ {
		hash = (hash*a + uint64(s[i])) % m
	}
	return hash
}


type Generator struct {
	Vals  []uint64
	State []uint64
}

func NewGenerator(n uint64, v []uint64) *Generator {
	state := make([]uint64, n)
	g := &Generator{
		Vals:  v,
		State: state,
	}

	return g
}

func (this *Generator) Next() []uint64 {
	r := make([]uint64, len(this.State))
	for i := 0; i < len(this.State); i++ {
		r[i] = this.Vals[this.State[i]]
	}
	this.inc(len(this.State) - 1)
	return r
}

func (this *Generator) inc(i int) {
	if i == -1 {
		for j := 0; j < len(this.State); j++ {
			this.State[j] = 0
		}
		return
	}
	this.State[i]++
	if this.State[i] >= uint64(len(this.Vals)) {
		this.State[i] = 0
		this.inc(i - 1)
	}
}
