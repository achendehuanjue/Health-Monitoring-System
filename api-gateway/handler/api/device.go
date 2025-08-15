package api

import (
	"api-gateway/basic/config"
	__ "api-gateway/basic/device-proto"
	"api-gateway/handler/request"
	"api-gateway/handler/resp"
	"github.com/gin-gonic/gin"
)

func GetDeviceList(c *gin.Context) {
	var req request.GetDeviceListReq
	err := c.ShouldBind(&req)
	if err != nil {
		resp.RespDataErr(c, err)
		return
	}

	// 调用服务
	list, err := config.DeviceSrv.DeviceList(c, &__.DeviceListReq{
		UserId:   req.UserId,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		resp.RespServerErr(c, err)
		return
	}
	resp.RespSuccess(c, "success", list)
}
