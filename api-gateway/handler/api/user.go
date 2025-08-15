package api

import (
	"api-gateway/basic/config"
	__ "api-gateway/basic/user-proto"
	"api-gateway/handler/request"
	"api-gateway/handler/resp"
	"api-gateway/middleware"
	"github.com/gin-gonic/gin"
	midd "github.com/zhanghanchen1014/pubilc_pkg/middleware"
	"strconv"
)

// SendSms
func SendSms(c *gin.Context) {
	//接收参数
	var req request.SendSmsReq
	if err := c.ShouldBind(&req); err != nil {
		resp.RespDataErr(c, err)
		return
	}

	//业务逻辑
	sms, err := config.UserSrv.SendSms(c, &__.SendSmsReq{
		Mobile: req.Mobile,
	})
	if err != nil {
		resp.RespServerErr(c, err)
		return
	}

	//返回参数
	resp.RespSuccess(c, "短信发送成功", sms)
}

// Register
func Register(c *gin.Context) {
	//接收参数
	var req request.RegisterReq
	if err := c.ShouldBind(&req); err != nil {
		resp.RespDataErr(c, err)
		return
	}

	//业务逻辑
	register, err := config.UserSrv.Register(c, &__.RegisterReq{
		Mobile:   req.Mobile,
		Password: req.Password,
		SmsCode:  req.SmsCode,
	})
	if err != nil {
		resp.RespServerErr(c, err)
		return
	}

	//返回参数
	resp.RespSuccess(c, "用户注册成功", register)
}

// Login
func Login(c *gin.Context) {
	//接收参数
	var req request.LoginReq
	if err := c.ShouldBind(&req); err != nil {
		resp.RespDataErr(c, err)
		return
	}

	//业务逻辑
	login, err := config.UserSrv.Login(c, &__.LoginReq{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		resp.RespServerErr(c, err)
		return
	}

	//创建token
	token, err := midd.TokenHandler(strconv.FormatInt(login.UserId, 10))
	if err != nil {
		return
	}

	//返回参数
	resp.RespSuccess(c, "用户登录成功", token)
}

// PerfectUser
func PerfectUser(c *gin.Context) {
	//接收参数
	var req request.PerfectUserReq
	if err := c.ShouldBind(&req); err != nil {
		resp.RespDataErr(c, err)
		return
	}

	//业务逻辑
	perfectUser, err := config.UserSrv.PerfectUserInfo(c, &__.PerfectUserInfoReq{
		UserId:   middleware.CheckToken(c),
		NickName: req.NickName,
		Email:    req.Email,
	})
	if err != nil {
		resp.RespServerErr(c, err)
		return
	}

	//返回参数
	resp.RespSuccess(c, "完善用户信息成功", perfectUser)
}

// GetUserList
func GetUserList(c *gin.Context) {
	//接收参数
	var req request.GetUserListReq
	if err := c.ShouldBind(&req); err != nil {
		resp.RespDataErr(c, err)
		return
	}

	//业务逻辑
	getUserList, err := config.UserSrv.GetUserList(c, &__.GetUserListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		resp.RespServerErr(c, err)
		return
	}

	//返回参数
	resp.RespSuccess(c, "获取用户列表成功", getUserList)
}

// GetUserDetail
func GetUserDetail(c *gin.Context) {
	//接收参数
	var req request.GetUserDetailReq
	if err := c.ShouldBind(&req); err != nil {
		resp.RespDataErr(c, err)
		return
	}

	//业务逻辑
	getUserDetail, err := config.UserSrv.GetUserDetail(c, &__.GetUserDetailReq{
		UserId: req.UserId,
	})
	if err != nil {
		resp.RespServerErr(c, err)
		return
	}

	//返回参数
	resp.RespSuccess(c, "获取用户详情成功", getUserDetail)
}

// DeleteUser
func DeleteUser(c *gin.Context) {
	//接收参数
	var req request.DeleteUserReq
	if err := c.ShouldBind(&req); err != nil {
		resp.RespDataErr(c, err)
		return
	}

	//业务逻辑
	deleteUser, err := config.UserSrv.DeleteUser(c, &__.DeleteUserReq{
		UserId: req.UserId,
	})
	if err != nil {
		resp.RespServerErr(c, err)
		return
	}

	//返回参数
	resp.RespSuccess(c, "用户删除成功", deleteUser)
}
