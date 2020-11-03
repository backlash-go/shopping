package models

import "time"

/*
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cellphone` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
  `account` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '0' COMMENT '密码',
  `nick_name` varchar(200) NOT NULL DEFAULT '' COMMENT '展示名称',
  `real_name` varchar(16) DEFAULT NULL COMMENT '老师真实姓名',
  `email` varchar(128) NOT NULL DEFAULT '' COMMENT '邮箱',
  `avatar_url` varchar(200) NOT NULL DEFAULT '' COMMENT '头像地址',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  AUTO_INCREMENT=0  DEFAULT CHARSET=utf8mb4;

insert into user values(1,"18273041051","backlash","xixianbin","nickback",realback,"466711901@qq.com","sdsa")
*/

type User struct {
	Id        uint64     `gorm:"column:id" form:"id" json:"id"`
	Account   string     `gorm:"column:account" form:"account" json:"account"`
	NickName  string     `gorm:"column:nick_name" form:"nick_name" json:"nick_name"`
	RealName  string     `gorm:"column:real_name" form:"real_name" json:"real_name"`
	Cellphone string     `gorm:"column:cellphone" form:"cellphone" json:"cellphone"`
	Password  string     `gorm:"column:password" form:"password" json:"password"`
	Email     string     `gorm:"column:email", form:"email",json:"email"`
	AvatarUrl string     `gorm:"column:avatar_url", form:"avatar_url",json:"avatar_url"`
	CreatedAt *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" form:"deleted_at" json:"deleted_at"`
}

//帐号密码登陆验证
func (m *User) TableName() string {
	return "user"

}
