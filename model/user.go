/**
 * @Author: korei
 * @Description:
 * @File:  user.go
 * @Version: 1.0.0
 * @Date: 2020/9/10 下午4:26
 */

package model

type User struct {
	ID 			string		`gorm:"primaryKey;type:varchar(11)"`
	Email		string		`gorm:"unique;type:varchar(50)"`
	UserName 	string		`gorm:"type:varchar(30)"`
	Password	string		`gorm:"type:varchar(30)"`
}

func createUser(user *User){
	db.Create(user)
}



