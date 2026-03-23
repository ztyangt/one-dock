package template

import (
	"fmt"
	"one-dock/core/comm"

	"gorm.io/gorm"
)

type TemplateModel struct {
	ID    int64  `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
	Name  string `gorm:"size:256;comment:配置名称" json:"name"`
	Value string `gorm:"type:text;comment:配置值" json:"value"`
	comm.BaseModel
}

func (TemplateModel) TableName() string {
	return "template"
}

func InitTemplate(db *gorm.DB) error {
	err := db.AutoMigrate(&TemplateModel{})
	if err != nil {
		return fmt.Errorf("数据表Template迁移失败: %w", err)
	}
	return nil
}
