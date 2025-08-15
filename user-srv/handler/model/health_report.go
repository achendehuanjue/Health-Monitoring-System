package model

import "gorm.io/gorm"

type HealthReport struct {
	gorm.Model
	UserId     int64  `gorm:"type:int(11);not null;index;comment:用户ID" json:"userId"`
	ReportType string `gorm:"type:varchar(50);not null;index;comment:报告类型" json:"reportType"`
	Content    string `gorm:"type:text;comment:报告内容" json:"content"`
	Suggestion string `gorm:"type:text;comment:建议" json:"suggestion"`
	Status     int64  `gorm:"type:int(11);not null;default:0;comment:状态 0未完成 1已完成" json:"status"`
}
