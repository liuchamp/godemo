package main

import (
	"fmt"
)

func Count(cha chan int) {
	cha <- 1
	fmt.Println("Counting")
}

func main() {

	ch := make(chan int)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
		case x := <-ch:
			fmt.Println(x)
		}
	}
}

func add(x, y int) {
	z := x + y
	fmt.Println(z)
}
