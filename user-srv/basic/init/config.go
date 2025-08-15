package init

import (
	"github.com/spf13/viper"
	"log"
	"user-srv/basic/config"
)

func InitViper() {
	viper.SetConfigFile("../../dev.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic("动态配置读取失败:" + err.Error())
	}
	if err := viper.Unmarshal(&config.AppConf); err != nil {
		panic("动态配置解析失败:" + err.Error())
	}
	log.Println("动态配置读取成功:", config.AppConf)
}
