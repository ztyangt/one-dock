package entity

import "mime/multipart"

// StorageListRequest 文件列表请求
type StorageListRequest struct {
	ParentId string `query:"parent_id" json:"parent_id" alias:"父目录ID"`
	Category int64  `query:"category" json:"category" alias:"文件分类"`
}

// CreateFolderRequest 创建文件夹请求
type CreateFolderRequest struct {
	Name     string `json:"name" validate:"required" alias:"文件夹名称"`
	ParentId string `json:"parent_id" alias:"父目录ID"`
}

// RenameStorageRequest 重命名请求
type RenameStorageRequest struct {
	Id   string `json:"id" validate:"required" alias:"文件ID"`
	Name string `json:"name" validate:"required" alias:"文件名"`
}

// DeleteStorageRequest 删除请求
type DeleteStorageRequest struct {
	Ids []string `json:"ids" validate:"required" alias:"文件ID列表"`
}

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
