package service

import (
	"context"
	__ "pex-srv/basic/proto"
	"pex-srv/handler/dao"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedEMQXSrvServer
}

// TemperatureEMQX  温度
func (s *Server) TemperatureEMQX(_ context.Context, in *__.TemperatureEMQXReq) (*__.TemperatureEMQXResp, error) {
	dao.TemperatureEmqxClient()
	return &__.TemperatureEMQXResp{}, nil
}

// BloodGlucoseEMQX  血糖
func (s *Server) BloodGlucoseEMQX(_ context.Context, in *__.BloodGlucoseEMQXReq) (*__.BloodGlucoseEMQXResp, error) {

	dao.BloodGlucoseEmqxClient()
	return &__.BloodGlucoseEMQXResp{}, nil
}

// BloodPressureEMQX  血压
func (s *Server) BloodPressureEMQX(_ context.Context, in *__.BloodPressureEMQXReq) (*__.BloodPressureEMQXResp, error) {
	dao.BloodPressureEmqxClient()
	return &__.BloodPressureEMQXResp{}, nil
}

// BloodOxygenEMQX  血氧
func (s *Server) BloodOxygenEMQX(_ context.Context, in *__.BloodOxygenEMQXReq) (*__.BloodOxygenEMQXResp, error) {
	dao.BloodOxygenEmqxClient()
	return &__.BloodOxygenEMQXResp{}, nil
}
