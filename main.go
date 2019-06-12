package main

import (
	"fmt"
)

func Count(cha chan int) {
	cha <- 1
	fmt.Println("Counting")
}

func main() {

	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	var cout int
	for _, ch := range chs {

		cout += <-ch
		fmt.Println(cout)
	}
}

func add(x, y int) {
	z := x + y
	fmt.Println(z)
}
