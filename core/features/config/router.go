package config

import (
	"log"
	"one-dock/app/config"
	"one-dock/core/middleware"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// Setup 是 user 模块暴露给全局的唯一启动入口
func Setup(routerGroup fiber.Router, db *gorm.DB, cfg *config.Cfg) {

	// 初始化模板配置项表
	err := InitConfig(db)
	if err != nil {
		log.Fatal(err)
	}

	// 1. 实例化依赖
	repo := newRepository(db)
	svc := newService(repo)
	h := newHandler(svc)

	// 2. 注册路由
	templates := routerGroup.Group("/config")
	templates.Get("/", h.Take)
	templates.Post("/", middleware.AdminMiddleware(cfg), h.Create)
	templates.Patch("/", middleware.AdminMiddleware(cfg), h.Update)
	templates.Delete("/", middleware.AdminMiddleware(cfg), h.Delete)
}
