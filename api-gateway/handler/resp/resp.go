package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 通用响应结构体（用于 Swagger 文档展示）
type Response struct {
	Code int         `json:"code" example:"200"`
	Msg  string      `json:"msg" example:"操作成功"`
	Data interface{} `json:"data"`
}
type Responses struct {
	Code int         `json:"code" example:"400"`
	Msg  string      `json:"msg" example:"操作失败"`
	Data interface{} `json:"data"`
}

// 请求参数报错
func RespDataErr(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": http.StatusBadRequest,
		"msg":  "参数错误",
		"data": err.Error(),
	})
	return
}

// 服务端错误
func RespServerErr(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  "服务端错误",
		"data": err.Error(),
	})
	return
}

// 成功
func RespSuccess(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  msg,
		"data": data,
	})
	return
}
