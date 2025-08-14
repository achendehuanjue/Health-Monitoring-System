package model

import "gorm.io/gorm"

type UserRole struct { //用户角色关联
	gorm.Model
	UserID uint `gorm:"type:int(11);not null;index;comment:用户ID"`
	RoleID uint `gorm:"type:int(11);not null;index;comment:角色ID"`
}
