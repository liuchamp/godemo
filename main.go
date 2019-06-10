package main

import (
	"fmt"
	yrd "github.com/garyburd/redigo/redis"
	"github.com/go-redis/redis"
)

func main() {
	ExampleNewClient2()
}
func ExampleClientGary() {
	c, err := yrd.Dial("tcp", "192.168.0.193:7003")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := yrd.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
}
func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr: "192.168.0.193:7005",
		// Password: "", // no password set
		DB: 0, // use default DB
	})
	defer client.Close()
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>

	Get := func(client *redis.Client, key string) *redis.StringCmd {
		cmd := redis.NewStringCmd("get", key)
		client.Process(cmd)
		return cmd
	}

	v, err := Get(client, "key_does_not_exist").Result()
	fmt.Printf("%q %s \n", v, err)

	v, e := client.Do("get", "key_does_not_exist").String()
	fmt.Printf("%q %s \n", v, e)

}
func ExampleNewClient2() {
	client := redis.NewClient(&redis.Options{
		Addr: "192.168.0.193:7005",
		// Password: "", // no password set
		DB: 0, // use default DB
	})
	defer client.Close()

	SAVE := func(client *redis.Client, key string, value string) *redis.StringCmd {
		cmd := redis.NewStringCmd("set", key, value)
		client.Process(cmd)
		return cmd
	}

	v, err := SAVE(client, "key_does_not_exist", "smul").Result()
	fmt.Printf("%q %s \n", v, err)

	v, e := client.Do("get", "key_does_not_exist").String()
	fmt.Printf("%q %s \n", v, e)

}

func ExampleClient() {
	client := redis.NewClient(&redis.Options{
		Addr: "92.168.0.193:7003",
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
