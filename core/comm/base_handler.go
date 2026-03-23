package comm

import (
	e "one-dock/app/error"
	"one-dock/pkgs/validator"

	"github.com/gofiber/fiber/v3"
)

type BaseHandler struct {
}

// BindAndValidate 绑定并验证请求参数
func (h *BaseHandler) BindAndValidate(c fiber.Ctx, req any) error {
	method := c.Method()
	if method == "GET" {
		err := c.Bind().Query(req)
		if err != nil {
			return e.New(400, err.Error())
		}
	} else {
		err := c.Bind().Body(req)
		if err != nil {
			return e.New(400, err.Error())
		}
	}

	err := validator.Valid(req)
	if err != nil {
		return e.New(400, err.Error())
	}
	return nil
}
