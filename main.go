package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"unicode"
)

func main() {
	//通过复杂的方式解析字符串的首字母问
	fmdk := "似当年发动大"
	smdk := []rune(fmdk)
	fmt.Println(unicode.IsLetter(smdk[0]))
	fmt.Println(unicode.Is(unicode.Han, smdk[0]))

	if unicode.Is(unicode.Han, smdk[0]) {
		f := string(smdk[0])
		fs := pinyin.LazyConvert(f, nil)
		fmt.Println(fs)
	}
}
