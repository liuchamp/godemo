package main

import (
	"fmt"
	"reflect"
)

// 定义一个Enum类型
type Enum int

const (
	Zero Enum = 0
)

func main() {
	demo2()
}
func demo1() {
	// 声明一个空结构体
	type cat struct {
	}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(cat{})
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())
	// 获取Zero常量的反射类型对象
	typeOfA := reflect.TypeOf(Zero)
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfA.Name(), typeOfA.Kind())

}

// 通过反射获取具体值
func demo2() {
	type Cat struct {
		Name string
		Uo   string `json:"uo"`
		Next *Cat   `json:"child"`
	}
	d1 := Cat{
		Name: "sdafgsdaf",
		Uo:   "dsadsa",
		Next: &Cat{
			Name: "1欧克",
			Uo:   "第三范德萨",
		},
	}

	printChildByTag(&d1, "")

}

func printChildByTag(d1 interface{}, prefix string) {
	taofcat := reflect.ValueOf(d1)
	fmt.Println(taofcat.Type().Kind())
	if taofcat.Kind() == reflect.Ptr {
		taofcat = taofcat.Elem()
	}

	for i := 0; i < taofcat.NumField(); i++ {
		childType := taofcat.Field(i)
		if childType.Type().Kind() == reflect.Ptr {
			if childType.IsNil() {
				continue
			}
			printChildByTag(childType.Interface(), "")
		} else if childType.Type().Kind().String() == "struct" {
			if childType.IsValid() {
				continue
			}
			printChildByTag(childType.Interface(), "")
		} else {
			fmt.Println(childType.Type().Name(), childType.Kind(), childType)
		}
	}
}
