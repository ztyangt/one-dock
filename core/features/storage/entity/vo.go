package entity

// InitChunkUploadResponse 初始化分片上传响应
type InitChunkUploadResponse struct {
	UploadId   int64    `json:"upload_id"`
	FileId     string   `json:"file_id"`
	ChunkParts []string `json:"chunk_parts"`
}

type MergeChunkResponse struct {
	FileId string `json:"file_id"`
}
