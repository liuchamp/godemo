package main

import (
	"fmt"
	"strings"
)

func main() {
	test2()
}

func test1() {
	s := "1231"
	strings.Join([]string{s, "ewrwer"}, ",")
	fmt.Println(s)
}

func test2() {
	var s []string
	s = append(s, "123123")
	s = append(s, "sjadingisdajf")
	sm := strings.Join(s, ",")
	fmt.Println(sm)
}
