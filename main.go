package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//定义conf类型
//类型里的属性，全是配置文件里的属性
type conf struct {
	Host   string `yaml: "host"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
}

func main() {
	var c conf
	//读取yaml配置文件
	conf := c.getConf()
	fmt.Println(conf)

	//将对象，转换成json格式
	data, err := json.Marshal(conf)

	if err != nil {
		fmt.Println("err:\t", err.Error())
		return
	}

	//最终以json格式，输出
	fmt.Println("data:\t", string(data))
}

//读取Yaml配置文件,
//并转换成conf对象
func (c *conf) getConf() *conf {
	//应该是 绝对地址
	yamlFile, err := ioutil.ReadFile("E:\\Program\\go2\\goPath\\src\\xingej-go\\xingej-go\\xingej-go666\\lib\\yaml\\conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		fmt.Println(err.Error())
	}

	return c
}
