package template

import (
	e "one-dock/app/error"
	"one-dock/app/response"
	"one-dock/core/comm"

	"github.com/gofiber/fiber/v3"
)

// handler 完全私有，不需要对外暴露任何接口
type handler struct {
	comm.BaseHandler
	svc Service
}

// newHandler 私有构造函数，注入 Service
func newHandler(svc Service) *handler {
	return &handler{svc: svc}
}

// Take 处理获取模板请求
func (h *handler) Take(c fiber.Ctx) error {

	var req GetByIdReq
	if err := h.BindAndValidate(c, &req); err != nil {
		return err
	}

	u, err := h.svc.GetTemplateById(c.Context(), req.ID)
	if err != nil {
		return e.New(400, err.Error())
	}
	return response.Success(c, u)

}

// Create 创建模板
func (h *handler) Create(c fiber.Ctx) error {
	var req CreateReq
	if err := h.BindAndValidate(c, &req); err != nil {
		return err
	}
	err := h.svc.CreateTemplate(c.Context(), &req)
	if err != nil {
		return e.New(400, err.Error())
	}
	return response.Success(c, nil)
}

// Update 更新模板
func (h *handler) Update(c fiber.Ctx) error {
	var req UpdateReq
	if err := h.BindAndValidate(c, &req); err != nil {
		return err
	}
	err := h.svc.UpdateTemplate(c.Context(), &req)
	if err != nil {
		return e.New(400, err.Error())
	}
	return response.Success(c, nil)
}

func (h *handler) Delete(c fiber.Ctx) error {
	var req DeleteReq
	if err := h.BindAndValidate(c, &req); err != nil {
		return err
	}
	err := h.svc.DeleteTemplate(c.Context(), &req)
	if err != nil {
		return e.New(400, err.Error())
	}
	return response.Success(c, nil)
}
