package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		go add(i, i+1)
	}
}

func add(x, y int) {
	z := x + y
	fmt.Println(z)
}
