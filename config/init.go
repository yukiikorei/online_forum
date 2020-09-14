/**
 * @Author: korei
 * @Description: load config and check
 * @File:  init
 * @Version: 1.0.0
 * @Date: 2020/9/14 上午11:12
 */

package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

var GlobalConfig Config

type Config struct {
	Network 	Network
	Admin		Admin
	Database	Database
	JWT			JWT
}

type Network struct {
	AllowDomain	[]string
	AllowHTTP   bool
	HTTPPort    uint
	AllowHTTPS  bool
	HTTPSPort   uint
}

type Admin struct {
	AdminID     string
	AdminName   string
	Pass        string
	Email       string
	AllowRemote bool
}

type Database struct {
	Type        string
	Host        string
	Database    string
	Port        string
	User        string
	Pass        string
}

type JWT struct {
	SignedKey  	string
}


func init()  {
	configPath := "./config"
	_,err := toml.DecodeFile(configPath,&GlobalConfig)
	if err== nil{
		log.Println(GlobalConfig)
	}else {
		log.Println(err.Error())
	}
}
