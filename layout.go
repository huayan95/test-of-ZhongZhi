package controllers

import (
	"fmt"
	"log"
	"html/template"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Contenthtml(c *gin.Context){

    //模板文件的拼接
	t, err := template.ParseFiles("views/layout.html", "views/T.header.tpl", 
		"views/T.navbar.html","views/sidebar.tpl","views/scripts.tpl")

    if err != nil {
        log.Fatal(err)
    }
	fmt.Println(t)
	//渲染html文件
	c.HTML(http.StatusOK,"layout.html", gin.H{
		"title": "布局页面",
	})
}