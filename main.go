package main

import (
	"fmt"
	"github.com/liuchamp/godemo/outtime"
	"time"
)

func main() {
	outtime.Outtimetest()
}
func smread() {
	queue := make(chan int, 10)
	queue <- 10
	queue <- 20

	close(queue)

	for i := 0; i < 10; i++ {
		fmt.Println(<-queue)
	}
}

func ti1() {
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
func ti() {
	//初始化定时器
	t := time.NewTimer(2 * time.Second)
	//当前时间
	now := time.Now()
	fmt.Printf("Now time : %v.\n", now)
	fmt.Printf("n type : %T.\n", 2)

	expire := <-t.C
	fmt.Printf("Expiration time: %v.\n", expire)
}
