package system

import (
	"one-dock/app/response"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v3"
)

// systemInfo 获取系统信息
func systemInfo(c fiber.Ctx) error {
	return response.Success(c, map[string]interface{}{
		"domain":     c.BaseURL(),
		"go_version": runtime.Version(),
		"arch":       runtime.GOARCH,
		"os":         runtime.GOOS,
		"cpu_num":    runtime.NumCPU(),
		"date":       time.Now().Format("2006-01-02 15:04:05.000"),
		"unix":       time.Now().UnixMilli(),
	})
}
