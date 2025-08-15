package untils

import (
	"gorm.io/gorm"
	"log"
)

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
