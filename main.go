package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
)

func main() {
	hans := "中国人"

	// 默认
	a := pinyin.NewArgs()
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zhong] [guo] [ren]]

	// 包含声调
	a.Style = pinyin.Tone
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zhōng] [guó] [rén]]

	// 声调用数字表示
	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zho1ng] [guo2] [re2n]]

	// 开启多音字模式
	a = pinyin.NewArgs()
	a.Heteronym = true
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zhong zhong] [guo] [ren]]
	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zho1ng zho4ng] [guo2] [re2n]]

	fmt.Println(pinyin.LazyPinyin(hans, pinyin.NewArgs()))
	// [zhong guo ren]

	fmt.Println(pinyin.Convert(hans, nil))
	// [[zhong] [guo] [ren]]

	fmt.Println(pinyin.LazyConvert(hans, nil))
	// [zhong guo ren]

	// 段落转换，支持完整支持多音字，保留符号
	fmt.Println(pinyin.Paragraph("交给团长，告诉他我们给予期望。前线的供给一定要能自给自足！"))
}
