package main

import (
	"fmt"
	"github.com/Masterminds/goutils"
)

func main() {
	s := make(map[string]int64)
	ddms(s)
	prting(s)
	ddms(s)
	prting(s)
	ddms(s)
	prting(s)
}

func ddms(ks map[string]int64) error {
	smk, _ := goutils.RandomAlphabetic(9)
	ks[smk] = 8989
	return nil
}

func prting(ks map[string]int64) {
	l := len(ks)
	fmt.Println(l)
}
