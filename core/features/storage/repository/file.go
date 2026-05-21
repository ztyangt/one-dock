package repository

import (
	"context"
	"errors"
	"fmt"
	"one-dock/core/features/storage/entity"
	"one-dock/core/features/storage/models"

	"gorm.io/gorm"
)

// CheckFileExist 检查文件是否存在
func (r *repository) CheckFileExist(ctx context.Context, fileHash string) (*models.FileModel, error) {
	var file models.FileModel
	err := r.db.WithContext(ctx).
		Where("hash = ?", fileHash).
		First(&file).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &file, nil
}

// UpdateFileRefCount 更新文件物理存储引用计数
// delta 为增加或减少的引用计数
// 返回更新后的引用计数
func (r *repository) UpdateFileRefCount(ctx context.Context, fileHash string, delta int) (int64, error) {
	var result *gorm.DB

	// 如果是为了减少引用计数，确保不会变成负数
	if delta >= 0 {
		result = r.db.WithContext(ctx).
			Model(&models.FileModel{}).
			Where("hash = ?", fileHash).
			Update("ref_count", gorm.Expr("ref_count + ?", delta))
	} else {
		// 减少引用计数时，确保不会小于0
		result = r.db.WithContext(ctx).
			Model(&models.FileModel{}).
			Where("hash = ? AND ref_count + ? >= 0", fileHash, delta).
			Update("ref_count", gorm.Expr("ref_count + ?", delta))
	}

	if result.Error != nil {
		return 0, result.Error
	}

	// 检查是否有行被更新（文件是否存在）
	if result.RowsAffected == 0 {
		// 如果是减少引用计数且没有更新，可能是因为 ref_count 会变成负数
		if delta < 0 {
			return 0, fmt.Errorf("file %s reference count would become negative", fileHash)
		}
		// 如果是增加引用计数但没有更新，说明文件不存在
		return 0, fmt.Errorf("file %s not found", fileHash)
	}

	// 查询更新后的引用计数
	var file models.FileModel
	if err := r.db.WithContext(ctx).
		Select("ref_count").
		Where("hash = ?", fileHash).
		First(&file).Error; err != nil {
		return 0, fmt.Errorf("failed to get updated ref_count: %w", err)
	}

	return result.RowsAffected, nil
}

// CreateFileRecord 创建物理文件记录。
func (r *repository) CreateFileRecord(ctx context.Context, req *entity.CreateFileDto) (int64, error) {
	fileModel := &models.FileModel{
		Hash:      req.Hash,
		Size:      req.Size,
		Extension: req.Extension,
		MimeType:  req.MimeType,
		Category:  req.Category,
		RefCount:  1,
	}

	err := r.db.WithContext(ctx).Create(fileModel).Error
	if err != nil {
		return 0, err
	}
	return fileModel.ID, nil
}
