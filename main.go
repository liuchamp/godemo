package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	clousuretest2()
}

func chanstest1() {
	// 待处理的字符串列表
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}
	// 处理函数链
	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}
	// 处理字符串
	StringProccess(list, chain)
	// 输出处理好的字符串
	for _, str := range list {
		fmt.Println(str)
	}

}

// 字符串处理函数，传入字符串切片和处理链
func StringProccess(list []string, chain []func(string) string) {
	// 遍历每一个字符串
	for index, str := range list {
		// 第一个需要处理的字符串
		result := str
		// 遍历每一个处理链
		for _, proc := range chain {
			// 输入一个字符串进行处理，返回数据作为下一个处理链的输入。
			result = proc(result)
		}
		// 将结果放回切片
		list[index] = result
	}
}

// 自定义的移除前缀的处理函数
func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func anonymoustest() {
	func(data int) {
		fmt.Println("hello", data)
	}(100)
}

func an2func() {
	// 将匿名函数体保存到f()中
	f := func(data int) {
		fmt.Println("hello", data)
	}
	// 使用f()调用
	fmt.Printf("%T sn\n", f)
	f(100)
}

var skillParam = flag.String("skill", "", "skill to perform")

func an2op() {
	flag.Parse()
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
}

func clousuretest() {
	// 准备一个字符串
	str := "hello world"
	// 创建一个匿名函数
	foo := func() {

		// 匿名函数中访问str
		str = "hello dude"
	}
	// 调用匿名函数
	fmt.Println(str)
	foo()
	fmt.Println(str)
}

func clousuretest2() {
	// 创建一个累加器, 初始值为1
	accumulator := Accumulate(1)
	// 累加1并打印
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", accumulator)
	// 创建一个累加器, 初始值为1
	accumulator2 := Accumulate(10)
	// 累加1并打印
	fmt.Println(accumulator2())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", accumulator2)
}

// 提供一个值, 每次调用函数会指定对值进行累加
func Accumulate(value int) func() int {
	// 返回一个闭包
	return func() int {
		// 累加
		value++
		// 返回一个累加值
		return value
	}
}
