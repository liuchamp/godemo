package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {

	sdfsa := []ChanRegissdfassdfa{
		{CRID: 1,
			Channel: "sdafdsa",
			Enable:  true,
		},
	}
	bst, _ := json.Marshal(sdfsa)
	fmt.Println(string(bst))
}

type ChanRegissdfassdfa struct {
	CRID    int64  `json:"id" bson:"_id"`
	Channel string `json:"channel" bson:"channel"` // 渠道名称
	Enable  bool   `json:"enable" bson:"enable"`   // 开关
}

func deom() {
	vYml, err := newFileViper("conf", "setting", "json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var csds []ChanRegissdfassdfa
	err = vYml.UnmarshalKey("channelReg", &csds)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(csds) > 0 {
		for _, v := range csds {
			fmt.Println(v)
		}
	}

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

type appConfig struct {
	Name    string
	Port    int
	Tags    []string
	Include []string
}

func getAppconfig() {
	ymlViper, err := newFileViper("conf", "application", "yml")
	if err != nil {
		log.Fatal(err)
	}
	var deappconfig appConfig
	if err := ymlViper.UnmarshalKey("app", &deappconfig); err != nil {
		log.Fatal("配置解析失败", err)
	}
	fmt.Println(deappconfig)
}

func newFileViper(path, name, fileType string) (*viper.Viper, error) {
	vpApp := viper.New()
	vpApp.AddConfigPath(path)
	vpApp.SetConfigName(name)
	vpApp.SetConfigType(fileType)
	err := vpApp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return vpApp, nil
}
