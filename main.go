package main

import (
	"encoding/json"
	"fmt"
	"github.com/json-iterator/go"
	"github.com/liuchamp/godemo/ref"
	"gopkg.in/mgo.v2/bson"
	"reflect"
)

// 定义一个Enum类型
type Enum int

const (
	Zero Enum = 0
)

func main() {
	demo()
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
		Name string `json:"uml"`
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

	j, err := jsoniter.Marshal(d1)
	j, err = json.Marshal(d1)
	fmt.Println(string(j), err)
	j, err = json.Marshal(d1)
	fmt.Println(string(j), err)
	j, err = bson.Marshal(d1)
	fmt.Println(string(j), err)
}

func printChildByTag(d1 interface{}) {

}

func printchilds(d1 interface{}, prefix string) {
	taoftype := reflect.TypeOf(d1)
	fmt.Println(taoftype.Kind())
	fmt.Println(taoftype.NumField())

	for i := 0; i < taoftype.NumField(); i++ {
		fmt.Println(taoftype.Field(i))
	}

}

func demo() {
	p := ref.Person{"sddaf", 32}
	v := reflect.ValueOf(p)
	tf := reflect.TypeOf(p)
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		fmt.Println(f.String())
		// fs,_ :=tf.FieldByName(i)
		fs := tf.Field(i)
		switch f.Kind() {
		case reflect.String:
			fmt.Println(f.String())
			fmt.Println(fs.Tag.Get("mata"))

		case reflect.Int:
			fmt.Println(f.Int())
		}
	}
}
