package services

import (
	"context"
	"one-dock/core/features/storage/entity"

	"github.com/spf13/cast"
)

func (s *service) InitChunkUpload(ctx context.Context, req *entity.InitChunkUploadRequest) (res entity.InitChunkUploadResponse, err error) {

	// 1. 如果传了目录id，检查目录是否存在
	if req.DirId != "" {
		err = s.repo.CheckDirExist(ctx, req.DirId)
		if err != nil {
			return
		}
	}

	// 2. 创建上传记录
	// 2.1 查询逻辑存储表，检查上传目标目录下是否有相同文件（同名），如果有则更改文件名
	exist, err := s.repo.CheckStorageExistByFileNameAndDirId(ctx, req.FileName, req.DirId)
	if err != nil {
		return
	}
	if exist {
		req.FileName = req.FileName + " (copy)"
	}

	// 2.2 查询上传记录表，检查是否有重复的上传记录（同名+目录id+文件哈希），如果有则复用该记录，否则创建新记录
	uploadId, err := s.repo.CheckUploadRecordExistByFileNameAndDirId(ctx, req.FileName, req.DirId, req.FileHash)
	if err != nil {
		return
	}
	if uploadId == 0 {
		uploadId, err = s.repo.CreateUploadRecord(ctx, &entity.CreateUploadDto{
			FileName: req.FileName,
			FileSize: req.FileSize,
			FileExt:  req.FileExt,
			FileHash: req.FileHash,
			DirId:    req.DirId,
		})
		if err != nil {
			return
		}
	}
	res.UploadId = uploadId

	// 3. 检查文件物理文件是否已存在，若存在则秒传
	file, err := s.repo.CheckFileExist(ctx, req.FileHash)
	if err != nil {
		return
	}

	if file != nil {
		// 秒传逻辑：更新上传记录表，将has_uploaded设置为true；创建逻辑存储表，将文件信息写入逻辑存储表；更新物理存储表引用
		err = s.repo.UpdateUploadStatus(ctx, uploadId, true)
		if err != nil {
			return
		}
		var storageId int64
		storageId, err = s.repo.CreateStorageRecord(ctx, &entity.CreateStorageDto{
			FileId:       file.ID,
			FileName:     req.FileName,
			FileSize:     req.FileSize,
			FileExt:      req.FileExt,
			FileHash:     req.FileHash,
			DirId:        req.DirId,
			FileMimeType: file.MimeType,
			Category:     file.Category,
			IsFolder:     false,
		})
		if err != nil {
			return
		}
		// 4.3 更新物理存储表引用计数
		_, err = s.repo.UpdateFileRefCount(ctx, req.FileHash, 1)
		if err != nil {
			return
		}
		// 返回逻辑文件id，实现秒传
		res.FileId = cast.ToString(storageId)
		return
	}

	// 4. 如果文件不存在，则查询已上传的分片列表
	var hashes []string
	hashes, err = s.repo.QueryChunkRecord(ctx, cast.ToString(uploadId))
	if err != nil {
		return
	}
	res.ChunkParts = hashes
	return
}
