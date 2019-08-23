package main

import (
	"encoding/json"
	"fmt"
	"github.com/Masterminds/goutils"
)

func main() {
	Muts()
}

type NameStdt struct {
	D string
	E string
}

// 验证对象clone
func Muts() {
	sd := new(NameStdt)
	sd.D = "dsafdsa"
	sd.E = "dsafdsi"

	od(sd)
	sdbs, _ := json.Marshal(sd)
	fmt.Println(string(sdbs))
	od(sd)
	sdbs, _ = json.Marshal(sd)
	fmt.Println(string(sdbs))
}

func od(s *NameStdt) {
	smk, _ := goutils.RandomAlphabetic(9)
	s.E = smk
	smd, _ := goutils.RandomAlphabetic(9)
	s.D = smd
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
