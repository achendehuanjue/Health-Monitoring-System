package init

import "user-srv/untils"

func init() {
	untils.InitNacos()
	InitMysql()
	InitRedis()
	InitDB()
}
