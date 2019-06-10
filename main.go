package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strings"
	"time"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	test1()
}
func test1() {
	session, err := mgo.Dial("mongodb://192.168.0.193:27017,mongodb://192.168.0.193:27018,mongodb://192.168.0.193:27019")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}

func test2() {
	host := "mongodb://192.168.0.193:27017,mongodb://192.168.0.193:27018,mongodb://192.168.0.193:27019"
	hosts := strings.Split(host, ",")
	mongoDBDialInfo := mgo.DialInfo{
		Addrs:    hosts,
		Timeout:  60 * time.Second,
		Database: "yexingyun",
		Username: "",
		Password: "",
	}
	var err error
	mongoSession, err := mgo.DialWithInfo(&mongoDBDialInfo)
	if err != nil {
		panic(err)
	}
	defer mongoSession.Close()
	mongoSession.SetMode(mgo.Monotonic, true)
	c := mongoSession.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}
