package routers

import (
  "github.com/gin-gonic/gin"
  . "project1/apis" 
  . "project1/controllers" 
 )

func InitRouter() *gin.Engine{   
  router := gin.Default()                
  router.GET("/", HomePage)              
  router.GET("/login",LoginPage)       
  router.LoadHTMLGlob("views/*")            
  
  //add
  router.GET("/index", ShowHtmlPage)   
  router.GET("/category", ListHtml)        
  router.POST("/PageData", GetDataList)    
  router.POST("/PageNextData", PageNextData)
  router.GET("/topic",TopicHtml)
  router.GET("/add",AddHtml)
  router.POST("/saveadd", AddPersonApi)
  
   //edit
  router.GET("/edit", EditHtml)
  router.POST("/saveedit", EditPersonApi)

   //delete
  router.POST("/delete", DeletePersonApi)

   //Bootstrap and Content
  router.GET("/bootstrap", Bootstraphtml)
  router.GET("/content", Contenthtml)
    
    return router
 }