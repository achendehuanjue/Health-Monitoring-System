package request

// SendSms
type SendSmsReq struct {
	Mobile string `json:"mobile" form:"mobile" binding:"required"`
}

// Register
type RegisterReq struct {
	Mobile   string `json:"mobile" form:"mobile" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	SmsCode  string `json:"sms_code" json:"sms_code" binding:"required"`
}

// Login
type LoginReq struct {
	Mobile   string `json:"mobile" form:"mobile" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// PerfectUser
type PerfectUserReq struct {
	NickName string `json:"nick_name" form:"nick_name"`
	Email    string `json:"email" form:"email"`
}

// GetUserList
type GetUserListReq struct {
	Page     int64 `json:"page" form:"page"`
	PageSize int64 `json:"page_size" form:"page_size"`
}

// GetUserDetail
type GetUserDetailReq struct {
	UserId int64 `json:"user_id" form:"user_id" binding:"required"`
}

// DeleteUser
type DeleteUserReq struct {
	UserId int64 `json:"user_id" form:"user_id" binding:"required"`
}
