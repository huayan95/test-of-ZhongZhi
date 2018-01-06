package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func Bootstraphtml(c *gin.Context){
	c.HTML(http.StatusOK, "bootstrap.html", gin.H{
		"title": "GIN: Bootstrap布局页面",
	})
}