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
	SubForumName	string	`gorm:"type:varchar(40)"`
	Name 			string	`gorm:"unique;type:varchar(40)"`
	MasterID		string
	Master 			User
	Themes			[]Theme
}

//TODO: Block 置顶

//TODO: Block 主题
type Theme struct {
	BlockID 	uint 		`gorm:"primaryKey"`
	BlockName	string		`gorm:"type:varchar(40)"`
	Name 		string		`gorm:"primaryKey;type:varchar(30)"`
	Threads		[]Thread	`gorm:"foreignKey:BlockID,ThemeName;references:BlockID,Name"`
}

//----------------------------------------------------------
//-----------Access Method----------------------------------
//----------------------------------------------------------


func AddSubForum(subForum *SubForum) bool{
	result := db.Create(subForum)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func DeleteSubForum(subForum *SubForum) bool{
	result := db.Where("name = ?",subForum.Name).Unscoped().Delete(subForum)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func GetSubForums()(subForums []SubForum){
	db.Preload("Blocks").Find(&subForums)
	return
}

func AddBlock(newblock *Block) bool {
	//block name ,sub forum name ,master needed
	result := db.Create(newblock)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func DeleteBlock(block *Block) bool{
	result := db.Unscoped().Delete(block)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func ChangeMaster(blockID uint,newMasterID string) bool {
	result := db.Model(&Block{}).Where("id = ?",blockID).Update("master_id",newMasterID)
	return  !(result.RowsAffected == 0)
}
