/**
 * @Author: korei
 * @Description: 论坛系统主模块，读取配置，创建服务器，设置路由
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2020/9/10 下午4:26
 */

package main

import (
	"github.com/gin-gonic/gin"
	"korei/onlineForum/controlFake"
	_ "korei/onlineForum/model"
	"korei/onlineForum/requestHandler"
	"net/http"
)

func main1()  {
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


	//for ADMIN
	adminRouter := router.Group("/admin")
	router.LoadHTMLGlob("./template/*")
	adminRouter.GET("", requestHandler.ControlPanel)
	adminRouter.GET("/login",requestHandler.LoginPage)
	adminRouter.POST("/login",requestHandler.LoginProcess)

	//start service
	router.Run(":9999")
}

/**
 *	this is a simple main function for test
 */
func main()  {
	print("begin main")
	controlFake.Test1()
}
