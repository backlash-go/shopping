package service

import (
	"github.com/rs/xid"
	"shopping/models"
	"shopping/resource"
	"time"
)

type Session struct {
	Key       string
	Id        uint64
	Account   string
	Cellphone string
	NickName  string
	RealName  string
	AvatarUrl string
	Expire    int64
	Update    int64
	LastUsed  int64
}

func CreateSession(user models.User) (*Session, error) {
	guid := xid.New()
	expireTime := time.Now().Add(10 * time.Hour).Unix()
	session := &Session{
		Key:       guid.String(),
		Id:        user.Id,
		Account:   user.Account,
		Cellphone: user.Cellphone,
		NickName:  user.NickName,
		RealName:  user.RealName,
		AvatarUrl: user.AvatarUrl,
		Expire:    expireTime,
	}
	m := map[string]interface{}{
		"Id":        user.Id,
		"Account":   user.Account,
		"Cellphone": user.Cellphone,
		"NickName":  user.NickName,
		"RealName":  user.RealName,
		"AvatarUrl": user.AvatarUrl,
	}
	err := resource.SetHashValue(session.Key, m)
	if err != nil {
		return session, err
	}

	if err := resource.SetKeyTtl(session.Key, time.Minute*30); err != nil {
		return session, err
	}
	return session, nil
}
