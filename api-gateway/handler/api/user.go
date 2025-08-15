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
// @Summary      短信发送
// @Description  发送短信验证码
// @Tags         用户模块
// @Accept       json
// @Produce      json
// @Param        data  body      request.SendSmsReq  true  "用户参数"
// @Success      200  {object}  resp.Response
// @Failure      400   {object}  resp.Responses
// @Router       /user/sendSms [post]
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
// @Summary      用户注册
// @Description  用户注册
// @Tags         用户模块
// @Accept       json
// @Produce      json
// @Param        data  body      request.RegisterReq  true  "用户参数"
// @Success      200  {object}  resp.Response
// @Failure      400   {object}  resp.Responses
// @Router       /user/register [post]
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
// @Summary      用户登录
// @Description  用户登录
// @Tags         用户模块
// @Accept       json
// @Produce      json
// @Param        data  body      request.LoginReq  true  "用户参数"
// @Success      200  {object}  resp.Response
// @Failure      400   {object}  resp.Responses
// @Router       /user/login [post]
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
// @Summary      完善用户信息
// @Description  完善用户信息
// @Tags         用户模块
// @Accept       json
// @Produce      json
// @Param        data  body      request.PerfectUserReq  true  "用户参数"
// @Success      200  {object}  resp.Response
// @Failure      400   {object}  resp.Responses
// @Router       /user/perfect [post]
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
// @Summary      用户列表展示
// @Description  用户列表展示
// @Tags         用户模块
// @Accept       json
// @Produce      json
// @Param        data  body      request.GetUserListReq  true  "用户参数"
// @Success      200  {object}  resp.Response
// @Failure      400   {object}  resp.Responses
// @Router       /user/list [get]
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
// @Summary      用户详情
// @Description  展示用户的数据
// @Tags         用户模块
// @Accept       json
// @Produce      json
// @Param        data  body      request.GetUserDetailReq  true  "用户参数"
// @Success      200  {object}  resp.Response
// @Failure      400   {object}  resp.Responses
// @Router       /user/detail [get]
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
// @Summary      删除用户
// @Description  删除用户
// @Tags         用户模块
// @Accept       json
// @Produce      json
// @Param        data  body      request.DeleteUserReq  true  "用户参数"
// @Success      200  {object}  resp.Response
// @Failure      400   {object}  resp.Responses
// @Router       /user/delete [post]
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
