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

	Content 	ThreadContent	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ThreadContent struct {
	ThreadID	uint
	Content 	string
}

type Comment struct{
	gorm.Model
	ThreadID 	uint
	UserID  	string
	User 		User
	Content 	string
}

func AddThread(thread *Thread) bool {
	result := db.Create(thread)
	return result.RowsAffected!=0
}

func DeleteThread(userID string,threadID uint)bool {
	result := db.Where("id = ?" ,threadID).Where("user_id = ?" ,userID).Delete(&Thread{})
	return result.RowsAffected!=0
}

func DeleteThreadAsMaster(threadID uint)bool {
	result := db.Where("id = ?" ,threadID).Delete(&Thread{})
	return result.RowsAffected!=0
}

func GetThreadFromID(threadID uint)*Thread{
	thread := Thread{}
	db.Preload("Content").Preload("Comments").Where("id = ?",threadID).Find(&thread)
	return &thread
}

func GetThreadsOfUser(userID string)(threads []Thread){
	db.Where("user_id = ?",userID).Find(&threads)
	return
}

func AddComment(comment *Comment) bool {
	result := db.Create(comment)
	return result.RowsAffected!=0
}

func SearchThreads(query string)(threads []Thread){
	db.Where("tittle like ?","%"+query+"%").Find(&threads)
	return
}
