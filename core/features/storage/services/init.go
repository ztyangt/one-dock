package services

import (
	"context"
	"one-dock/core/features/storage/entity"
	"one-dock/core/features/storage/repository"
)

type Service interface {
	// InitChunkUpload 初始化分片上传
	InitChunkUpload(ctx context.Context, req *entity.InitChunkUploadRequest) (entity.InitChunkUploadResponse, error)
	// UploadChunk 上传分片
	UploadChunk(ctx context.Context, req *entity.UploadChunkRequest) error
	// MergeChunk 合并分片
	MergeChunk(ctx context.Context, req *entity.MergeChunkRequest) (entity.MergeChunkResponse, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}
