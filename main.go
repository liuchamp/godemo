package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool, 1)
	ch := make(chan int)
	go func() {
		time.Sleep(1e9) // sleep one second
		ch <- 1
		timeout <- true
	}()
	select {
	case u := <-ch:
		fmt.Println(u)
	case <-timeout:
		fmt.Println("timeout!")
	}
}
