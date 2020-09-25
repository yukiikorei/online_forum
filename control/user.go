/**
 * @Author: korei
 * @Description: 用户逻辑处理
 * @File:  user.go
 * @Version: 1.0.0
 * @Date: 2020/9/18 上午9:03
 */

package control

import (
	"errors"
	"korei/onlineForum/model"
)

func GetUserFromID(userID string)(user *model.User,err error){
	user,err = model.GetUserFromID(userID)
	if err != nil {
		err = errors.New("user not exit")
	}
	return
}

func GetBlock(blockID uint) *model.Block {
	return model.GetBlock(blockID)
}

func GetThreadsByBlockID(blockID uint) []model.Thread{
	return model.GetBlockThreads(blockID)
}

func GetBlockThreads()  {

}

func GetThemesThreads(blockID uint,themeName string)(theme *model.Theme){
	return model.GetThemeWithThreads(blockID,themeName)
}

func PostThread(userID string,blockID uint,theme string,tittle string,content string)(err error){
	//TODO: check if frozen in block
	threadContent := model.ThreadContent{
		Content: content,
	}
	thread := model.Thread{
		Tittle: tittle,
		UserID: userID,
		BlockID: blockID,
		ThemeName: theme,
		Content: threadContent,
	}
	ok := model.AddThread(&thread)
	if !ok {
		err = errors.New("failure at add thread")
	}
	return
}

func GetThread(threadID uint) *model.Thread {
	return model.GetThreadFromID(threadID)
}

func AddComment(threadID uint,userID string,content string)(err error) {
	comment := model.Comment{
		ThreadID: threadID,
		UserID: userID,
		Content: content,
	}
	ok := model.AddComment(&comment)
	if !ok {
		err = errors.New("failure at add comment")
	}
	return

}

func DeleteThread(userID string,threadID uint)(err error){
	ok := model.DeleteThread(userID,threadID)
	if !ok {
		err = errors.New("No permission or Thread do not exist ")
	}
	return
}

func GetThreadsOfUser(userID string)[]model.Thread{
	return model.GetThreadsOfUser(userID)
}

func Register(ID string,name string,email string,password string)(err error){
	user := model.User{
		ID: ID,
		UserName: name,
		Password: password,
		Email: email,
	}
	ok := model.CreateUser(&user)
	if !ok {
		err = errors.New("ID or Email exist ")
	}
	return
}

func SearchThreads(query string)[]model.Thread {
	return model.SearchThreads(query)
}

func ChangeInfo(userID string,newUserName string,newEmail string)(err error) {
	model.SetUserNameAndEmail(userID,newUserName,newEmail)
	return
}