package model

import (
	"device-srv/basic/config"
	"device-srv/untils"
	"gorm.io/gorm"
)

type UserDevice struct { //用户设备关联
	gorm.Model
	UserID   int64 `gorm:"type:int(11);not null;index;comment:用户ID"`
	DeviceID int64 `gorm:"type:int(11);not null;index;comment:设备ID"`
}

type UserDeviceInfo struct {
	Id         int64  `json:"id"`
	UserID     int64  `json:"user_id"`
	DeviceID   int64  `json:"device_id"`
	Mobile     string `json:"mobile"`
	NickName   string `json:"nick_name"`
	Email      string `json:"email"`
	DeviceName string `json:"device_name"`
}

func (ud *UserDevice) GetUserDeviceInfo(page, pageSize int) (userDeviceList []*UserDeviceInfo, err error) {
	var where map[string]interface{}
	if ud.UserID != 0 {
		where["user_id"] = ud.UserID
	}
	err = config.DB.Table("user_device").
		Joins("left join user on user_device.user_id = user.id").
		Joins("left join device on user_device.device_id = device.id").
		Select("user_device.*,user.mobile,user.nick_name,user.email,device.name as device_name").
		Scopes(untils.Paginate(page, pageSize)).
		Where(where).
		Find(&userDeviceList).Error

	return
}
