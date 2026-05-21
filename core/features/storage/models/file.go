package models

import (
	"fmt"
	"one-dock/core/comm"
	"one-dock/core/features/storage/entity"

	"gorm.io/gorm"
)

// FileModel 物理文件
type FileModel struct {
	ID        int64           `gorm:"primaryKey;autoIncrement:false" json:"id"`
	Hash      string          `gorm:"type:varchar(64);uniqueIndex;not null;comment:文件MD5或SHA256" json:"hash"`
	Size      int64           `gorm:"not null;comment:文件大小(字节)" json:"size"`
	Extension string          `gorm:"type:varchar(32);comment:文件扩展名" json:"extension"`
	MimeType  string          `gorm:"size:100; comment:MIME类型;" json:"mime_type"`
	RefCount  int             `gorm:"not null;default:0;comment:引用计数(0时可清理物理文件)" json:"ref_count"`
	Category  entity.Category `gorm:"index;not null;default:0;comment:文件分类" json:"category"`
	comm.BaseModel
}

func (FileModel) TableName() string {
	return "files"
}

func InitFile(db *gorm.DB) error {
	err := db.AutoMigrate(&FileModel{})
	if err != nil {
		return fmt.Errorf("数据表File迁移失败: %w", err)
	}
	return nil
}

func (f *FileModel) BeforeCreate(tx *gorm.DB) (err error) {
	// 生成id
	if f.ID == 0 {
		f.ID = comm.SnowNode.Generate().Int64()
	}
	return
}
