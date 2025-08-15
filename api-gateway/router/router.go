package router

import (
	"api-gateway/middleware"
	"api-gateway/untils"
	"github.com/gin-gonic/gin"
)

func LoadRouters(r *gin.Engine) {
	// 加载日志中间件
	logger, _ := untils.SetUpLogger()
	r.Use(middleware.LoggerMiddleware(logger))
	v1 := r.Group("/v1")
	{
		BrandRouter(v1)
		PexRouter(v1)
		UserRouter(v1)
	}
}
