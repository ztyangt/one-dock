package install

import (
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

func (h *handler) InstallBySqlite(c fiber.Ctx) error {
	var req BySqliteRequest
	if err := h.BindAndValidate(c, &req); err != nil {
		return err
	}
	return response.Success(c, nil)
}
