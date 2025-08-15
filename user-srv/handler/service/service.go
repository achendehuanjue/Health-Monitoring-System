package service

import (
	"context"
	"errors"
	"github.com/zhanghanchen1014/pubilc_pkg/pkg"
	"github.com/zhanghanchen1014/pubilc_pkg/untils"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"time"
	"user-srv/basic/config"
	__ "user-srv/basic/proto"
	"user-srv/handler/model"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedUserServer
}

// SendSms
func (s *Server) SendSms(_ context.Context, in *__.SendSmsReq) (*__.SendSmsResp, error) {

	sms := untils.ALiYun{}

	//定义发送短信次数
	count := config.Rdb.Get(config.Ctx, "sendSmsErr"+in.Mobile).Val()
	if count >= "3" {
		return nil, errors.New("短信发送次数频繁，1分钟值允许发送3次")
	}

	code := rand.Intn(9000) + 1000

	err := sms.ALiYunSms(in.Mobile, strconv.Itoa(code))
	if err != nil {
		return nil, errors.New("短信发送失败")
	}

	//存入缓存
	config.Rdb.Set(config.Ctx, "sendSms"+in.Mobile, code, time.Minute*30)
	//发送次数自增
	config.Rdb.Incr(config.Ctx, "sendSmsErr"+in.Mobile)
	//锁住发送时间
	config.Rdb.Expire(config.Ctx, "sendSmsErr"+in.Mobile, time.Hour*1)

	return &__.SendSmsResp{}, nil
}

// Register
func (s *Server) Register(_ context.Context, in *__.RegisterReq) (*__.RegisterResp, error) {
	var err error

	//短信验证
	get := config.Rdb.Get(config.Ctx, "sendSms"+in.Mobile).Val()
	if get != in.SmsCode {
		return nil, errors.New("短信验证失败")
	}

	var user model.User
	if err = user.GetUserByMobile(in.Mobile); err != nil {
		return nil, errors.New("用户查询失败")
	}

	//注册
	user = model.User{
		Mobile:   in.Mobile,
		Password: pkg.Md5(in.Password),
	}
	if err = user.CreateUser(); err != nil {
		return nil, errors.New("用户注册失败")
	}

	return &__.RegisterResp{
		UserId: int64(user.ID),
	}, nil
}

// Login
func (s *Server) Login(_ context.Context, in *__.LoginReq) (*__.LoginResp, error) {
	var err error

	var user model.User
	if err = user.GetUserByMobile(in.Mobile); err != nil {
		return nil, errors.New("用户查询失败")
	}

	if user.ID == 0 {
		return nil, errors.New("该用户不存在")
	}

	if user.Password != pkg.Md5(in.Password) {
		return nil, errors.New("账号密码错误")
	}

	return &__.LoginResp{
		UserId: int64(user.ID),
	}, nil
}

// PerfectUserInfo
func (s *Server) PerfectUserInfo(_ context.Context, in *__.PerfectUserInfoReq) (*__.PerfectUserInfoResp, error) {
	var err error

	var user model.User
	if err = user.GetUserById(int(in.UserId)); err != nil {
		return nil, errors.New("用户查询失败")
	}

	user = model.User{
		Model:    gorm.Model{ID: uint(in.UserId)},
		Nickname: in.NickName,
		Email:    in.Email,
	}
	if err = user.PerfectUserInfo(int(in.UserId)); err != nil {
		return nil, errors.New("完善用户信息失败")
	}

	return &__.PerfectUserInfoResp{
		UserId: int64(user.ID),
	}, nil
}

// GetUserList
func (s *Server) GetUserList(_ context.Context, in *__.GetUserListReq) (*__.GetUserListResp, error) {

	var user model.User
	list, err := user.GetUserList(int(in.Page), int(in.PageSize))
	if err != nil {
		return nil, errors.New("用户查询失败")
	}

	var userList []*__.GetUserInfo
	for _, u := range list {
		userList = append(userList, &__.GetUserInfo{
			UserId:   int64(u.ID),
			Mobile:   u.Mobile,
			Password: "***",
			NickName: u.Nickname,
			Email:    u.Email,
		})
	}

	return &__.GetUserListResp{
		UserList: userList,
	}, nil
}

// GetUserDetail
func (s *Server) GetUserDetail(_ context.Context, in *__.GetUserDetailReq) (*__.GetUserDetailResp, error) {
	var err error

	var user model.User
	if err = user.GetUserDetail(int(in.UserId)); err != nil {
		return nil, errors.New("用户查询失败")
	}

	if user.ID == 0 {
		return nil, errors.New("该用户不存在")
	}

	return &__.GetUserDetailResp{
		UserId:   int64(user.ID),
		Mobile:   user.Mobile,
		Password: "****",
		NickName: user.Nickname,
		Email:    user.Email,
	}, nil
}

func (s *Server) DeleteUser(_ context.Context, in *__.DeleteUserReq) (*__.DeleteUserResp, error) {
	var err error

	var user model.User
	if err = user.GetUserById(int(in.UserId)); err != nil {
		return nil, errors.New("用户查询失败")
	}

	if user.ID == 0 {
		return nil, errors.New("该用户不存在")
	}

	if err = user.DeleteUser(int(in.UserId)); err != nil {
		return nil, errors.New("用户删除失败")
	}

	return &__.DeleteUserResp{}, nil
}
