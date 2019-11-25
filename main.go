package main

import (
	"fmt"
	"time"
)

func main() {
	timelocal, err := time.LoadLocation("Asia/Chongqing")
	fmt.Println(err, timelocal)
	st := time.Now().In(timelocal)
	fmt.Println(st.String())
}

func smul(s1 struct{}) {
	ti := time.Now().UnixNano() / 1e6
	fmt.Println(ti)
}

func todaystartandend() {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	tm2 := tm1.AddDate(0, 0, 1)
	tm3 := tm1.AddDate(0, 0, -1)
	tm4 := tm1.AddDate(0, 0, -1000)

	fmt.Println("yesterday start", tm3.UnixNano()/1e6)
	fmt.Println("today start", tm1.UnixNano()/1e6)
	fmt.Println("today end", tm2.UnixNano()/1e6)
	fmt.Println("1000 days ago", tm4.UnixNano()/1e6)
}

func spliteDate(start int, end int, step int) map[int]bool {
	set := make(map[int]bool)
	for s := start; s < end; s += step {
		set[s] = true
	}
	if _, ok := set[end]; !ok {
		set[end] = true
	}
	return set
}
