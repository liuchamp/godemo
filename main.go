package main

import (
	"fmt"
	"time"
)

func main() {
	todaystartandend()
}

func smul(s1 struct{}) {
	ti := time.Now().UnixNano() / 1e6
	fmt.Println(ti)
}

func todaystartandend() {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	tm2 := tm1.AddDate(0, 0, 1)
	println(t.Location())
	fmt.Println(tm1.UnixNano()/1e6, tm2.UnixNano()/1e6)
}
