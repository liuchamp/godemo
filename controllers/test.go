package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/liuchamp/godemo/services"
	"net/http"
)

type TestController struct {
	Service services.TestService
}

func (ts *TestController) Demo(c *gin.Context) {
	t := ts.Service.GetOne()
	c.JSON(http.StatusOK, t)
}
