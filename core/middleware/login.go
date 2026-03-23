package middleware

import (
	"one-dock/app/config"
	e "one-dock/app/error"

	"github.com/gofiber/fiber/v3"
)

func LoginMiddleware(cfg *config.Cfg) fiber.Handler {
	return func(c fiber.Ctx) error {

		ctx := c.Context()
		err := ctx.Value("err")

		if err != nil {
			return err.(error)
		}

		user := ctx.Value("user")

		if user == nil {
			return e.New(401, "请先登录！")
		}

		return c.Next()
	}
}
