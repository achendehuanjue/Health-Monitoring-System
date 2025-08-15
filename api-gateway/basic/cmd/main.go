package main

import (
	_ "api-gateway/basic/init"
	"api-gateway/router"
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//@title 健康检测系统
//@version 1.0
//@description This is an API Gateway service
//@termsOfService http://swagger.io/terms/

//@contact.name API Support
//@contact.url http://www.swagger.io/support
//@contact.email support@swagger.io

//@license.name Apache 2.0
//@license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /v1
func main() {
	r := gin.Default()

	router.LoadRouters(r)
	// 添加 Swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
