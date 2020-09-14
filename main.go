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
	"html/template"
	_ "korei/onlineForum/config"
	"korei/onlineForum/controlFake"
	_ "korei/onlineForum/model"
	"korei/onlineForum/requestHandler"
	"net/http"
	"time"
)

func main()  {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"after":time.Time.After,
	})
	router.LoadHTMLGlob("./template/*")
	//router.HTMLRender=createMultiTemplateRender()

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
	//adminRouter.LoadHTMLGlob("./template/*")
	adminRouter.GET("", requestHandler.ControlPanel)
	adminRouter.GET("/block",requestHandler.BlockManage)
	adminRouter.POST("/block",requestHandler.BlockManagePost)
	//adminRouter.GET("/login",requestHandler.LoginPage)
	adminRouter.POST("/login",requestHandler.LoginProcess)
	adminRouter.GET("/user",requestHandler.UserManageGet)
	adminRouter.POST("/user",requestHandler.UserManagePost)
	//start service
	router.Run(":9999")
}


/**
 *	this is a simple main function for test
 */
func main1()  {
	print("begin main")
	controlFake.Test1()
}
