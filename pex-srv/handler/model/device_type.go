package model

import "gorm.io/gorm"

type DeviceType struct { //设备类型
	gorm.Model
	Name string `gorm:"type:varchar(50);not null;index;comment:设备类型名称"`
	Desc string `gorm:"type:varchar(255);not null;index;comment:设备类型描述"`
}
