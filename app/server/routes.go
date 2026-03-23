package server

import (
	"one-dock/app/config"
	configuration "one-dock/core/features/config"
	"one-dock/core/features/system"
	"one-dock/core/features/user"
	"one-dock/pkgs/console"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func RegisterRoutes(router fiber.Router, db *gorm.DB, cfg *config.Cfg) {

	apiRouter := router.Group("/api")

	system.Setup(apiRouter)
	configuration.Setup(apiRouter, db, cfg)
	user.Setup(apiRouter, db, cfg)

	console.Info("路由挂载成功！")
}
