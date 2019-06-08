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
		fmt.Println(err)
	}
}
