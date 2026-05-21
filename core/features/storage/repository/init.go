package repository

import (
	"context"
	"one-dock/core/features/storage/entity"
	"one-dock/core/features/storage/models"

	"gorm.io/gorm"
)

type Repository interface {
	// CheckDirExist 检查目录是否存在
	CheckDirExist(ctx context.Context, dirId string) error

	// CheckFileExist 检查真实文件是否存在
	CheckFileExist(ctx context.Context, fileHash string) (*models.FileModel, error)

	// CheckStorageExistByFileNameAndDirId 通过文件名和目录id判断逻辑文件是否存在
	CheckStorageExistByFileNameAndDirId(ctx context.Context, fileName string, dirId string) (bool, error)

	// CheckUploadRecordExistByFileNameAndDirId 查上传记录是否存在（同名+目录id+文件哈希）
	CheckUploadRecordExistByFileNameAndDirId(ctx context.Context, fileName string, dirId string, fileHash string) (int64, error)

	// CreateUploadRecord 创建上传记录
	CreateUploadRecord(ctx context.Context, req *entity.CreateUploadDto) (int64, error)

	// UpdateUploadStatus 更新上传记录状态
	UpdateUploadStatus(ctx context.Context, uploadId int64, hasUploaded bool) error

	//	CreateStorageRecord 创建逻辑存储记录
	CreateStorageRecord(ctx context.Context, req *entity.CreateStorageDto) (int64, error)

	// UpdateFileRefCount 更新文件物理存储引用计数
	// delta 为增加或减少的引用计数
	UpdateFileRefCount(ctx context.Context, fileHash string, delta int) (int64, error)

	// CreateFileRecord 创建真实文件记录
	CreateFileRecord(ctx context.Context, req *entity.CreateFileDto) (int64, error)

	// QueryUploadRecord 查询上传记录
	QueryUploadRecord(ctx context.Context, uploadId int64) (*models.UploadModel, error)

	//	QueryChunkRecord 查询上传分片记录
	QueryChunkRecord(ctx context.Context, uploadId string) ([]string, error)

	// QueryChunkRecords 查询上传分片完整记录
	QueryChunkRecords(ctx context.Context, uploadId int64) ([]models.ChunkModel, error)

	// CreateChunkRecord 创建上传分片记录
	CreateChunkRecord(ctx context.Context, req *entity.CreateChunkDto) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
