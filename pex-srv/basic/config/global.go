package config

import (
	"context"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	Es      *elasticsearch.Client
	AppConf AppConfig
	DB      *gorm.DB
	Ctx     = context.Background()
	Rdb     *redis.Client
	Mdb     *mongo.Client
)
var Opts *mqtt.ClientOptions
