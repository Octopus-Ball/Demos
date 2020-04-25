package model

import (
	"github.com/jinzhu/gorm"
	"Cron/init"
)

// Job 任务模型
type Job struct {
	gorm.Model
	Name string `json:"name,omitempty" db:"name" gorm:"name"`
	Cron string `json:"cron,omitempty" db:"cron" gorm:"cron"`
}

func init() {
	init.DB.AutoMigrate(new(Job))	// 自动迁移Job
}