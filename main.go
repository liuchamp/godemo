package main

import (
	"fmt"
	"runtime"
	"sync"
)

var couter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	couter++
	fmt.Println("counter = ", couter)
	lock.Unlock()
}

func main() {
	lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go Count(lock)

	}

	for {
		lock.Lock()
		c := couter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}

func add(x, y int) {
	z := x + y
	fmt.Println(z)
}
