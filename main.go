package main

import (
	"one-dock/app/config"
	"one-dock/app/database"
	"one-dock/app/logger"
	"one-dock/app/server"

	"github.com/fsnotify/fsnotify"
	"gorm.io/gorm"
)

func main() {

	// 1. 初始化配置
	cfg := config.Init()

	// 2. 初始化核心依赖
	log, db := setupDependencies(cfg)

	// 3. 启动服务
	srv := server.NewServer(cfg, log, db)
	go srv.Start()

	// 4. 监听配置变更
	cfg.Viper.OnConfigChange(func(in fsnotify.Event) {

		if err := cfg.Reload(); err == nil {
			// 重新加载依赖并重启服务
			newLog, newDB := setupDependencies(cfg)
			srv.Restart(cfg, newLog, newDB)
		}
	})

	// 5. 等待优雅退出
	srv.WaitForShutdown()
}

// setupDependencies 封装依赖初始化的逻辑（DRY原则）
func setupDependencies(cfg *config.Cfg) (*logger.LogFactor, *gorm.DB) {
	log := logger.Init(cfg.Log)
	db := database.Init(cfg.DB)

	return log, db
}
