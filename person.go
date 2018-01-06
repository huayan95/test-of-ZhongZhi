package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//初始页面

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "GIN: HTML页面",
	})
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "GIN:  用户登录",
	})
}

func ShowHtmlPage(c *gin.Context) {
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"title": "GIN: HTML页面",
	})
}

func ListHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "list.html", gin.H{
		"title": "GIN: 列表页面",
	})
}

//文章页面
func TopicHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "views.html", gin.H{
		"title": "GIN: 浏览页面",
	})
}

//新增页面
func AddHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "add.html", gin.H{
		"title": "GIN: 新增用户页面",
	})
}
