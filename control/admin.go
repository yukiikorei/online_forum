/**
 * @Author: korei
 * @Description:
 * @File: admin.go
 * @Version: 1.0.0
 * @Date: 2020/9/10 下午4:26
 */

package control

import (
	"errors"
	"gorm.io/gorm"
	"korei/onlineForum/model"
	"time"
)

//------------------------------------------------
//1.forum and block manage function
//------------------------------------------------

func GetForums() []model.SubForum {
	return model.GetSubForums()
}

func AddSubForums(name string) error{
	ok := model.AddSubForum(&model.SubForum{Name: name})
	if !ok {
		return errors.New("forums add failure")
	}
	return nil
}

func DeleteSubForums(name string) error{
	ok := model.DeleteSubForum(&model.SubForum{Name: name})
	if !ok {
		return errors.New("forums delete failure")
	}
	return nil
}

func AddBlock(forumName string,name string,masterID string) error{
	//TODO: judge error reason：
	// 1. block exit
	// 2. forum no exit
	// 3. master no exit
	//			try to get all error type
	ok := model.AddBlock(&model.Block{
		SubForumName: 	forumName,
		Name: 			name,
		MasterID: 		masterID,
	})
	if !ok {
		return errors.New("block add failure")
	}
	return nil
}

func DeleteBlock(blockID uint) error{
	ok := model.DeleteBlock(&model.Block{
		Model:gorm.Model{ID: blockID},
	})
	if !ok {
		return errors.New("block delete failure")
	}
	return nil
}

func ChangeMasterID(blockID uint,masterID string) error {
	ok := model.ChangeMaster(blockID, masterID)
	if !ok {
		return errors.New("failure in update master : user \""+masterID+"\" no found")
	}
	return nil
}

func SearchUser(ID string,name string,IP string,email string)[]model.User{
	user := model.User{
		UserName: name,
		ID:ID,
		LastIP: IP,
		Email: email,
	}

	return model.UsersFilter(&user)
}

func SearchFrozenUser(ID string,name string,IP string,email string)[]model.User{
	user := model.User{
		UserName: name,
		ID:ID,
		LastIP: IP,
		Email: email,
	}
	return model.UsersFilterFrozenOnly(&user,true)
}

func FreezeUser(userID string,duration time.Duration) error {
	ok := model.SetFrozenTime(userID,time.Now().Add(duration))
	if ok {return nil}else{
		return errors.New("freeze user failure")
	}
}

func ThawUser(userID string) error {
	ok := model.SetFrozenTime(userID,time.Now())
	if ok {return nil}else{
		return errors.New("thaw failure")
	}
}
