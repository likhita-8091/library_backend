package model

import (
	"gorm.io/gorm"
	"time"
)

// 公共字段
type CommonModel struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

func NewCommonModel() *CommonModel {
	return &CommonModel{
		CreatedAt: time.Now(),
	}
}
