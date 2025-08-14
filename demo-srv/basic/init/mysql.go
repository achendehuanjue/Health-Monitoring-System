package init

import (
	"demo-srv/basic/config"
	"demo-srv/handler/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitMysql() {
	var err error
	data := config.AppConf.Mysql
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		data.User,
		data.Password,
		data.Host,
		data.Port,
		data.Database,
	)
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("mysql init failed:" + err.Error())
	}
	log.Println("mysql init success")
	err = config.DB.AutoMigrate(&model.User{},
		&model.Device{},
		&model.HealthData{},
		&model.DeviceType{},
		&model.UserDevice{},
		&model.Role{},
		&model.UserRole{},
		&model.Brand{},
		&model.HealthReport{},
		&model.HealthDataType{},
		&model.HealthData{},
	)
	if err != nil {
		panic("mysql migrate failed:" + err.Error())
	}
	log.Println("mysql migrate success")
}
