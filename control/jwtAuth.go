/**
 * @Author: korei
 * @Description: jwt auth API
 * @File:  jwtAuth
 * @Version: 1.0.0
 * @Date: 2020/9/12 上午12:01
 */

package control

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"korei/onlineForum/model"
	"time"
)


var signedKey = "mySimpleKey"
var adminID	  = "admin"

//-------------------------------------------------------------
//---------------Method for login check------------------------
func Login(userID string,password string) (token string,err error){
	user,err := model.GetUserFromID(userID)
	if err!=nil {
		return
	}
	if user.Password != password {
		err = errors.New("password error")
		return
	}
	ui := UserInfo{UserID: userID,Role: "user"}
	token,err = createToken(&ui)
	return
}

func LoginAsAdmin(userID string,password string) (token string,err error){
	user,err := model.GetUserFromID(userID)
	if err!=nil {
		return
	}
	if user.Password != password {
		err = errors.New("password error")
		return
	}
	if userID != adminID {
		err = errors.New("no permission")
		return
	}
	ui := UserInfo{UserID: userID,Role: "admin"}
	token,err = createToken(&ui)
	return
}

//TODO:Login AS BLOCK MASTER
func LoginAsBlockMaster(){}


//--------------------------------------------------------------
//---------------Method for jwt---------------------------------


type UserInfo struct {
	UserID 		string
	Role 		string
	BlockID		uint		//used in block master
}

func createToken(user *UserInfo)(tokenss string,err error){
	//自定义claim
	claim := jwt.MapClaims{
		"UserID":   user.UserID,
		"Role": 	user.Role,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	tokenss,err  = token.SignedString([]byte(signedKey))
	return
}

func secret()jwt.Keyfunc{
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(signedKey),nil
	}
}

func ParseToken(tokenss string)(user *UserInfo,err error){
	user = &UserInfo{}
	token,err := jwt.Parse(tokenss,secret())
	if err != nil{
		return
	}
	claim,ok := token.Claims.(jwt.MapClaims)
	if !ok{
		err = errors.New("token format error: " + tokenss)
		return
	}
	if  !token.Valid{
		err = errors.New("token is invalid")
		return
	}

	user.UserID =	claim["UserID"].(string)
	user.Role   = 	claim["Role"].(string)
	return
}

