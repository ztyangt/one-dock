package repository

import (
	"context"
	"errors"
	"one-dock/core/features/storage/models"

	"gorm.io/gorm"
)

func (r *repository) QueryUploadRecord(ctx context.Context, uploadId int64) (*models.UploadModel, error) {
	uploadModel := &models.UploadModel{}
	err := r.db.WithContext(ctx).
		Where("id = ?", uploadId).
		First(uploadModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("上传记录不存在！")
		}
		return nil, err
	}
	return uploadModel, nil
}
