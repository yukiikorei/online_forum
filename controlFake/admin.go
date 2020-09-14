/**
 * @Author: korei
 * @Description:
 * @File:  admin.go
 * @Version: 1.0.0
 * @Date: 2020/9/11 上午9:18
 */

package controlFake

func CheckAdminLogin(userCookie string) bool {
	if userCookie == "1234567890" {
		return true
	}else {
		return false
	}
}

func AdminLogin(userID string,password string)(cookie string,result bool){

	if userID=="ADMIN" && password == "ADMIN" {
		return "1234567890",true
	}
	return "",false
}



