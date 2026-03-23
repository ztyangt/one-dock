package middleware

import (
	"one-dock/app/config"
	e "one-dock/app/error"
	"one-dock/core/comm"

	"github.com/gofiber/fiber/v3"
	"github.com/spf13/cast"
)

func AdminMiddleware(cfg *config.Cfg) fiber.Handler {
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

		role := user.(map[string]interface{})["role"]
		if comm.UserRole(cast.ToInt64(role)) != comm.UserRoleAdmin {
			return e.New(403, "该操作须具有管理员权限，您无权操作！")
		}

		return c.Next()
	}
}
