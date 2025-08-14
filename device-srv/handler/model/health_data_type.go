package model

import "gorm.io/gorm"

type HealthDataType struct { //数据类型
	gorm.Model
	Name        string `gorm:"type:varchar(50);not null;index;comment:数据类型名称"`
	Desc        string `gorm:"type:varchar(255);comment:数据类型描述"`
	NormalRange string `gorm:"type:varchar(100);comment:正常范围"`
	DeviceId    int64  `gorm:"type:int(11);index;comment:设备ID"`
}
