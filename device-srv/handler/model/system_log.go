package model

import "time"

// 系统日志表
type SystemLog struct {
	ID         int64     `gorm:"primaryKey;autoIncrement;comment:日志ID，主键"`
	LogLevel   string    `gorm:"type:varchar(20);not null;index:idx_level;comment:日志级别（如：INFO, WARN, ERROR）"`
	Module     string    `gorm:"type:varchar(50);not null;index:idx_module;comment:模块名称"`
	Message    string    `gorm:"type:text;not null;comment:日志消息"`
	StackTrace string    `gorm:"type:mediumtext;comment:异常堆栈信息"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP;index:idx_created_at;comment:创建时间"`
}
