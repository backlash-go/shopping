package models

import "time"
/*
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cellphone` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '0' COMMENT '密码',
  `email` varchar(128) NOT NULL DEFAULT '' COMMENT '邮箱',
  `address` varchar(128) NOT NULL DEFAULT '' COMMENT '地址',
 // `active` tinyint(1) NOT NULL DEFAULT '' COMMENT '是否激活:0不是，1是',
 //`power` tinyint(1) NOT NULL DEFAULT '' COMMENT '权限设置  0 表示普通用户  1表示管理员用户',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
*/

type User struct {
	Id uint  `gorm:"column:id" form:"id" json:"id"`
	Cellphone string  `gorm:"column:cellphone" form:"cellphone" json:"cellphone"`
	Password string `gorm:"column:password" form:"password" json:"password"`
	Email string	`gorm:"column:email", form:"email",json:"email"`
	Address string  `gorm:"column:address", form:"address",json:"address"`
	//Active uint     `gorm:"column:active", form:"active",json:"active"`
	//Power uint      `gorm:"column:power", form:"power",json:"power"`
	Created_at *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	Updated_at *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	Deleted uint `gorm:"column:deleted" form:"deleted" json:"deleted"`
}


//帐号密码登陆验证
func (m *User) TableName() string  {
	return "user"

}

