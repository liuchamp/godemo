package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liuchamp/godemo/controllers"
	"github.com/liuchamp/godemo/services"
)

func main() {

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	s := services.TestService{}
	c := controllers.TestController{Service: s}
	router.GET("test", c.Demo)
	router.GET("sim", c.Deok)
	router.PUT("test", c.Menu)
	router.Run()
}
