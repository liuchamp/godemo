package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	urlDemo()
}

func baidurequ() {
	response, err := http.Get("http://www.baidu.com")
	if err != nil {
		// handle error
	}
	//程序在使用完回复后必须关闭回复的主体。
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

func postDemo() {
	body := "{\"action\":20}"
	res, err := http.Post("http://192.168.0.103:8080/users", "application/json;charset=utf-8", bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	fmt.Println(string(content))
}

func urlDemo() {
	u := url.URL{}
	pq := url.Values{}
	pq.Set("user", "champ")
	pq.Set("find", "或等OK")
	pq.Set("ok", "champ")
	u.Scheme = "http"
	u.Host = "127.0.0.1:8080"
	u.Path = "/users"
	u.RawQuery = pq.Encode()
	u.User = url.UserPassword("cahp", "dsagindsafs")
	fmt.Println(u.String())
	body := "{\"action\":20}"
	res, err := http.Post(u.String(), "application/json;charset=utf-8", bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	sim := u.Query()
	sim.Get("find")
	fmt.Println(string(content))
	fmt.Println(sim)
}
