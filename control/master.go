/**
 * @Author: korei
 * @Description: 版主逻辑处理
 * @File:  master
 * @Version: 1.0.0
 * @Date: 2020/9/24 下午8:14
 */

package control

import (
	"errors"
	"korei/onlineForum/config"
	"korei/onlineForum/model"
)

func IsMasterOrAdmin(userID string,blockID uint) bool {
	if userID == config.GlobalConfig.Admin.AdminID{
		return true
	}
	block := GetBlock(blockID)
	if userID == block.MasterID {
		return true
	}
	return false
}

func AddTheme(blockID uint,themeName string) (err error) {
	ok := model.AddTheme(blockID,themeName)
	if !ok {
		err = errors.New("failure at add theme")
	}
	return
}

func DeleteTheme(blockID uint,themeName string) (err error) {
	ok := model.DeleteTheme(blockID,themeName)
	if !ok {
		err = errors.New("failure at delete theme")
	}
	return
}

func DeleteThreadAsMaster(threadID uint)(err error)  {
	ok := model.DeleteThreadAsMaster(threadID)
	if !ok {
		err = errors.New("thread not exist ")
	}
	return
}
