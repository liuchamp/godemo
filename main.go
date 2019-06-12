package main

import "fmt"

func main() {
	son := Son{Base{"Start"}}
	//
	son.B()
	son.A()
	x := X{"sdaingsdia"}
	y := Y{}
	y.X = x
	fmt.Println(x.Name)
	fmt.Println(y.Name)
}

type Base struct {
	Name string
}

func (b *Base) A() {
	fmt.Println("Base method A called ......")
}
func (b *Base) B() {
	fmt.Println("Base method B called ......")
}

type Son struct {
	Base
}

// 方法重写
func (son *Son) B() {
	// 调用父类，相当于super
	son.Base.B()
	fmt.Println("Son method B called ......")
}
