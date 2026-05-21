package services

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"one-dock/core/features/storage/entity"
	"os"
	"path/filepath"
	"strings"
)

// UploadChunk 上传分片
func (s *service) UploadChunk(ctx context.Context, req *entity.UploadChunkRequest) error {
	// 1. 检查上传记录
	uploadModel, err := s.repo.QueryUploadRecord(ctx, req.UploadId)
	if err != nil {
		return err
	}

	// 2. 创建临时目录（如果不存在）
	chunkDir := filepath.Join("storage", ".temp", uploadModel.FileHash)
	if err := os.MkdirAll(chunkDir, 0755); err != nil {
		return fmt.Errorf("创建临时目录失败: %w", err)
	}

	// 3. 打开上传的分片文件
	file, err := req.Chunk.Open()
	if err != nil {
		return fmt.Errorf("打开分片文件失败: %w", err)
	}
	defer file.Close()

	// 4. 创建目标文件
	chunkPath := filepath.Join(chunkDir, req.ChunkHash)
	targetFile, err := os.Create(chunkPath)
	if err != nil {
		return fmt.Errorf("创建目标文件失败: %w", err)
	}

	// 5. 将分片内容写入目标文件
	var chunkHasher hash.Hash
	switch len(req.ChunkHash) {
	case 32:
		chunkHasher = md5.New()
	case 64:
		chunkHasher = sha256.New()
	}
	writer := io.Writer(targetFile)
	if chunkHasher != nil {
		writer = io.MultiWriter(targetFile, chunkHasher)
	}
	if _, err := io.Copy(writer, file); err != nil {
		_ = targetFile.Close()
		return fmt.Errorf("写入分片文件失败: %w", err)
	}
	if err := targetFile.Close(); err != nil {
		return fmt.Errorf("关闭分片文件失败: %w", err)
	}
	if chunkHasher != nil && !strings.EqualFold(hex.EncodeToString(chunkHasher.Sum(nil)), req.ChunkHash) {
		_ = os.Remove(chunkPath)
		return fmt.Errorf("分片哈希校验失败")
	}

	// 6. 创建分片上传记录
	if err := s.repo.CreateChunkRecord(ctx, &entity.CreateChunkDto{
		UploadId: req.UploadId,
		Hash:     req.ChunkHash,
		Offset:   req.Offset,
		Size:     req.Chunk.Size,
		FileHash: uploadModel.FileHash,
	}); err != nil {
		return fmt.Errorf("创建分片上传记录失败: %w", err)
	}

	return nil
}
