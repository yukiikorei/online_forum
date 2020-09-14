/**
 * @Author: korei
 * @Description: 存储一些论坛中用到的名单或者列表
 * @File:  list
 * @Version: 1.0.0
 * @Date: 2020/9/11 上午8:29
 */

package model

import (
	"gorm.io/gorm"
	"time"
)

//TODO:
//	1. 版内发文权限列表
//	2. 全局冻结列表
//	3. 屏蔽IP列表

type BlockFrozenUser struct {
	//TODO: add reference of block !!
	gorm.Model
	UserID 		string		`gorm:"unique"`
	User		User
	ThawTime	time.Time
}

type ForumFrozenUser struct {
	gorm.Model
	UserID 		string		`gorm:"unique"`
	User		User
	ThawTime	time.Time
}

type BanedIP struct {
	gorm.Model
	IP 		string		`gorm:"unique;index;type:varchar(12)"`
}

func AddBanedIP(ip *BanedIP) bool {
	result := db.Create(&ip)
	return result.RowsAffected==0
}

func DeleteBanedIP(ip *BanedIP) bool {
	result := db.Unscoped().Delete(ip)
	return result.RowsAffected==0
}


