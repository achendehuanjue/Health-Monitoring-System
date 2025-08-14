package request

type GetBrandListReq struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
	Name     string `json:"name" form:"name"`
}

type AddBrandReq struct {
	Name string `json:"name" form:"name" binding:"required"`
	Desc string `json:"desc" form:"desc" binding:"required"`
	Logo string `json:"logo" form:"logo" binding:"required"`
}

type DelBrandReq struct {
	Id int64 `json:"id" form:"id" binding:"required"`
}
type UpdateBrandReq struct {
	Id   int64  `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name"`
	Desc string `json:"desc" form:"desc"`
	Logo string `json:"logo" form:"logo"`
}
