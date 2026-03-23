package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	e "one-dock/app/error"
)

// InstallMiddleware 安装检查中间件
func InstallMiddleware() fiber.Handler {
	return func(ctx fiber.Ctx) error {

		isInstallPath := strings.HasPrefix(ctx.Path(), "/api/install")

		// 检查是否安装（install.lock文件存在）
		if _, err := os.Stat("install.lock"); err == nil {
			if isInstallPath {
				return e.New(http.StatusServiceUnavailable, "系统已安装，禁止访问！")
			}
			return ctx.Next()
		}

		if isInstallPath {
			return ctx.Next()
		}

		return e.New(http.StatusNotImplemented, "系统未安装！")

	}
}
