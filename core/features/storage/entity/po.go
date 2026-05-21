package entity

type CreateStorageDto struct {
	FileId       int64
	FileName     string
	FileSize     int64
	FileExt      string
	FileHash     string
	DirId        string
	FileMimeType string
	Category     Category
	IsFolder     bool
}

type CreateFileDto struct {
	Hash      string
	Size      int64
	Extension string
	MimeType  string
	Category  Category
}
