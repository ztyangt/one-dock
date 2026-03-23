package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func CorsMiddleware() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		// 必须允许 Range 头，否则部分播放器在跨Z域时无法拖动进度条
		AllowHeaders: []string{"Range", "Content-Length", "Content-Type", "Authorization", "X-Requested-With"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH", "HEAD"},
		//AllowCredentials: true,
	})
}
