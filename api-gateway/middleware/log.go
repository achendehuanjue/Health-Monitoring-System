package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	fmt.Println("触发日志中间件")
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		logger.Info("请求日志",
			zap.String("method", c.Request.Method),
			zap.String("host", c.Request.Host),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("duration", time.Since(start)),
		)
	}
}
