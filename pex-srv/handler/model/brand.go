package model

import (
	"gorm.io/gorm"
	"log"
	"pex-srv/basic/config"
)

type Brand struct { //品牌
	gorm.Model
	Name string `gorm:"type:varchar(20);not null;index;comment:品牌名称"`
	Desc string `gorm:"type:varchar(20);not null;index;comment:品牌描述"`
	Logo string `gorm:"type:varchar(300);not null;comment:品牌logo"`
}

func (b *Brand) GetBrandDetailById() error {
	return config.DB.Where("id = ?", b.ID).First(&b).Error
}

func (b *Brand) GetBrandDetailByName() error {
	return config.DB.Where("name = ?", b.Name).Limit(1).Find(&b).Error
}

func (b *Brand) AddBrand() error {
	return config.DB.Create(&b).Error
}

func (b *Brand) DelBrand() error {
	return config.DB.Where("id = ?", b.ID).Delete(&b).Error
}

func (b *Brand) UpdateBrand() error {
	return config.DB.Where("id = ?", b.ID).Updates(&b).Error
}

func (b *Brand) GetBrandList(page, pageSize int) (BrandList []*Brand, err error) {
	err = config.DB.Where("name like ?", "%"+b.Name+"%").Scopes(Paginate(page, pageSize)).Find(&BrandList).Error
	return
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		log.Println("offset:", offset)
		return db.Offset(offset).Limit(pageSize)
	}
}
