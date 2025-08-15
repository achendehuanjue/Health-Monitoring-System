package model

import (
	"gorm.io/gorm"
	"user-srv/basic/config"
)

type User struct { //用户
	gorm.Model
	Mobile   string `gorm:"type:varchar(11);not null;index;unique;comment:手机号"`
	Password string `gorm:"type:varchar(50);not null;comment:密码"`
	Nickname string `gorm:"type:varchar(20);index;comment:昵称"`
	Email    string `gorm:"type:varchar(50);index;unique;comment:邮箱"`
}

// 注册
func (u *User) CreateUser() error {
	return config.DB.Create(&u).Error
}

// 根据id查询
func (u *User) GetUserById(id int) error {
	return config.DB.Model(&User{}).Where("id = ?", id).Limit(1).Find(&u).Error
}

// 根据mobile
func (u *User) GetUserByMobile(mobile string) error {
	return config.DB.Model(&User{}).Where("mobile = ?", mobile).Limit(1).Find(&u).Error
}

// 完善用户信息
func (u *User) PerfectUserInfo(id int) error {
	return config.DB.Model(&User{}).Where("id = ?", id).Updates(&u).Error
}

// 用户列表
func (u *User) GetUserList(page, pageSize int) (userList []*User, err error) {
	db := config.DB.Model(&User{})
	offset := (page - 1) * pageSize
	if page > 0 || pageSize > 0 {
		db = db.Offset(offset).Limit(pageSize)
	}
	db.Find(&userList)
	return userList, err
}

// 用户详情
func (u *User) GetUserDetail(id int) error {
	return config.DB.Model(&User{}).Where("id = ?", id).Limit(1).Find(&u).Error
}

// 用户删除
func (u *User) DeleteUser(id int) error {
	return config.DB.Model(&User{}).Where("id = ?", id).Delete(&u).Find(&u).Error
}
