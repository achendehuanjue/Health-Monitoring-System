package router

import (
	"api-gateway/handler/api"
	"github.com/gin-gonic/gin"
)

func PexRouter(v1 *gin.RouterGroup) {
	pex := v1.Group("/pex")
	{
		pex.POST("/temperature", api.PexTemperature) //体温EMQX订阅

		pex.POST("/blood_glucose", api.PexBloodGlucose) //血糖EMQX订阅

		pex.POST("/blood_pressure", api.PexBloodPressure) //血压EMQX订阅

		pex.POST("/blood_oxygen", api.PexBloodOxygen) //血氧EMQX订阅
	}
}
