package models

import (
	"fmt"
	"one-dock/core/comm"

	"gorm.io/gorm"
)

// UploadModel 文件上传
type UploadModel struct {
	ID       int64  `gorm:"primaryKey;autoIncrement:true" json:"id"`
	FileHash string `gorm:"type:varchar(64);not null;comment:文件MD5或SHA256" json:"file_hash"`
	FileSize int64  `gorm:"not null;comment:文件大小(字节)" json:"file_size"`
	FileExt  string `gorm:"type:varchar(255);not null;default:'';comment:文件扩展名" json:"file_ext"`
	FileName string `gorm:"type:varchar(255);not null;default:'';comment:文件名称" json:"file_name"`
	//TotalParts int64  `gorm:"not null;default:0;comment:总分片数" json:"total_parts"`
	HasUploaded bool  `gorm:"type:bool;not null;default:0;comment:是否已上传完成" json:"has_uploaded"`
	DirId       int64 `gorm:"not null;default:0;comment:目录ID" json:"dir_id"`
	comm.BaseModel
}

func (UploadModel) TableName() string {
	return "upload"
}

func InitUpload(db *gorm.DB) error {
	err := db.AutoMigrate(&UploadModel{})
	if err != nil {
		return fmt.Errorf("数据表Upload迁移失败: %w", err)
	}
	return nil
}
