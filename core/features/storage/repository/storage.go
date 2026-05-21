package repository

import (
	"context"
	"errors"
	"one-dock/core/features/storage/entity"
	"one-dock/core/features/storage/models"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// CheckDirExist 检查目录是否存在
func (r *repository) CheckDirExist(ctx context.Context, dirId string) error {
	err := r.db.WithContext(ctx).
		Where("id = ?", dirId).
		Where("is_folder = ?", true).
		First(&models.StorageModel{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("目录不存在！")
		}
		return err
	}
	return nil
}

// CheckStorageExistByFileNameAndDirId 通过文件名和目录id判断逻辑文件是否存在
func (r *repository) CheckStorageExistByFileNameAndDirId(ctx context.Context, fileName string, dirId string) (bool, error) {
	err := r.db.WithContext(ctx).
		Model(&models.StorageModel{}).
		Where("file_name = ? AND parent_id = ?", fileName, cast.ToInt64(dirId)).
		First(&models.StorageModel{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *repository) CreateStorageRecord(ctx context.Context, req *entity.CreateStorageDto) (int64, error) {
	storageModel := models.StorageModel{
		FileID:   req.FileId,
		FileName: req.FileName,
		ParentID: cast.ToInt64(req.DirId),
		IsFolder: req.IsFolder,
		MIMEType: req.FileMimeType,
		Category: req.Category,
		FileExt:  req.FileExt,
		FileSize: req.FileSize,
		FileHash: req.FileHash,
	}
	err := r.db.WithContext(ctx).Create(&storageModel).Error
	if err != nil {
		return 0, err
	}
	return storageModel.ID, nil
}
