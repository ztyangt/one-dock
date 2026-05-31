package entity

type StorageItemResponse struct {
	Id        string   `json:"id"`
	ParentId  string   `json:"parent_id"`
	FileId    string   `json:"file_id"`
	FileName  string   `json:"file_name"`
	FileHash  string   `json:"file_hash"`
	IsFolder  bool     `json:"is_folder"`
	Category  Category `json:"category"`
	FileExt   string   `json:"file_ext"`
	FileSize  int64    `json:"file_size"`
	MimeType  string   `json:"mime_type"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
}

type StorageListResponse struct {
	List []StorageItemResponse `json:"list"`
}

type StorageStatsResponse struct {
	TotalFiles int64 `json:"total_files"`
	TotalSize  int64 `json:"total_size"`
	ImageCount int64 `json:"image_count"`
	VideoCount int64 `json:"video_count"`
	AudioCount int64 `json:"audio_count"`
	DocCount   int64 `json:"doc_count"`
	OtherCount int64 `json:"other_count"`
}
