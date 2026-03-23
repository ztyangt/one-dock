package install

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"one-dock/app/config"
)

func Setup(routerGroup fiber.Router, db *gorm.DB, cfg *config.Cfg) {

	// 1. 实例化依赖
	//repo := newRepository(db)
	svc := newService()
	h := newHandler(svc)

	// 2. 注册路由
	templates := routerGroup.Group("/install")
	templates.Post("/sqlite", h.InstallBySqlite)
}
