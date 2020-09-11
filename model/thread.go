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
	Tittle		string			`gorm:"type:varchar(100)"`

	BlockID 	uint
	ThemeName	string			`gorm:"type:varchar(30)"`

	UserID 		string			`gorm:"type:varchar(11)"`
	User 		User			`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Comments 	[]Comment		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Comment struct{
	gorm.Model
	ThreadID 	uint
	UserID  	string
	User 		User
	Content 	string
}

func addThread()  {

}

