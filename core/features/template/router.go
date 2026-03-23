package template

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// Setup 是 user 模块暴露给全局的唯一启动入口
func Setup(routerGroup fiber.Router, db *gorm.DB) {

	// 初始化模板表
	InitTemplate(db)

	// 1. 实例化依赖
	repo := newRepository(db)
	svc := newService(repo)
	h := newHandler(svc)

	// 2. 注册路由
	templates := routerGroup.Group("/template")
	templates.Get("/", h.Take)
	templates.Post("/", h.Create)
	templates.Patch("/", h.Update)
	templates.Delete("/", h.Delete)
}
