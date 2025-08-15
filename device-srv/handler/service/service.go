package service

import (
	"context"
	__ "device-srv/basic/proto"
	"device-srv/handler/model"
	"errors"
	"gorm.io/gorm"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedDeviceSrvServer
}

// BrandAdd
func (s *Server) BrandAdd(_ context.Context, in *__.BrandAddReq) (*__.BrandAddResp, error) {
	var brand model.Brand
	brand.Name = in.Name
	if err := brand.GetBrandDetailByName(); err != nil {
		return nil, err
	}
	if brand.ID != 0 {
		return nil, errors.New("该品牌已存在")
	}
	brand = model.Brand{
		Name: in.Name,
		Desc: in.Desc,
		Logo: in.Logo,
	}
	if err := brand.AddBrand(); err != nil {
		return nil, err
	}
	return &__.BrandAddResp{
		Id: int64(brand.ID),
	}, nil
}

// BrandDel
func (s *Server) BrandDel(_ context.Context, in *__.BrandDelReq) (*__.BrandDelResp, error) {
	var brand model.Brand
	brand.ID = uint(in.Id)
	if err := brand.GetBrandDetailById(); err != nil {
		return nil, errors.New("该品牌不存在")
	}
	if err := brand.DelBrand(); err != nil {
		return nil, err
	}
	return &__.BrandDelResp{
		Id: int64(brand.ID),
	}, nil
}

func (s *Server) BrandUpdate(_ context.Context, in *__.BrandUpdateReq) (*__.BrandUpdateResp, error) {
	var brand model.Brand
	brand.ID = uint(in.Id)
	if err := brand.GetBrandDetailById(); err != nil {
		return nil, errors.New("该品牌不存在")
	}
	brand = model.Brand{
		Model: gorm.Model{ID: uint(in.Id)},
		Name:  in.Name,
		Desc:  in.Desc,
		Logo:  in.Logo,
	}
	if err := brand.UpdateBrand(); err != nil {
		return nil, err
	}
	return &__.BrandUpdateResp{
		Id: int64(brand.ID),
	}, nil
}

func (s *Server) GetBrandList(_ context.Context, in *__.GetBrandListReq) (*__.GetBrandListResp, error) {
	var brand model.Brand
	brand.Name = in.Name
	BrandList, err := brand.GetBrandList(int(in.Page), int(in.PageSize))
	if err != nil {
		return nil, err
	}
	var brandListResp []*__.GetBrandList
	for _, v := range BrandList {
		brandListResp = append(brandListResp, &__.GetBrandList{
			Id:   int64(v.ID),
			Name: v.Name,
			Desc: v.Desc,
			Logo: v.Logo,
		})
	}
	return &__.GetBrandListResp{
		List:     brandListResp,
		Page:     in.Page,
		PageSize: in.PageSize,
		Total:    0,
	}, nil
}

func (s *Server) DeviceList(_ context.Context, in *__.DeviceListReq) (*__.DeviceListResp, error) {
	var userDevice model.UserDevice
	userDevice.UserID = in.UserId
	brandList, err := userDevice.GetUserDeviceInfo(int(in.Page), int(in.PageSize))
	if err != nil {
		return nil, err
	}
	var list []*__.DeviceList
	for _, info := range brandList {
		list = append(list, &__.DeviceList{
			Id:         info.Id,
			NickName:   info.NickName,
			DeviceName: info.DeviceName,
			Mobile:     info.Mobile,
			Email:      info.Email,
		})
	}
	return &__.DeviceListResp{
		List:     list,
		Page:     in.Page,
		PageSize: in.PageSize,
	}, nil
}
