package services

import (
	"context"
	"one-dock/core/features/storage/entity"
	"one-dock/core/features/storage/repository"
)

type Service interface {
	// ListStorage 查询文件列表
	ListStorage(ctx context.Context, req *entity.StorageListRequest) (entity.StorageListResponse, error)
	// CreateFolder 创建文件夹
	CreateFolder(ctx context.Context, req *entity.CreateFolderRequest) (entity.StorageItemResponse, error)
	// RenameStorage 重命名文件
	RenameStorage(ctx context.Context, req *entity.RenameStorageRequest) error
	// DeleteStorage 删除文件
	DeleteStorage(ctx context.Context, req *entity.DeleteStorageRequest) error
	// StorageStats 查询统计
	StorageStats(ctx context.Context) (entity.StorageStatsResponse, error)
	// DownloadPath 查询下载文件路径
	DownloadPath(ctx context.Context, id string) (string, string, error)
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
