package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuchamp/godemo/services"
	"io/ioutil"
	"net/http"
)

type TestController struct {
	Service services.TestService
}

func (ts *TestController) Demo(c *gin.Context) {
	t := ts.Service.GetOne()
	c.JSON(http.StatusOK, t)
}
func (ts *TestController) Menu(c *gin.Context) {
	getUser(c)
	t := ts.Service.GetOne()
	c.JSON(http.StatusOK, t)
}
func getUser(c *gin.Context) map[string]interface{} {
	var btm map[string]interface{}

	s, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(s, &btm)
	fmt.Println(btm)
	return btm
}
