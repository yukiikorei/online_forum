/**
 * @Author: korei
 * @Description: 响应版主的请求
 * @File:  master
 * @Version: 1.0.0
 * @Date: 2020/9/18 下午3:21
 */

package requestHandler

import (
	"github.com/gin-gonic/gin"
	"korei/onlineForum/control"
	"korei/onlineForum/model"
	"net/http"
	"strconv"
)

//FunctionName: GetMasterPage
//Path:			/master/:blockID
//Method:		GET
//Description:  Get page from master
func GetMasterPage(context *gin.Context)  {
	ok,blockID,user := MasterCheck(context)
	if !ok {return}
	block := control.GetBlock(blockID)
	context.HTML(http.StatusOK,"blockMaster.html",gin.H{
		"block":block,
		"user":user,
	})
}

//FunctionName: ProcessMaster
//Path:			/master/:blockID
//Method:		POST
//Description:  process Post request
func ProcessMaster(context *gin.Context)  {
	ok,blockID,_ := MasterCheck(context)
	if !ok {return}
	action := context.PostForm("action")
	if action=="addTheme" {
		themeName := context.PostForm("themeName")
		err := control.AddTheme(blockID,themeName)
		if err!= nil {
			context.String(http.StatusBadRequest,err.Error())
		}else{
			GetMasterPage(context)
			return
		}
	}else if action=="deleteTheme" {
		themeName := context.PostForm("themeName")
		err := control.DeleteTheme(blockID,themeName)
		if err!= nil {
			context.String(http.StatusBadRequest,err.Error())
		}else{
			GetMasterPage(context)
			return
		}
	}else {
		context.String(http.StatusBadRequest,"post form without action")
		return
	}
}

//---------------------------------------------------------------------------------
//----------------------sub func---------------------------------------------------

//FunctionName: MasterCheck
//Path:			sub func
//Method:		sub func
//Description:  check if master or admin
func MasterCheck(context *gin.Context) (ifMaster bool,blockID uint,user *model.User) {
	user = loginNeeded(context)
	blockIDOrigin,_ := strconv.ParseUint(context.Param("blockID"),10,64)
	if !control.IsMasterOrAdmin(user.ID,uint(blockIDOrigin)) {
		context.String(http.StatusBadRequest,"no permission")
		return false,0,nil
	}
	return true,uint(blockIDOrigin),user
}