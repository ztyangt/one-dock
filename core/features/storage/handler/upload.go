package handler

import (
	e "one-dock/app/error"
	"one-dock/app/response"
	"one-dock/core/comm"
	"one-dock/core/features/storage/entity"
	"one-dock/core/features/storage/services"

	"github.com/gofiber/fiber/v3"
)

type UploadHandler interface {
	InitChunkUpload(ctx fiber.Ctx) error
	UploadChunk(ctx fiber.Ctx) error
	MergeChunk(ctx fiber.Ctx) error
}

type uploadHandler struct {
	comm.BaseHandler
	svc services.Service
}

func newUploadHandler(svc services.Service) UploadHandler {
	return &uploadHandler{svc: svc}
}

// InitChunkUpload 初始化分片上传
func (h *uploadHandler) InitChunkUpload(c fiber.Ctx) error {
	req := new(entity.InitChunkUploadRequest)

	if err := h.BindAndValidate(c, req); err != nil {
		return err
	}

	resp, err := h.svc.InitChunkUpload(c.Context(), req)
	if err != nil {
		return e.New(400, err.Error())
	}

	return response.Success(c, resp)
}

// UploadChunk 上传分片
func (h *uploadHandler) UploadChunk(c fiber.Ctx) error {
	req := new(entity.UploadChunkRequest)

	if err := h.BindAndValidate(c, req); err != nil {
		return err
	}

	if err := h.svc.UploadChunk(c.Context(), req); err != nil {
		return e.New(400, err.Error())
	}

	return response.Success(c, nil)
}

// MergeChunk 合并分片
func (h *uploadHandler) MergeChunk(c fiber.Ctx) error {
	req := new(entity.MergeChunkRequest)

	if err := h.BindAndValidate(c, req); err != nil {
		return err
	}

	resp, err := h.svc.MergeChunk(c.Context(), req)
	if err != nil {
		return e.New(400, err.Error())
	}

	return response.Success(c, resp)
}
