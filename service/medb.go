package service

import (
	"shopping/models"
	"shopping/resource"
)

//判断用户是否存在
func SelectUser(cellphone string) (user models.User,err error)  {
	err = resource.GetDB().Where("cellphone = ?",cellphone).First(&user).Error
	return
}

//验证用户密码

func VerifyPassword(cellphone string,password string)(verify models.User,err error){
	err = resource.GetDB().Where("cellphone = ? and password = ?",cellphone,password).First(&verify).Error
	return
}

//注册用户

func RegistryUser(name string,password string,email string,address string) (err error)  {
	err = resource.GetDB().Create(&models.User{Cellphone:name,Password:password,Email:email,Address:address}).Error
	return
}


