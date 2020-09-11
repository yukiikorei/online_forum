/**
 * @Author: korei
 * @Description:	处理admin用户的请求，调用控制器中的逻辑,本页面中每一个函数前面应当声明路由相关
 * @File:  admin.go
 * @Version: 1.0.0
 * @Date: 2020/9/11 上午8:46
 */

package requestHandler

import (
	"github.com/gin-gonic/gin"
	"korei/onlineForum/controlFake"
	"net/http"
)

//FunctionName:	ControlPanel
//Path:			/admin/
//Method:		GET
//Description:	显示主页，如果用户没有登录，需要登录
func ControlPanel(context *gin.Context)  {
	//1.查看是否登录
	if !checkLogin(context){ return }
	//2.显示主页
	context.HTML(http.StatusOK,"controlpanel.html",gin.H{})
}

//FunctionName: LoginPage	
//Path:			/admin/login
//Method:		GET
//Description:  显示登录页面
func LoginPage(context *gin.Context)  {
	context.HTML(http.StatusOK,"adminlogin.html",gin.H{})
}

//FunctionName: LoginProcess
//Path:			/admin/login
//Method:		POST
//Description:  处理登录请求
func LoginProcess(context *gin.Context)  {
	//1.获取数据
	userID := context.PostForm("userID")
	password := context.PostForm("password")
	//2.登录
	cookie,ifSuc := controlFake.AdminLogin(userID, password)
	//3.响应
	if ifSuc {
		context.SetCookie("USER",cookie,3600,"/","127.0.0.1",false,true)
		context.Redirect(http.StatusMovedPermanently,"/admin")
	}else{
		context.HTML(http.StatusOK,"adminlogin.html",gin.H{"info":"帐号或密码错误"})
	}

}

//---------------------------sub function-----------------------------------------
//--------------------------------------------------------------------------------
//below are some sub function, which won't be used to handler request directly----
//--------------------------------------------------------------------------------

//FunctionName: CheckAdminLogin
//Path:			nil
//Method:		nil
//Description:  check if user has login as admin
func checkLogin(context *gin.Context) bool {
    cookie,err := context.Cookie("USER")
    if err==nil {
    	if controlFake.CheckAdminLogin(cookie) {
    		return true
    	}
	}
	context.Redirect(http.StatusMovedPermanently,"/admin/login")
	return false
}




