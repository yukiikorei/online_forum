/**
 * @Author: korei
 * @Description:
 * @File:  thread.go
 * @Version: 1.0.0
 * @Date: 2020/9/10 下午4:26
 */

package model

import (
	"gorm.io/gorm"
)

/*
 *thread 帖子类，帖子的实际内容根据帖子的id存放在帖子当中
 */
type Thread struct {
	gorm.Model
	Tittle		string
	Block 		Block
	Theme		string
	UserID 		string
	User 		User			`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Comments 	[]Comment
}

type Comment struct{
	gorm.Model
	UserID  	string
	User 		User
	Content 	string
}

func addThread()  {

}

