package model

import "gorm.io/gorm"

type HealthData struct { // 健康数据
	gorm.Model
	UserID     int64   `gorm:"type:int(11);not null;index;comment:用户ID"`
	DeviceId   int64   `gorm:"type:int(11);not null;index;comment:设备ID(手动上报时可为空)"`
	DataTypeId int64   `gorm:"type:int(11);not null;index;comment:数据类型ID"`
	DataValue  float32 `gorm:"type:float;not null;comment:数据值"`
	IsManual   int64   `gorm:"type:tinyint(1);not null;default:0;comment:是否手动上报"`
}
