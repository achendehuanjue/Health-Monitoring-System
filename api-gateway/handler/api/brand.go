package api

import (
	"api-gateway/basic/config"
	__ "api-gateway/basic/device-proto"
	"api-gateway/handler/request"
	"github.com/gin-gonic/gin"
)

func GetBrandList(c *gin.Context) {
	var req request.GetBrandListReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	resp, err := config.DeviceSrv.GetBrandList(c, &__.GetBrandListReq{
		Name:     req.Name,
		Page:     int64(req.Page),
		PageSize: int64(req.PageSize),
	})
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": resp,
	})
}

func AddBrand(c *gin.Context) {
	var req request.AddBrandReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	resp, err := config.DeviceSrv.BrandAdd(c, &__.BrandAddReq{
		Name: req.Name,
		Desc: req.Desc,
		Logo: req.Logo,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": resp,
	})
}

func DelBrand(c *gin.Context) {
	var req request.DelBrandReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	resp, err := config.DeviceSrv.BrandDel(c, &__.BrandDelReq{
		Id: req.Id,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": resp,
	})
}

func UpdateBrand(c *gin.Context) {
	var req request.UpdateBrandReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	resp, err := config.DeviceSrv.BrandUpdate(c, &__.BrandUpdateReq{
		Id:   req.Id,
		Name: req.Name,
		Desc: req.Desc,
		Logo: req.Logo,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": resp,
	})
}
