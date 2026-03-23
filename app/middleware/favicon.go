package middleware

import (
	"github.com/gofiber/fiber/v3"
)

func FaviconMiddleware() fiber.Handler {

	// 重定向 /favicon.ico 到 /favicon.svg
	return func(ctx fiber.Ctx) error {
		if ctx.Path() == "/favicon.ico" {
			return ctx.Redirect().To("/favicon.svg")
		}
		return ctx.Next()
	}
}
