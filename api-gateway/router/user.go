package router

import (
	"api-gateway/handler/api"
	"github.com/gin-gonic/gin"
)

func UserRouter(v1 *gin.RouterGroup) {
	//用户模块
	user := v1.Group("/user")
	{
		user.POST("/sendSms", api.SendSms)     //SendSms
		user.POST("/register", api.Register)   //Register
		user.POST("/login", api.Login)         //Login
		user.POST("/perfect", api.PerfectUser) //PerfectUser
		user.GET("/list", api.GetUserList)     //GetUserList
		user.GET("/detail", api.GetUserDetail) //GetUserDetail
		user.POST("/delete", api.DeleteUser)   //DeleteUser
	}
}
