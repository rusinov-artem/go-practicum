package main

import (
	"testing"
	"time"
)

var m = make(map[int]int)

func Test_Map(t *testing.T) {
	go read()
	go write()
	time.Sleep(time.Second)
}

func read() {
	for {
		_, _ = m[1]
	}
}

func write() {
	for {
		m[1] = 1
	}
}
