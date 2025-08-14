package config

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	AppConf AppConfig
	DB      *gorm.DB
	Ctx     = context.Background()
	Rdb     *redis.Client
	Mdb     *mongo.Client
)
