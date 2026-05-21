package services

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"mime"
	"net/http"
	"one-dock/core/features/storage/entity"
	"one-dock/core/features/storage/models"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cast"
)

const (
	chunkTempRoot = "storage/.temp"
	fileStoreRoot = "storage/files"
)

func (s *service) MergeChunk(ctx context.Context, req *entity.MergeChunkRequest) (entity.MergeChunkResponse, error) {
	uploadModel, err := s.repo.QueryUploadRecord(ctx, req.UploadId)
	if err != nil {
		return entity.MergeChunkResponse{}, err
	}
	if uploadModel.HasUploaded {
		return entity.MergeChunkResponse{}, fmt.Errorf("上传任务已完成")
	}

	chunks, err := s.repo.QueryChunkRecords(ctx, req.UploadId)
	if err != nil {
		return entity.MergeChunkResponse{}, err
	}
	if len(chunks) == 0 {
		return entity.MergeChunkResponse{}, fmt.Errorf("未找到已上传分片")
	}

	tempDir := filepath.Join(chunkTempRoot, uploadModel.FileHash)
	if err := validateChunks(tempDir, uploadModel.FileSize, chunks); err != nil {
		return entity.MergeChunkResponse{}, err
	}

	if err := os.MkdirAll(fileStoreRoot, 0755); err != nil {
		return entity.MergeChunkResponse{}, fmt.Errorf("创建文件存储目录失败: %w", err)
	}

	targetPath := filepath.Join(fileStoreRoot, uploadModel.FileHash)
	if err := mergeChunks(tempDir, targetPath, uploadModel.FileHash, chunks); err != nil {
		_ = os.Remove(targetPath)
		return entity.MergeChunkResponse{}, err
	}

	mimeType, category := detectFileMeta(targetPath, uploadModel.FileExt)
	fileId, err := s.ensureFileRecord(ctx, uploadModel.FileHash, uploadModel.FileSize, uploadModel.FileExt, mimeType, category)
	if err != nil {
		return entity.MergeChunkResponse{}, err
	}

	storageId, err := s.repo.CreateStorageRecord(ctx, &entity.CreateStorageDto{
		FileId:       fileId,
		FileName:     uploadModel.FileName,
		FileSize:     uploadModel.FileSize,
		FileExt:      uploadModel.FileExt,
		FileHash:     uploadModel.FileHash,
		DirId:        cast.ToString(uploadModel.DirId),
		FileMimeType: mimeType,
		Category:     category,
		IsFolder:     false,
	})
	if err != nil {
		return entity.MergeChunkResponse{}, err
	}

	if err := s.repo.UpdateUploadStatus(ctx, req.UploadId, true); err != nil {
		return entity.MergeChunkResponse{}, err
	}
	_ = os.RemoveAll(tempDir)

	return entity.MergeChunkResponse{FileId: cast.ToString(storageId)}, nil
}

func validateChunks(tempDir string, fileSize int64, chunks []models.ChunkModel) error {
	var expectedOffset int64
	for _, chunk := range chunks {
		if chunk.Offset != expectedOffset {
			return fmt.Errorf("分片不完整，缺少偏移量 %d 的分片", expectedOffset)
		}
		if chunk.Size <= 0 {
			return fmt.Errorf("分片大小非法: %s", chunk.Hash)
		}
		stat, err := os.Stat(filepath.Join(tempDir, chunk.Hash))
		if err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("分片文件不存在: %s", chunk.Hash)
			}
			return fmt.Errorf("读取分片文件失败: %w", err)
		}
		if stat.Size() != chunk.Size {
			return fmt.Errorf("分片大小不一致: %s", chunk.Hash)
		}
		expectedOffset += chunk.Size
	}
	if expectedOffset != fileSize {
		return fmt.Errorf("分片总大小不一致，期望 %d，实际 %d", fileSize, expectedOffset)
	}
	return nil
}

func mergeChunks(tempDir, targetPath, fileHash string, chunks []models.ChunkModel) error {
	fileHasher, err := newFileHasher(fileHash)
	if err != nil {
		return err
	}

	targetFile, err := os.Create(targetPath)
	if err != nil {
		return fmt.Errorf("创建合并文件失败: %w", err)
	}
	defer targetFile.Close()

	writer := io.MultiWriter(targetFile, fileHasher)
	for _, chunk := range chunks {
		chunkFile, err := os.Open(filepath.Join(tempDir, chunk.Hash))
		if err != nil {
			return fmt.Errorf("打开分片文件失败: %w", err)
		}
		if _, err := io.Copy(writer, chunkFile); err != nil {
			_ = chunkFile.Close()
			return fmt.Errorf("合并分片失败: %w", err)
		}
		if err := chunkFile.Close(); err != nil {
			return fmt.Errorf("关闭分片文件失败: %w", err)
		}
	}

	actualHash := hex.EncodeToString(fileHasher.Sum(nil))
	if !strings.EqualFold(actualHash, fileHash) {
		return fmt.Errorf("文件哈希校验失败，期望 %s，实际 %s", fileHash, actualHash)
	}
	return nil
}

func newFileHasher(fileHash string) (hash.Hash, error) {
	switch len(fileHash) {
	case 32:
		return md5.New(), nil
	case 64:
		return sha256.New(), nil
	default:
		return nil, fmt.Errorf("不支持的文件哈希长度")
	}
}

func (s *service) ensureFileRecord(ctx context.Context, fileHash string, size int64, ext string, mimeType string, category entity.Category) (int64, error) {
	file, err := s.repo.CheckFileExist(ctx, fileHash)
	if err != nil {
		return 0, err
	}
	if file != nil {
		_, err = s.repo.UpdateFileRefCount(ctx, fileHash, 1)
		if err != nil {
			return 0, err
		}
		return file.ID, nil
	}

	return s.repo.CreateFileRecord(ctx, &entity.CreateFileDto{
		Hash:      fileHash,
		Size:      size,
		Extension: normalizeExt(ext),
		MimeType:  mimeType,
		Category:  category,
	})
}

func detectFileMeta(path string, ext string) (string, entity.Category) {
	mimeType := mime.TypeByExtension(normalizeExt(ext))
	if mimeType == "" {
		file, err := os.Open(path)
		if err == nil {
			defer file.Close()
			buffer := make([]byte, 512)
			n, _ := file.Read(buffer)
			mimeType = http.DetectContentType(buffer[:n])
		}
	}
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	return mimeType, categoryByMimeType(mimeType)
}

func categoryByMimeType(mimeType string) entity.Category {
	switch {
	case strings.HasPrefix(mimeType, "image/"):
		return entity.ImageCategory
	case strings.HasPrefix(mimeType, "video/"):
		return entity.VideoCategory
	case strings.HasPrefix(mimeType, "audio/"):
		return entity.AudioCategory
	case strings.HasPrefix(mimeType, "text/"),
		strings.Contains(mimeType, "pdf"),
		strings.Contains(mimeType, "document"),
		strings.Contains(mimeType, "spreadsheet"),
		strings.Contains(mimeType, "presentation"):
		return entity.DocumentCategory
	default:
		return entity.OtherCategory
	}
}

func normalizeExt(ext string) string {
	if ext == "" {
		return ""
	}
	if strings.HasPrefix(ext, ".") {
		return ext
	}
	return "." + ext
}
