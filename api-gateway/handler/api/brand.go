package api

import (
	"api-gateway/basic/config"
	__ "api-gateway/basic/device-proto"
	"api-gateway/handler/request"
	"api-gateway/handler/resp"
	"github.com/gin-gonic/gin"
)

// GetBrandList
// @Summary      品牌列表获取
// @Description  展示品牌的数据
// @Tags         品牌模块
// @Accept       json
// @Produce      json
// @Param        data  body      request.GetBrandListReq  true  "品牌参数"
// @Success      200  {object}  resp.Response
// @Failure      400   {object}  resp.Response
// @Router       /brand/list [get]
func GetBrandList(c *gin.Context) {
	var req request.GetBrandListReq
	if err := c.ShouldBind(&req); err != nil {
		resp.RespDataErr(c, err)
		return
	}

	list, err := config.DeviceSrv.GetBrandList(c, &__.GetBrandListReq{
		Name:     req.Name,
		Page:     int64(req.Page),
		PageSize: int64(req.PageSize),
	})
	if err != nil {
		resp.RespServerErr(c, err)
		return
	}

	resp.RespSuccess(c, "品牌列表获取成功", list)
}

// AddBrand
// @Summary      品牌发布
// @Description  发布一个新品牌
// @Tags         品牌模块
// @Accept       json
// @Produce      json
// @Param        data  body      request.AddBrandReq  true  "品牌参数"
// @Success      200  {object}  resp.Response
// @Failure      400   {object}  resp.Response
// @Router       /brand/add [post]
func AddBrand(c *gin.Context) {
	var req request.AddBrandReq
	if err := c.ShouldBind(&req); err != nil {
		resp.RespDataErr(c, err)
		return
	}

	add, err := config.DeviceSrv.BrandAdd(c, &__.BrandAddReq{
		Name: req.Name,
		Desc: req.Desc,
		Logo: req.Logo,
	})
	if err != nil {
		resp.RespServerErr(c, err)
		return
	}

	resp.RespSuccess(c, "添加成功", add)
}

// DelBrand
// @Summary      品牌删除
// @Description  删除一个品牌
// @Tags         品牌模块
// @Accept       json
// @Produce      json
// @Param        data  body      request.DelBrandReq  true  "品牌参数"
// @Success      200  {object}  resp.Response
// @Failure      400   {object}  resp.Response
// @Router       /brand/del [post]
func DelBrand(c *gin.Context) {
	var req request.DelBrandReq
	if err := c.ShouldBind(&req); err != nil {
		resp.RespDataErr(c, err)
		return
	}

	del, err := config.DeviceSrv.BrandDel(c, &__.BrandDelReq{
		Id: req.Id,
	})
	if err != nil {
		resp.RespServerErr(c, err)
		return
	}
	resp.RespSuccess(c, "删除成功", del)
}

// UpdateBrand
// @Summary      品牌修改
// @Description  修改品牌数据
// @Tags         品牌模块
// @Accept       json
// @Produce      json
// @Param        data  body      request.UpdateBrandReq  true  "品牌参数"
// @Success      200  {object}  resp.Response
// @Failure      400   {object}  resp.Response
// @Router       /brand/update [post]
func UpdateBrand(c *gin.Context) {
	var req request.UpdateBrandReq
	if err := c.ShouldBind(&req); err != nil {
		resp.RespDataErr(c, err)
		return
	}

	update, err := config.DeviceSrv.BrandUpdate(c, &__.BrandUpdateReq{
		Id:   req.Id,
		Name: req.Name,
		Desc: req.Desc,
		Logo: req.Logo,
	})
	if err != nil {
		resp.RespServerErr(c, err)
		return
	}

	resp.RespSuccess(c, "修改成功", update)
}
