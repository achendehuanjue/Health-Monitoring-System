package init

import (
	"demo-srv/untils"
)

func init() {
	untils.InitNacos()
	InitMysql()
	InitRedis()
	InitDB()
	InitEs()
}
