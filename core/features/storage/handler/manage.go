package handler

import (
	e "one-dock/app/error"
	"one-dock/app/response"
	"one-dock/core/comm"
	"one-dock/core/features/storage/entity"
	"one-dock/core/features/storage/services"

	"github.com/gofiber/fiber/v3"
)

type ManageHandler interface {
	List(c fiber.Ctx) error
	CreateFolder(c fiber.Ctx) error
	Rename(c fiber.Ctx) error
	Delete(c fiber.Ctx) error
	Stats(c fiber.Ctx) error
	Download(c fiber.Ctx) error
}

type manageHandler struct {
	comm.BaseHandler
	svc services.Service
}

func newManageHandler(svc services.Service) ManageHandler {
	return &manageHandler{svc: svc}
}

func (h *manageHandler) List(c fiber.Ctx) error {
	req := new(entity.StorageListRequest)
	if err := h.BindAndValidate(c, req); err != nil {
		return err
	}
	res, err := h.svc.ListStorage(c.Context(), req)
	if err != nil {
		return e.New(400, err.Error())
	}
	return response.Success(c, res)
}

func (h *manageHandler) CreateFolder(c fiber.Ctx) error {
	req := new(entity.CreateFolderRequest)
	if err := h.BindAndValidate(c, req); err != nil {
		return err
	}
	res, err := h.svc.CreateFolder(c.Context(), req)
	if err != nil {
		return e.New(400, err.Error())
	}
	return response.Success(c, res)
}

func (h *manageHandler) Rename(c fiber.Ctx) error {
	req := new(entity.RenameStorageRequest)
	if err := h.BindAndValidate(c, req); err != nil {
		return err
	}
	if err := h.svc.RenameStorage(c.Context(), req); err != nil {
		return e.New(400, err.Error())
	}
	return response.Success(c, nil)
}

func (h *manageHandler) Delete(c fiber.Ctx) error {
	req := new(entity.DeleteStorageRequest)
	if err := h.BindAndValidate(c, req); err != nil {
		return err
	}
	if err := h.svc.DeleteStorage(c.Context(), req); err != nil {
		return e.New(400, err.Error())
	}
	return response.Success(c, nil)
}

func (h *manageHandler) Stats(c fiber.Ctx) error {
	res, err := h.svc.StorageStats(c.Context())
	if err != nil {
		return e.New(400, err.Error())
	}
	return response.Success(c, res)
}

func (h *manageHandler) Download(c fiber.Ctx) error {
	path, name, err := h.svc.DownloadPath(c.Context(), c.Params("id"))
	if err != nil {
		return e.New(400, err.Error())
	}
	return c.Download(path, name)
}
