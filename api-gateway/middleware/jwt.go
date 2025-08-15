package middleware

import (
	"api-gateway/handler/resp"
	"errors"
	"github.com/gin-gonic/gin"
	middleware "github.com/zhanghanchen1014/pubilc_pkg/middleware"
	"strconv"
)

// 获取token
func CheckToken(c *gin.Context) int64 {
	r := c.Request

	token := r.Header.Get("token")

	if token == "" {
		resp.RespDataErr(c, errors.New("token不能为空"))
		return 0
	}

	getToken, s := middleware.GetToken(token)

	if s != "" {
		resp.RespDataErr(c, errors.New(s))
		return 0
	}

	r.ParseForm()
	r.Form.Set("userId", getToken["userId"].(string))

	c.Next()
	
	userId, _ := strconv.ParseInt(getToken["userId"].(string), 10, 64)
	return userId

}
