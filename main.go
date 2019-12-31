package main

import (
	"github.com/patrickmn/go-cache"
	"time"
)

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
}
