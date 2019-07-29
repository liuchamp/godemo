package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	demo2()
}

func demo1() {
	vYml := viper.New()
	vYml.SetConfigName("app")  // name of config file (without extension)
	vYml.AddConfigPath("conf") // path to look for the config file in
	//vYml.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	//vYml.AddConfigPath(".")              // optionally look for config in the working directory
	vYml.SetConfigType("yml")
	fmt.Println(vYml.ReadInConfig())
	urls := vYml.GetString("redis.urls")
	fmt.Println(urls)
}

func demo2() {
	vYml := viper.New()
	vYml.SetConfigName("application")
	vYml.AddConfigPath("conf")
	if err := vYml.ReadInConfig(); err != nil {
		log.Fatal("配置读写失败", err)
	}
	var deappconfig defaultAppConfig
	if err := vYml.UnmarshalKey("app", &deappconfig); err != nil {
		log.Fatal("配置解析失败", err)
	}
	v, _ := json.Marshal(deappconfig)
	fmt.Print(string(v))
}

type defaultAppConfig struct {
	Port int64
	Name string
	Tags []string
}

func demoEvn() {

	os.Setenv("SPF_PORT", "13")
	os.Setenv("SPF_NAME", "13")
	os.Setenv("SPF_TAGS", "13,dsaigsdaf")

	vEnv := viper.New()
	vEnv.SetEnvPrefix("spf")
	vEnv.AutomaticEnv()
	vEnv.BindEnv("port")
	vEnv.BindEnv("name")
	vEnv.BindEnv("tags")

	var deappconfig defaultAppConfig
	if err := vEnv.Unmarshal(&deappconfig); err != nil {
		log.Fatal("配置解析失败", err)
	}
	v, _ := json.Marshal(deappconfig)
	fmt.Print(string(v))
}
