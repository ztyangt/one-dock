package repository

import (
	"context"
	"one-dock/core/features/storage/entity"
	"one-dock/core/features/storage/models"
)

// QueryChunkRecord 查询上传分片记录
func (r *repository) QueryChunkRecord(ctx context.Context, uploadId string) ([]string, error) {
	var hashes []string
	if err := r.db.WithContext(ctx).Model(&models.ChunkModel{}).
		Where("upload_id = ?", uploadId).
		Order("offset asc").
		Pluck("hash", &hashes).Error; err != nil {
		return nil, err
	}
	return hashes, nil
}

func (r *repository) QueryChunkRecords(ctx context.Context, uploadId int64) ([]models.ChunkModel, error) {
	var chunks []models.ChunkModel
	if err := r.db.WithContext(ctx).
		Where("upload_id = ?", uploadId).
		Order("offset asc").
		Find(&chunks).Error; err != nil {
		return nil, err
	}
	return chunks, nil
}

func (r *repository) CreateChunkRecord(ctx context.Context, req *entity.CreateChunkDto) error {
	var chunk models.ChunkModel
	return r.db.WithContext(ctx).
		Where("upload_id = ? AND hash = ?", req.UploadId, req.Hash).
		Assign(models.ChunkModel{
			Size:     req.Size,
			Offset:   req.Offset,
			FileHash: req.FileHash,
		}).
		FirstOrCreate(&chunk, models.ChunkModel{
			UploadId: req.UploadId,
			Hash:     req.Hash,
		}).Error
}
