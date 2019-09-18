package main

import (
	"fmt"
	"time"
)

func main() {
	bdemo()
}

func bdemo() {
	FirstNames := []string{"aaa", "bbb", "ccc"}
	LastNames := []string{"111", "222", "333"}

Loop:
	for _, firstName := range FirstNames {
		for _, lastName := range LastNames {
			fmt.Printf("Name: %s %s\n", firstName, lastName)
			time.Sleep(time.Second)
			if firstName == "bbb" && lastName == "111" {
				break Loop
			}
		}
	}
	println("Over.")
}

func demos() {
	FirstNames := []string{"aaa", "bbb", "ccc"}
	LastNames := []string{"111", "222", "333"}

Loop:
	for _, firstName := range FirstNames {
		for _, lastName := range LastNames {
			fmt.Printf("Name: %s %s %d\n", firstName, lastName, time.Now().Unix())
			time.Sleep(time.Second)
			if firstName == "bbb" && lastName == "111" {
				continue Loop
			}
		}
	}
	println("Over.")
}
