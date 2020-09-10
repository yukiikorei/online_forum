/*
论坛系统主模块，读取配置，创建服务器，设置路由

create_time: 2020-9-9
author: korei
 */
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	router := gin.Default()
	router.LoadHTMLGlob("./template/*")

	//set static file system
	router.StaticFS("./static/",http.Dir("./static"))
	router.StaticFS("./css/",http.Dir("./css"))
	router.StaticFS("./javascript/",http.Dir("./javascript"))

	//static page get
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK,"index.html",gin.H{})
	})




	//fro manager
	managerRouter := router.Group("/manager")

	managerRouter.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK,"index.html",gin.H{})
	})



	//start service
	router.Run(":9999")
}
