package model

import "gorm.io/gorm"

type User struct { //用户
	gorm.Model
	Mobile   string `gorm:"type:varchar(11);not null;index;unique;comment:手机号"`
	Password string `gorm:"type:varchar(50);not null;comment:密码"`
	Nickname string `gorm:"type:varchar(20);index;comment:昵称"`
	Email    string `gorm:"type:varchar(50);index;unique;comment:邮箱"`
}
