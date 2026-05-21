package entity

import "mime/multipart"

// InitChunkUploadRequest 初始化分片上传请求
type InitChunkUploadRequest struct {
	FileName string `json:"file_name" validate:"required" alias:"文件名"`
	FileSize int64  `json:"file_size" validate:"required" alias:"文件大小"`
	FileExt  string `json:"file_ext"  alias:"文件扩展名"`
	FileHash string `json:"file_hash" validate:"required" alias:"文件哈希"`
	DirId    string `json:"dir_id" alias:"目录ID"`
}

// UploadChunkRequest 上传分片请求
type UploadChunkRequest struct {
	UploadId  int64                 `form:"upload_id" validate:"required" alias:"上传ID"`
	ChunkHash string                `form:"chunk_hash" validate:"required" alias:"分片哈希"`
	Offset    int64                 `form:"offset" alias:"分片偏移量"`
	Chunk     *multipart.FileHeader `form:"chunk" validate:"required" alias:"分片内容"`
}

// CreateUploadDto 创建上传记录请求DTO
type CreateUploadDto struct {
	FileName string `json:"file_name" validate:"required" alias:"文件名"`
	FileSize int64  `json:"file_size" validate:"required" alias:"文件大小"`
	FileExt  string `json:"file_ext"  alias:"文件扩展名"`
	FileHash string `json:"file_hash" validate:"required" alias:"文件哈希"`
	DirId    string `json:"dir_id" alias:"目录ID"`
}

// CreateChunkDto 创建上传分片记录DTO
type CreateChunkDto struct {
	UploadId int64
	Hash     string
	Offset   int64
	Size     int64
	FileHash string
}

// MergeChunkRequest 合并分片请求
type MergeChunkRequest struct {
	UploadId int64 `json:"upload_id" validate:"required" alias:"上传ID"`
}
