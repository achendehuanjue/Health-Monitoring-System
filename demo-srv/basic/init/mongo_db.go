package init

import (
	"context"
	"demo-srv/basic/config"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() {
	// 设置客户端连接配置
	cfg := config.AppConf.MongoDb
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.User, cfg.Password, cfg.Host, cfg.Port)
	clientOptions := options.Client().ApplyURI(uri)

	// 连接到MongoDB
	var err error
	config.Mdb, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = config.Mdb.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB 连接成功")

	config.Mdb.Database("health_db").Collection("user")
}
