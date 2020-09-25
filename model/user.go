/**
 * @Author: korei
 * @Description:
 * @File:  user.go
 * @Version: 1.0.0
 * @Date: 2020/9/10 下午4:26
 */

package model

import (
	"errors"
	"time"
)

type User struct {
	ID 			string		`gorm:"primaryKey;type:varchar(11)"`
	Email		string		`gorm:"unique;type:varchar(50)"`
	UserName 	string		`gorm:"type:varchar(30)"`
	Password	string		`gorm:"type:varchar(30)"`
	LastIP		string		`gorm:"type:varchar(16)"`

	ThawTime	time.Time
}

func CreateUser(user *User) bool {
	user.ThawTime = time.Unix(0,0)
	result := db.Create(user)
	return result.RowsAffected!= 0
}

func IfFrozen(user User) bool {
	return user.ThawTime.After(time.Now())
}

func GetUserFromID(userID string)(user *User,err error){
	temUser := User{ID:userID}
	result := db.Find(&temUser)
	if result.RowsAffected == 0 {
		err = errors.New("user not exit")
	}
	user = &temUser
	return
}

func UsersFilter(user *User)(users []User){
	return UsersFilterFrozenOnly(user,false)
}

func UsersFilterFrozenOnly(user *User,frozenOnly bool)(users []User){
	filter := db
	if user.ID != ""{
		filter = filter.Where("ID like ?","%"+user.ID+"%")
	}
	if user.Email != ""{
		filter = filter.Where("Email like ?","%"+user.Email+"%")
	}
	if user.UserName != ""{
		filter = filter.Where("User_Name like ?","%"+user.UserName+"%")
	}
	if user.LastIP != ""{
		filter = filter.Where("Last_IP like ?",user.LastIP+"%")
	}
	if frozenOnly {
		filter = filter.Where("thaw_time > now()")
	}
	filter.Find(&users)
	return
}

func SetFrozenTime(userID string,thawTime time.Time) bool{
	result := db.Where("id = ? ",userID).Model(&User{}).Update("thaw_time",thawTime)
	return !(result.RowsAffected == 0)
}

func SetUserNameAndEmail(userID string,newUserName string,newEmail string)bool {
	result := db.Where("id = ?",userID).Model(&User{}).Update("user_name",
		newUserName).Update("email",newEmail)
	return result.RowsAffected != 0
}


