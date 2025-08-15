package init

import "device-srv/untils"

func init() {
	untils.InitNacos()
	InitMysql()
	InitRedis()
	InitDB()
	InitEs()
}
