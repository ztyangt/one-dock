package config

import (
	e "one-dock/app/error"
	"one-dock/app/response"
	"one-dock/core/comm"

	"github.com/gofiber/fiber/v3"
)

type handler struct {
	comm.BaseHandler
	svc Service
}

func newHandler(svc Service) *handler {
	return &handler{svc: svc}
}

// Take 获取配置项
func (h *handler) Take(c fiber.Ctx) error {
	var req GetReq
	if err := h.BindAndValidate(c, &req); err != nil {
		return err
	}

	config, err := h.svc.GetConfig(c.Context(), &req)
	if err != nil {
		return err
	}
	return response.Success(c, config)
}

// Create 创建配置项
func (h *handler) Create(c fiber.Ctx) error {
	req := new(CreateReq)
	if err := h.BindAndValidate(c, req); err != nil {
		return err
	}

	if err := h.svc.CreateConfig(c.Context(), req); err != nil {
		return e.New(400, err.Error())
	}
	return response.Success(c, nil)
}

// Update 更新配置项
func (h *handler) Update(c fiber.Ctx) error {
	var req UpdateReq
	if err := h.BindAndValidate(c, &req); err != nil {
		return err
	}

	if err := h.svc.UpdateConfig(c.Context(), &req); err != nil {
		return e.New(400, err.Error())
	}
	return response.Success(c, nil)
}

// Delete 删除配置项
func (h *handler) Delete(c fiber.Ctx) error {
	var req DeleteReq
	if err := h.BindAndValidate(c, &req); err != nil {
		return err
	}

	if err := h.svc.DeleteConfig(c.Context(), &req); err != nil {
		return e.New(400, err.Error())
	}
	return response.Success(c, nil)
}
