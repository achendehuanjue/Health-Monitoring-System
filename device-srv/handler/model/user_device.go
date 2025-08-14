package model

import "gorm.io/gorm"

type UserDevice struct { //用户设备关联
	gorm.Model
	UserID   uint `gorm:"type:int(11);not null;index;comment:用户ID"`
	DeviceID uint `gorm:"type:int(11);not null;index;comment:设备ID"`
}
