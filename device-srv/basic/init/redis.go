package init

import (
	"demo-srv/basic/config"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

func InitRedis() {
	addr := fmt.Sprintf("%s:%d", config.AppConf.Redis.Host, config.AppConf.Redis.Port)
	config.Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,                          // use default Addr
		Password: config.AppConf.Redis.Password, // no password set
		DB:       config.AppConf.Redis.Database, // use default DB
	})

	_, err := config.Rdb.Ping(config.Ctx).Result()
	if err != nil {
		panic("redis init failed:" + err.Error())
	}
	log.Println("redis init success")
}
