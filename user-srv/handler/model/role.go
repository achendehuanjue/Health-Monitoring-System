package model

import "gorm.io/gorm"

type Role struct { //角色
	gorm.Model
	Name string `gorm:"type:varchar(50);not null;index;comment:角色名称"`
	Desc string `gorm:"type:varchar(255);not null;index;comment:角色描述"`
}
