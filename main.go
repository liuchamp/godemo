package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/liuchamp/godemo/model"
)

func main() {
	TestRedisStruct()
}

func TestRedisStruct() {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		// Password: "", // no password set
		DB: 0, // use default DB
	})
	u := model.User{
		Username: "test",
		RoleID:   "123548894641685165",
	}
	client.Set("s", &u, 3600000)
	us := new(model.User)
	er := client.Get("s").Scan(us)
	if er != nil {
		fmt.Println("s does not exist", er)
	}
	fmt.Println(us)
}

func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
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
		Addr: "127.0.0.1:6379",
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
		Addr: "127.0.0.1:6379",
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
