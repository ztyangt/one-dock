package user

import (
	"one-dock/app/config"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// Setup 是 user 模块暴露给全局的唯一启动入口
func Setup(routerGroup fiber.Router, db *gorm.DB, cfg *config.Cfg) {

	// 初始化用户表
	InitUser(db)

	// 1. 实例化依赖
	repo := newRepository(db)
	svc := newService(repo, cfg)
	h := newHandler(svc, cfg)

	// 2. 注册路由
	users := routerGroup.Group("/user")
	users.Get("/", h.getUser)
	users.Post("/login", h.login)
}
