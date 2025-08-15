package router

import (
	"api-gateway/handler/api"
	"github.com/gin-gonic/gin"
)

func BrandRouter(v1 *gin.RouterGroup) {
	brand := v1.Group("/brand")
	{
		brand.POST("/add", api.AddBrand)
		brand.POST("/del", api.DelBrand)
		brand.POST("/update", api.UpdateBrand)
		brand.GET("/list", api.GetBrandList)
	}
	device := v1.Group("/device")
	{
		device.GET("/list", api.GetDeviceList)

	}
}
