package main

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type User struct {
	Name string
	Age  int
}

func main() {
	usec := cache.New(5*time.Minute, 6*time.Minute)
	usec.Set("user1111", 1000, cache.DefaultExpiration)
	usec.Set("user1112", "1000", cache.DefaultExpiration)
	s, ok := usec.Get("user1111")
	if ok {
		println(s.(int))
	}
	s2, ok := usec.Get("user1112")
	if ok {
		println(s2.(string))
	}
	usec.Set("users", User{
		Name: "yyyyyy",
		Age:  56,
	}, cache.DefaultExpiration)

	s3, ok := usec.Get("users")
	if ok {
		uinfo := s3.(User)
		println(uinfo.Name)
		println("ok")
	}
}
