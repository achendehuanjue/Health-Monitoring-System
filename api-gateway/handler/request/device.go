package request

type GetDeviceListReq struct {
	Page     int64 `json:"page" form:"page"`
	PageSize int64 `json:"page_size" form:"page_size"`
	UserId   int64 `json:"user_id" form:"user_id"`
}
