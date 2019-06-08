package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	ExampleNewClient()
	ExampleClient()
}
func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr: "192.168.0.193:7001",
		// Password: "", // no password set
		DB: 0, // use default DB
	})
	defer client.Close()
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func ExampleClient() {
	client := redis.NewClient(&redis.Options{
		Addr: "92.168.0.193:7001",
		// Password: "", // no password set
		DB: 0, // use default DB
	})
	defer client.Close()
	err := client.Set("okmg", "signd", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("okmg").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
