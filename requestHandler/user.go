/**
 * @Author: korei
 * @Description: 响应用户的请求
 * @File:  user.go
 * @Version: 1.0.0
 * @Date: 2020/9/18 上午8:46
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
//------------------------------------------------------------
//------------------------------------------------------------
//FunctionName: UserLoginProcess
//Path:
//Method:		Post
//Description:  login
func UserLoginProcess(context *gin.Context)  {
	//1.获取数据
	userID := context.PostForm("userID")
	password := context.PostForm("password")
	//2.登录
	token,err := control.Login(userID,password)
	//3.响应
	if err==nil {
		context.SetCookie("USER",token,3600,"/","127.0.0.1",false,true)
		//context.Redirect(http.StatusMovedPermanently,"/admin")
		context.Redirect(http.StatusMovedPermanently,"/")
	}else{
		//TODO: 去掉登录页面路径，改成刷新当前页
		context.HTML(http.StatusOK,"login.html",gin.H{"info":err.Error()})
	}
}

//FunctionName: HomePackageGet
//Path:			/
//Method:		Get
//Description:  home page for forum
func HomePackageGet(context *gin.Context)  {
	//user := loginNeeded(context)
	user := loginNoNeed(context)
	context.HTML(http.StatusOK,"index.html",gin.H{
		"forums":control.GetForums(),
		"user":user,
	})
}

//FunctionName: GetBlockPage
//Path:			/block/:blockID
//Method:		Get
//Description:  get block page by block id
func GetBlockPage(context *gin.Context)  {
	user := loginNoNeed(context)
	blockID,_ := strconv.ParseUint(context.Param("blockID") , 10 , 64)
	block := control.GetBlock(uint(blockID))
	themeName := context.DefaultQuery("theme","all")
	var threads []model.Thread
	if themeName=="all" {
		threads = control.GetThreadsByBlockID(uint(blockID))
	}else{
		theme := control.GetThemesThreads(uint(blockID),themeName)
		threads = theme.Threads
	}
	context.HTML(http.StatusOK,"block.html",gin.H{
		"user":user,
		"block":block,
		"theme":themeName,
		"threads":threads,
	})
}

//FunctionName: BlockPageProcess
//Path:			/block/:blockID
//Method:		POST
//Description:  process post method from block page
func BlockPageProcess(context *gin.Context)  {

	user := loginNeeded(context)
	if user == nil {
		return
	}
	//process

	blockID,_ := strconv.ParseUint(context.Param("blockID") , 10 , 64)
	action := context.PostForm("action")
	if action == "addThread" {
		theme 	:= context.PostForm("theme")
		tittle 	:= context.PostForm("tittle")
		content := context.PostForm("content")
		err := control.PostThread(user.ID,uint(blockID),theme,tittle,content)
		if err==nil {
			GetBlockPage(context)
			return
		}else{
			context.String(http.StatusBadRequest,err.Error())
		}

	} else {
		context.String(http.StatusBadRequest,"post form without action")
	}

}

//FunctionName: GetThreadPage
//Path:			/thread/:threadID
//Method:		/GET
//Description:  get thread page
func GetThreadPage(context *gin.Context)  {
    user := loginNoNeed(context)

    threadID,_ := strconv.ParseUint(context.Param("threadID"),10,64)
    thread := control.GetThread(uint(threadID))
    context.HTML(http.StatusOK,"thread.html",gin.H{
    	"user" : user,
    	"thread" : thread,
	})
}

//FunctionName: ThreadPageProcess
//Path:			/thread/:threadID
//Method:		POST
//Description:  process post request from thread page
func ThreadPageProcess(context *gin.Context)  {
    user := loginNeeded(context)
	threadID,_ := strconv.ParseUint(context.Param("threadID"),10,64)
	action := context.PostForm("action")
	if action == "addComment" {
		commentContent := context.PostForm("comment")
		control.AddComment(uint(threadID),user.ID,commentContent)
	}else if action=="delete" {
		blockID := model.GetThreadFromID(uint(threadID)).BlockID
		err := control.DeleteThread(user.ID,uint(threadID))
		if err == nil {
			context.Redirect(http.StatusMovedPermanently,"/block/"+strconv.FormatUint(uint64(blockID),10))
			return
		}
	}else {
		context.String(http.StatusBadRequest,"post form without action")
		return
	}
	thread := control.GetThread(uint(threadID))
	context.HTML(http.StatusOK,"thread.html",gin.H{
		"user" : user,
		"thread" : thread,
	})
}

//FunctionName: UserInfoPage
//Path:			/user/:userID
//Method:		GET
//Description:  get user page
func UserInfoPage(context *gin.Context)  {
	user := loginNoNeed(context)
	userID := context.Param("userID")
	host,_ := control.GetUserFromID(userID)
	threads := control.GetThreadsOfUser(userID)
	context.HTML(http.StatusOK,"user.html",gin.H{
		"user":user,
		"host":host,
		"threads":threads,
	})
}

//FunctionName: ChangeUserInfo
//Path:			/changeUserInfo
//Method:		POST
//Description:  change user info
func ChangeUserInfo(context *gin.Context)  {
    user := loginNoNeed(context)

    name := context.PostForm("userName")
    email := context.PostForm("email")

	control.ChangeInfo(user.ID,name,email)

    context.Redirect(http.StatusMovedPermanently,"/user/"+user.ID)
}

//FunctionName: Register
//Path:			/register
//Method:		POST
//Description:  Register new user
func Register(context *gin.Context)  {
    ID := context.PostForm("ID")
    name := context.PostForm("name")
    email := context.PostForm("email")
    password := context.PostForm("password")
    err := control.Register(ID,name,email,password)
	if err != nil {
		context.HTML(http.StatusOK, "login.html", gin.H{
			"info": err.Error(),
		})
	}else{
		context.HTML(http.StatusOK, "login.html", gin.H{
			"info": nil,
		})
	}
}

//FunctionName: SearchThreads
//Path:			/search
//Method:		Get
//Description:  search threads
func SearchThreads(context *gin.Context)  {
    user := loginNoNeed(context)
    query := context.Query("query")
    threads := control.SearchThreads(query)

    context.HTML(http.StatusOK,"searchThread.html",gin.H{
    	"user":user,
    	"threads":threads,
	})

}


//--------------------------------------------------------------
//--------------------sub process-------------------------------

//FunctionName: loginNeeded
//Path:			/login
//Method:		subProcess
//Description:  检查登录，没有登录跳转
func loginNeeded(context *gin.Context)(user *model.User) {
	token, err := context.Cookie("USER")
	if err == nil {
		ui,err := control.ParseToken(token)
		if err==nil{
			if ui.Role == "user" {
				if user,err = control.GetUserFromID(ui.UserID);err == nil{
					if user.ThawTime.After(time.Now()){
						context.String(http.StatusBadRequest,"ID \""+user.ID+"\" was frozen")
						return nil
					}else {
						return
					}
				}
			}
		}
	}
	//TODO : else 返回登录页面
	LoginPage(context)
	return nil
}

//FunctionName: loginNoNeed
//Path:
//Method:		subProcess
//Description:  检查登录，没有登录返回空
func loginNoNeed(context *gin.Context) (user *model.User) {
	token, err := context.Cookie("USER")
	if err == nil {
		ui,err := control.ParseToken(token)
		if err==nil{
			if ui.Role == "user" {
				user,err = control.GetUserFromID(ui.UserID)
			}
		}
	}
	return
}



//FunctionName: LoginPageGet
//Path:
//Method:		subProcess
//Description:  get login page
func LoginPage(context *gin.Context)  {
	context.HTML(http.StatusOK,"login.html",gin.H{"info":""})
}
