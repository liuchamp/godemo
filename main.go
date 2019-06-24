package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {

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
