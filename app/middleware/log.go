package middleware

import (
	"one-dock/app/logger"

	"github.com/gofiber/fiber/v3"
)

func LoggerMiddleware(logFactor *logger.LogFactor) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		logFactor.Info(map[string]any{
			"ip":     ctx.IP(),
			"path":   ctx.Path(),
			"method": ctx.Method(),
		})
		return ctx.Next()
	}
}
