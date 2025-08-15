package init

import (
	"pex-srv/untils"
)

func init() {
	untils.InitNacos()
	InitMysql()
	InitRedis()
	InitDB()
	InitEs()
	InitEmqx()
}
