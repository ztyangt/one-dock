package user

import (
	"one-dock/app/config"
	e "one-dock/app/error"
	"one-dock/app/response"
	"one-dock/core/comm"

	"github.com/gofiber/fiber/v3"
)

// handler 完全私有，不需要对外暴露任何接口
type handler struct {
	comm.BaseHandler
	svc Service
	cfg *config.Cfg
}

// newHandler 私有构造函数，注入 Service
func newHandler(svc Service, cfg *config.Cfg) *handler {
	return &handler{svc: svc, cfg: cfg}
}

// getUser 处理获取信息请求
func (h *handler) getUser(c fiber.Ctx) error {

	var req getUserRequest
	if err := h.BindAndValidate(c, &req); err != nil {
		return err
	}

	u, err := h.svc.UserInfo(c.Context(), req.ID)
	if err != nil {
		return e.New(400, err.Error())
	}
	return response.Success(c, u)

}

// login 处理登录请求
func (h *handler) login(c fiber.Ctx) error {
	var req loginRequest

	if err := h.BindAndValidate(c, &req); err != nil {
		return err
	}
	resp, err := h.svc.Login(c.Context(), &req)
	if err != nil {
		return e.New(400, err.Error())
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = resp.Token
	cookie.MaxAge = int(h.cfg.JWT.Expire)
	cookie.Path = "/"
	cookie.Secure = false
	cookie.HTTPOnly = false
	cookie.SameSite = "Lax"
	c.Cookie(cookie)

	return response.Success(c, resp, "登录成功！")
}
