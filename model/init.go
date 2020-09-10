/**
 * @Author: korei
 * @Description: 负责model包的初始化，负责一些包内变量的管理
 * @File:  init
 * @Version: 1.0.0
 * @Date: 2020/9/10 下午4:30
 */

package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

//some global var
var db *gorm.DB

/**
 * this function is to init and connect to the database
 */
func init()  {
	//TODO(load config from config file)
	//	include: user password host port dbname
	dsn := "root:test@tcp(127.0.0.1:3306)/forum?charset=utf8mb4&parseTime=True&loc=Local"
	temDB,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		panic(err)
	}else {
		db = temDB
	}

	checkAndCreateTable()
	//insertTestData()

	test()
}

func checkAndCreateTable()  {
	mig := db.Migrator()
	if !mig.HasTable(&User{}) {
		mig.CreateTable(&User{})
	}
	if !mig.HasTable(&Comment{}){
		mig.CreateTable(&Comment{})
	}
	if !mig.HasTable(&Block{}){
		mig.CreateTable(&Block{})
		//db.AutoMigrate(&Block{})
	}
	if !mig.HasTable(&SubForum{}){
		mig.CreateTable(&SubForum{})
		//db.AutoMigrate(&Block{},&SubForum{})
	}
//	if !mig.HasTable(&Thread{}){
//		mig.CreateTable(&Thread{})
//	}


}

func insertTestData() {
	yukii := User{
		ID:"yukii",
		Email: "yukii@gmail.com",
		UserName: "yukii",
		Password: "test",
	}
	korei := User{
		ID:"korei",
		Email: "korei@gmail.com",
		UserName: "korei",
		Password: "test",
	}
	sb1 := SubForum{
		Name: "sb1",
	}
	b1 := Block{
		SubForumName: "sb1",
		Name: "b1",
		MasterID: "yukii",
	}
	b2 := Block{
		SubForumName: "sb1",
		Name: "b2",
		MasterID: "korei",
	}

	db.Create(&yukii)
	db.Create(&korei)
	db.Create(&sb1)
	db.Create(&b1)
	db.Create(&b2)
}

func restoreDB()  {
	mig := db.Migrator()
	mig.DropTable(&User{})
	mig.DropTable(&Comment{})
	mig.DropTable(&Block{})
	mig.DropTable(&Thread{})
}

func test()  {
	var user 	User
	var sf		SubForum

	db.Find(&user)
	db.Preload("Blocks").Find(&sf)

	log.Println(user)
	log.Println(sf.Blocks[0].Name)
	log.Println(sf.Blocks[1].Name)

}