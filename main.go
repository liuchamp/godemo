package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
)

func main() {
	i := 0
	c := cron.New(cron.WithSeconds())
	spec := "*/15 * * * * ?"
	id, err := c.AddFunc(spec, func() {
		i++
		fmt.Println("cron running", i)
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
	c.Start()
	select {}
}
