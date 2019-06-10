package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	redisdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"192.168.0.193:7001",
			"192.168.0.193:7002",
			"192.168.0.193:7003",
			"192.168.0.193:7004",
			"192.168.0.193:7005",
			"192.168.0.193:7006"},
	})
	redisdb.Ping()
	err := redisdb.ReloadState()
	if err != nil {
		panic(err)
	}
	defer redisdb.Close()
	v, e := redisdb.Do("get", "key_does_not_exist").String()
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(v)
	}
	v, es := redisdb.Do("set", "name", "yexingyun").String()

	if es != nil {
		fmt.Println(e)
	} else {
		fmt.Println(v)
	}
	if m, ess := redisdb.Do("get", "name").String(); ess == nil {
		fmt.Println(m)
	}
	if m, esss := redisdb.Do("get", "key_does_not_exist1").String(); esss == nil {
		fmt.Println(m)
	} else {
		fmt.Println("error", esss)
	}
}
