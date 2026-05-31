package repository

import (
	"context"
	"errors"
	"one-dock/core/features/storage/entity"
	"one-dock/core/features/storage/models"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// CheckUploadRecordExistByFileNameAndDirId 检查上传记录是否存在（同名+目录id+文件哈希）
func (r *repository) CheckUploadRecordExistByFileNameAndDirId(ctx context.Context, fileName string, dirId string, fileHash string) (int64, error) {
	var upload models.UploadModel
	err := r.db.WithContext(ctx).
		Where("file_name = ? AND dir_id = ? AND file_hash = ?", fileName, cast.ToInt64(dirId), fileHash).
		First(&upload).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, err
	}
	return upload.ID, nil
}

// CreateUploadRecord 创建上传记录
func (r *repository) CreateUploadRecord(ctx context.Context, req *entity.CreateUploadDto) (int64, error) {
	uploadModel := &models.UploadModel{
		FileName: req.FileName,
		FileSize: req.FileSize,
		FileExt:  req.FileExt,
		FileHash: req.FileHash,
		DirId:    cast.ToInt64(req.DirId),
	}

	result := r.db.WithContext(ctx).Create(uploadModel)
	if result.Error != nil {
		return 0, result.Error
	}

	return uploadModel.ID, nil
}

// UpdateUploadStatus 更新上传记录状态
func (r *repository) UpdateUploadStatus(ctx context.Context, uploadId int64, hasUploaded bool) error {
	return r.db.WithContext(ctx).
		Model(&models.UploadModel{}).
		Where("id = ?", uploadId).
		Update("has_uploaded", hasUploaded).
		Error
}
