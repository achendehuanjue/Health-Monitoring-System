package api

import (
	"api-gateway/basic/config"
	pex "api-gateway/basic/pex-proto"
	"api-gateway/handler/resp"
	"github.com/gin-gonic/gin"
)

func PexTemperature(c *gin.Context) {
	_, err := config.PexSrv.TemperatureEMQX(c, &pex.TemperatureEMQXReq{})
	if err != nil {
		resp.RespServerErr(c, err)
		return
	}

	resp.RespSuccess(c, "success", nil)
}

func PexBloodGlucose(c *gin.Context) {
	_, err := config.PexSrv.BloodGlucoseEMQX(c, &pex.BloodGlucoseEMQXReq{})
	if err != nil {
		resp.RespServerErr(c, err)
		return
	}

	resp.RespSuccess(c, "success", nil)
}

func PexBloodPressure(c *gin.Context) {
	_, err := config.PexSrv.BloodPressureEMQX(c, &pex.BloodPressureEMQXReq{})
	if err != nil {
		resp.RespServerErr(c, err)
		return
	}

	resp.RespSuccess(c, "success", nil)
}

func PexBloodOxygen(c *gin.Context) {
	_, err := config.PexSrv.BloodOxygenEMQX(c, &pex.BloodOxygenEMQXReq{})
	if err != nil {
		resp.RespServerErr(c, err)
		return
	}

	resp.RespSuccess(c, "success", nil)
}
