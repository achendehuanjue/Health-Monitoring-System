package model

import (
	"gorm.io/gorm"
	"time"
)

type Device struct { //设备
	gorm.Model
	Name           string    `gorm:"type:varchar(50);not null;index;comment:设备名称"`
	Desc           string    `gorm:"type:varchar(255);not null;index;comment:设备描述"`
	Status         int64     `gorm:"type:tinyint(1);not null;index;comment:设备状态 1离线 2在线"`
	LastOnlineTime time.Time `gorm:"type:datetime;index;comment:最后在线时间"`
}
