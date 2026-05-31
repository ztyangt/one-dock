package services

import (
	"context"
	"errors"
	"one-dock/core/features/storage/entity"
	"one-dock/core/features/storage/models"
	"path/filepath"

	"github.com/spf13/cast"
)

func (s *service) ListStorage(ctx context.Context, req *entity.StorageListRequest) (entity.StorageListResponse, error) {
	list, err := s.repo.QueryStorageList(ctx, req.ParentId, req.Category)
	if err != nil {
		return entity.StorageListResponse{}, err
	}
	res := entity.StorageListResponse{List: make([]entity.StorageItemResponse, 0, len(list))}
	for _, item := range list {
		res.List = append(res.List, storageModelToResponse(item))
	}
	return res, nil
}

func (s *service) CreateFolder(ctx context.Context, req *entity.CreateFolderRequest) (entity.StorageItemResponse, error) {
	if err := s.repo.CheckDirExist(ctx, req.ParentId); err != nil {
		return entity.StorageItemResponse{}, err
	}
	exist, err := s.repo.CheckStorageExistByFileNameAndDirId(ctx, req.Name, req.ParentId)
	if err != nil {
		return entity.StorageItemResponse{}, err
	}
	if exist {
		return entity.StorageItemResponse{}, errors.New("同级目录下已存在同名文件或文件夹！")
	}
	id, err := s.repo.CreateStorageRecord(ctx, &entity.CreateStorageDto{
		FileName: req.Name,
		DirId:    req.ParentId,
		IsFolder: true,
	})
	if err != nil {
		return entity.StorageItemResponse{}, err
	}
	model, err := s.repo.QueryStorageById(ctx, cast.ToString(id))
	if err != nil {
		return entity.StorageItemResponse{}, err
	}
	return storageModelToResponse(*model), nil
}

func (s *service) RenameStorage(ctx context.Context, req *entity.RenameStorageRequest) error {
	model, err := s.repo.QueryStorageById(ctx, req.Id)
	if err != nil {
		return err
	}
	exist, err := s.repo.CheckStorageExistByFileNameAndDirId(ctx, req.Name, cast.ToString(model.ParentID))
	if err != nil {
		return err
	}
	if exist && req.Name != model.FileName {
		return errors.New("同级目录下已存在同名文件或文件夹！")
	}
	return s.repo.RenameStorage(ctx, req.Id, req.Name)
}

func (s *service) DeleteStorage(ctx context.Context, req *entity.DeleteStorageRequest) error {
	if len(req.Ids) == 0 {
		return errors.New("请选择要删除的文件！")
	}
	return s.repo.DeleteStorageRecords(ctx, req.Ids)
}

func (s *service) StorageStats(ctx context.Context) (entity.StorageStatsResponse, error) {
	return s.repo.QueryStorageStats(ctx)
}

func (s *service) DownloadPath(ctx context.Context, id string) (string, string, error) {
	model, err := s.repo.QueryStorageById(ctx, id)
	if err != nil {
		return "", "", err
	}
	if model.IsFolder {
		return "", "", errors.New("文件夹不支持下载！")
	}
	return filepath.Join(fileStoreRoot, model.FileHash), model.FileName, nil
}

func storageModelToResponse(item models.StorageModel) entity.StorageItemResponse {
	return entity.StorageItemResponse{
		Id:        cast.ToString(item.ID),
		ParentId:  cast.ToString(item.ParentID),
		FileId:    cast.ToString(item.FileID),
		FileName:  item.FileName,
		FileHash:  item.FileHash,
		IsFolder:  item.IsFolder,
		Category:  item.Category,
		FileExt:   item.FileExt,
		FileSize:  item.FileSize,
		MimeType:  item.MIMEType,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}
