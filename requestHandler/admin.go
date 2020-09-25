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
	"korei/onlineForum/control"
	"korei/onlineForum/model"
	"net/http"
	"strconv"
	"time"
)

//FunctionName:	ControlPanel
//Path:			/admin/
//Method:		GET
//Description:	显示主页，如果用户没有登录，需要登录
func ControlPanel(context *gin.Context)  {
	//1.查看是否登录
	if !checkLogin(context){ return }
	//2.查看参数选择，否则直接显示主页
	context.HTML(http.StatusOK,"controlpanel.html",gin.H{})
}

//FunctionName: BlockManage
//Path:			/admin/block
//Method:		GET
//Description:  分论坛和板块管理界面
func BlockManage(context *gin.Context)  {
	//1.查看是否登录
	if !checkLogin(context){ return }
	//2.处理
	context.HTML(http.StatusOK,"blockmanage.html",
		control.GetForums())
}

//FunctionName: BlockManagePost
//Path:			/admin/block
//Method:		POST
//Description:  处理分论坛和板块管理的更新请求
func BlockManagePost(context *gin.Context) {
	//TODO: 把ACTION嵌入表单里面
	if !checkLogin(context){ return }
	action, ok := context.GetQuery("action")
	if !ok {
		context.String(http.StatusBadRequest, "post without param \"action\" ")
	}
	if action == "addforum" {
		name := context.PostForm("name")
		err := control.AddSubForums(name)
		if err == nil {
			context.HTML(http.StatusOK, "blockmanage.html",
				control.GetForums())
		} else {
			context.String(http.StatusBadRequest, err.Error())
		}
	}else if action=="deleteforum"{
		name := context.PostForm("name")
		err := control.DeleteSubForums(name)
		if err == nil {
			context.HTML(http.StatusOK, "blockmanage.html",
				control.GetForums())
		} else {
			context.String(http.StatusBadRequest, err.Error())
		}
	}else if action=="addblock"{
		forumName 	:= context.PostForm("forumName")
		name 		:= context.PostForm("name")
		masterID 	:= context.PostForm("masterID")
		if err := control.AddBlock(forumName,name,masterID); err == nil{
			context.HTML(http.StatusOK,"blockmanage.html",
				control.GetForums())
		} else {
			context.String(http.StatusBadRequest,err.Error())
		}
	}else if action=="deleteblock"{
		blockID,_ := strconv.ParseUint(context.PostForm("blockID"),10,64)
		err := control.DeleteBlock(uint(blockID))
		if err == nil {
			context.HTML(http.StatusOK, "blockmanage.html",
				control.GetForums())
		} else {
			context.String(http.StatusBadRequest, err.Error())
		}
	}else if action == "changeMaster"{
		blockID,_ := strconv.ParseUint(context.PostForm("blockID"),10,64)
		masterID := context.PostForm("masterID")
		err := control.ChangeMasterID(uint(blockID),masterID)
		if err == nil {
			context.HTML(http.StatusOK, "blockmanage.html",
				control.GetForums())
		} else {
			context.String(http.StatusBadRequest, err.Error())
		}
	} else {
		context.String(http.StatusBadRequest,"post without param \"action\" ")
	}
}


//FunctionName: UserManageGet
//Path:			/admin/user
//Method:		GET
//Description:  返回用户管理界面
func UserManageGet(context *gin.Context)  {
	if !checkLogin(context){ return }
	context.HTML(http.StatusOK,"usermanage.html",
		gin.H{"users":[]model.User{}})
}

//FunctionName: UserManagePost
//Path:			/admin/user
//Method:		Post
//Description:  用户管理的更新和查询请求
func UserManagePost(context *gin.Context)  {
	if !checkLogin(context){ return }
	action := context.PostForm("action")
	if action=="getUsers"{
		userName := context.PostForm("userName")
		userID := context.PostForm("userID")
		IP:= context.PostForm("IP")
		email := context.PostForm("email")
		users := control.SearchUser(userID,userName,IP,email)
		context.HTML(http.StatusOK,"usermanage.html",
			gin.H{"users":users,"now":time.Now()})
	}else if action=="getFrozen"{
		userName := context.PostForm("userName")
		userID := context.PostForm("userId")
		IP:= context.PostForm("IP")
		email := context.PostForm("email")
		users := control.SearchFrozenUser(userID,userName,IP,email)
		context.HTML(http.StatusOK,"usermanage.html",
			gin.H{"users":users,"now":time.Now()})
	}else if action=="freezeUser"{
		userID := context.PostForm("userID")
		year,_ := strconv.Atoi(context.PostForm("year"))
		month,_ := strconv.Atoi(context.PostForm("month"))
		day,_ := strconv.Atoi(context.PostForm("day"))
		hour,_ := strconv.Atoi(context.PostForm("hour"))
		min,_ := strconv.Atoi(context.PostForm("min"))
		second := ((((year*365)+month*12 +day)*24+hour)*60+min)*60
		duration := time.Duration(int64(second)*1000000000)
		control.FreezeUser(userID,duration)
	}else if action=="thaw"{
		userID := context.PostForm("userID")
		control.ThawUser(userID)

	} else{
		context.String(http.StatusBadRequest,"Bad request without \"action\" in post form")
	}
}

/*
//FunctionName: LoginPage	
//Path:			/admin/login
//Method:		GET
//Description:  显示登录页面
func LoginPage(context *gin.Context)  {
	context.HTML(http.StatusOK,"adminlogin.html",gin.H{"info":""})
}
 */

//FunctionName: LoginProcess
//Path:			/admin/login
//Method:		POST
//Description:  处理登录请求
func LoginProcess(context *gin.Context)  {
	//1.获取数据
	userID := context.PostForm("userID")
	password := context.PostForm("password")
	//2.登录
	token,err := control.LoginAsAdmin (userID,password)
	//3.响应
	if err==nil {
		context.SetCookie("ADMIN",token,3600,"/","127.0.0.1",false,true)
		//context.Redirect(http.StatusMovedPermanently,"/admin")
		context.Redirect(http.StatusMovedPermanently,"/admin")
	}else{
		//TODO: 去掉登录页面路径，改成刷新当前页
		context.HTML(http.StatusOK,"adminlogin.html",gin.H{"info":err.Error()})
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
    token,err := context.Cookie("ADMIN")
    if err==nil {
    	ui,err :=control.ParseToken(token)
    	if err == nil {
    		if ui.Role == "admin"{
    			return true
			}
		}
	}
	context.HTML(http.StatusOK,"adminlogin.html",gin.H{"info":"请先登录再操作"})
	return false
}




