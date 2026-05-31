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
	if dirId == "" || dirId == "0" {
		return nil
	}
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

func (r *repository) QueryStorageList(ctx context.Context, parentId string, category int64) ([]models.StorageModel, error) {
	var list []models.StorageModel
	query := r.db.WithContext(ctx).
		Where("parent_id = ?", cast.ToInt64(parentId)).
		Order("is_folder DESC").
		Order("updated_at DESC")
	if category > 0 {
		query = query.Where("category = ? AND is_folder = ?", category, false)
	}
	err := query.Find(&list).Error
	return list, err
}

func (r *repository) QueryStorageById(ctx context.Context, id string) (*models.StorageModel, error) {
	var storage models.StorageModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&storage).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("文件不存在！")
		}
		return nil, err
	}
	return &storage, nil
}

func (r *repository) RenameStorage(ctx context.Context, id string, name string) error {
	return r.db.WithContext(ctx).
		Model(&models.StorageModel{}).
		Where("id = ?", id).
		Update("file_name", name).Error
}

func (r *repository) DeleteStorageRecords(ctx context.Context, ids []string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var records []models.StorageModel
		if err := tx.Where("id IN ?", ids).Find(&records).Error; err != nil {
			return err
		}
		for _, record := range records {
			if !record.IsFolder && record.FileHash != "" {
				if err := tx.Model(&models.FileModel{}).
					Where("hash = ?", record.FileHash).
					Update("ref_count", gorm.Expr("CASE WHEN ref_count > 0 THEN ref_count - 1 ELSE 0 END")).Error; err != nil {
					return err
				}
			}
		}
		return tx.Where("id IN ?", ids).Delete(&models.StorageModel{}).Error
	})
}

func (r *repository) QueryStorageStats(ctx context.Context) (entity.StorageStatsResponse, error) {
	var stats entity.StorageStatsResponse
	base := func() *gorm.DB {
		return r.db.WithContext(ctx).Model(&models.StorageModel{}).Where("is_folder = ?", false)
	}
	if err := base().Count(&stats.TotalFiles).Error; err != nil {
		return stats, err
	}
	if err := base().Select("COALESCE(SUM(file_size), 0)").Scan(&stats.TotalSize).Error; err != nil {
		return stats, err
	}
	if err := base().Where("category = ?", entity.ImageCategory).Count(&stats.ImageCount).Error; err != nil {
		return stats, err
	}
	if err := base().Where("category = ?", entity.VideoCategory).Count(&stats.VideoCount).Error; err != nil {
		return stats, err
	}
	if err := base().Where("category = ?", entity.AudioCategory).Count(&stats.AudioCount).Error; err != nil {
		return stats, err
	}
	if err := base().Where("category = ?", entity.DocumentCategory).Count(&stats.DocCount).Error; err != nil {
		return stats, err
	}
	if err := base().Where("category = ?", entity.OtherCategory).Count(&stats.OtherCount).Error; err != nil {
		return stats, err
	}
	return stats, nil
}
