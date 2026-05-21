package models

import (
	"fmt"
	"one-dock/core/comm"
	"one-dock/core/features/storage/entity"

	"gorm.io/gorm"
)

// StorageModel 逻辑文件
type StorageModel struct {
	ID       int64           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	ParentID int64           `gorm:"index;not null;default:0;comment:父目录ID(0为根目录)" json:"parent_id"`
	FileID   int64           `gorm:"index;comment:物理文件ID(如果是文件夹则为null)" json:"file_id"`
	FileName string          `gorm:"type:varchar(255);not null;comment:文件/文件夹名称" json:"file_name"`
	FileHash string          `gorm:"type:varchar(255);not null;default:'';comment:文件哈希" json:"file_hash"`
	IsFolder bool            `gorm:"not null;default:false;comment:是否为文件夹" json:"is_folder"`
	Category entity.Category `gorm:"index;not null;default:0;comment:文件分类" json:"category"`
	FileExt  string          `gorm:"type:varchar(255);not null;default:'';comment:文件扩展名" json:"file_ext"`
	FileSize int64           `gorm:"not null;default:0;comment:文件大小" json:"file_size"`
	MIMEType string          `gorm:"type:varchar(255);not null;default:'';comment:文件MIME类型" json:"mime_type"`
	comm.BaseModel
}

func (StorageModel) TableName() string {
	return "storage"
}

func InitStorage(db *gorm.DB) error {
	err := db.AutoMigrate(&StorageModel{})
	if err != nil {
		return fmt.Errorf("数据表Storage迁移失败: %w", err)
	}
	return nil
}

func (f *StorageModel) BeforeCreate(tx *gorm.DB) (err error) {
	if f.ID == 0 {
		f.ID = comm.SnowNode.Generate().Int64()
	}
	return
}
