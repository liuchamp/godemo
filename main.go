package main

import (
	"github.com/casbin/casbin"
	"github.com/gin-contrib/authz"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	// load the casbin model and policy from files, database is also supported.
	e, err := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")
	if err != nil {
		log.Fatal("find users")
	}
	// define your router, and use the Casbin authz middleware.
	// the access that is denied by authz will return HTTP 403 error.
	router := gin.New()
	router.Use(authz.NewAuthorizer(e))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
