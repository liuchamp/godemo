package main

import (
	"fmt"
)

func main() {

	impTest2()
}
func testsingleInherit() {
	son := Son{Base{"Start"}}
	//
	son.B()
	son.A()
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

// 名称冲突问题
type Se struct {
	Name string
}

func (s *Se) GetName() {
	fmt.Println("Se.Name=", s.Name)
}

type See struct {
	Se
	Name string
}

func (s *See) GetName() {
	fmt.Println("See.Name=", s.Name)
}

func testRepeatName() {
	s := Se{"siagnsd"}
	ss := See{s, "dsagdsina"}

	s.GetName()
	ss.GetName()
}

type Student struct {
	Name string
}

type Player struct {
	Student
}

type Action interface {
	sport()
	draw()
}

func (s Student) sport() {
	fmt.Println("sport")
}
func (p Player) sport() {
	fmt.Println("player sport")
}

func (s Student) draw() {

}

func impTest() {
	var a Action
	fmt.Println("a=", a)

	var s Student
	a = s
	fmt.Println(a)
}

func impTest2() {
	s := Student{"a"}
	p := Player{s}
	p.sport()
}
