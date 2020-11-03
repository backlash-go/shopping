package service

import (
	"shopping/models"
	"shopping/resource"
)

//登陆验证

func VerifyAccountLogin(account, password string) (user models.User, err error) {
	err = resource.GetDB().Model(&models.User{}).Where("account = ?", account).First(user).Error

	return
}

//判断用户是否存在
func VerifyUser(account string) (user models.User, err error) {
	err = resource.GetDB().Where("account = ?", account).First(&user).Error
	return
}

//验证用户密码

func VerifyPassword( id uint64) (verify models.User, err error) {
	err = resource.GetDB().Model(&models.User{}).Where("id = ?", id).First(&verify).Error
	return
}

//注册用户

func RegistryUser(name, password, email, address string) (err error) {
	err = resource.GetDB().Create(&models.User{Cellphone: name, Password: password, Email: email, Address: address}).Error
	return
}
