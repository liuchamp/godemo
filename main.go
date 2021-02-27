package main

import (
	"fmt"
	"sync"
)

//Singleton 是单例模式类
type Singleton struct {
	c int
}

var singleton *Singleton
var once sync.Once

//GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}

func (st *Singleton) GetS() string {
	st.c = st.c + 1
	return fmt.Sprintf("sxx %d", st.c)
}

func main() {
	s1 := GetInstance()
	fmt.Println(s1.GetS())
	s2 := GetInstance()
	
	fmt.Println(s2.GetS())
	fmt.Println(s1.GetS())
	fmt.Println(s2.GetS())
}
