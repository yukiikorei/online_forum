/**
 * @Author: korei
 * @Description:
 * @File: forum.go
 * @Version: 1.0.0
 * @Date: 2020/9/10 下午4:26
 */

package model

import (
	"gorm.io/gorm"
)

type SubForum struct {
	gorm.Model
	Name	string		`gorm:"unique;type:varchar(40)"`
	Blocks 	[]Block		`gorm:"foreignKey:SubForumName;references:Name"`
}

type Block struct {
	gorm.Model
	SubForumName	string
	Name 			string	`gorm:"unique;type:varchar(40)"`
	MasterID		string
	Master 			User
}
