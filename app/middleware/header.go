package middleware

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
)

func HeaderMiddleware() fiber.Handler {
	return func(ctx fiber.Ctx) error {

		start := time.Now()

		ctx.Set("X-Server-Name", "fiber-server")
		ctx.Set("X-Start-Time", start.Format(time.RFC3339Nano))

		err := ctx.Next()
		if err != nil {
			return err
		}

		// 计算时间指标
		latency := time.Since(start)
		latencyMs := float64(latency.Nanoseconds()) / 1e6

		ctx.Set("X-Processing-Time", strconv.FormatFloat(latencyMs, 'f', 3, 64)+" ms")

		return nil
	}
}
