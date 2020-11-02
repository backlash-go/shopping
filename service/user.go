package service

import "shopping/models"

type Session struct {
	id        uint64
	Account   string
	Cellphone string
	NickName  string
	RealName  string
	AvatarUrl string
}


func CreateSession (user models.User) {

}
