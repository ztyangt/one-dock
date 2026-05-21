package models

import (
	"fmt"

	"gorm.io/gorm"
)

// ChunkModel 文件分片
type ChunkModel struct {
	ID       int64  `gorm:"primaryKey;autoIncrement:true" json:"id"`
	UploadId int64  `gorm:"not null;comment:上传记录ID" json:"upload_id"`
	Hash     string `gorm:"type:varchar(64);not null;comment:文件MD5或SHA256" json:"hash"`
	Size     int64  `gorm:"not null;comment:文件大小(字节)" json:"size"`
	Offset   int64  `gorm:"not null;comment:偏移量(字节)" json:"offset"`
	FileHash string `gorm:"type:varchar(64);not null;comment:物理文件MD5或SHA256" json:"file_hash"`
}

func (ChunkModel) TableName() string {
	return "chunks"
}

func InitChunk(db *gorm.DB) error {
	err := db.AutoMigrate(&ChunkModel{})
	if err != nil {
		return fmt.Errorf("数据表Chunk迁移失败: %w", err)
	}
	return nil
}
