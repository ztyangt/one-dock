package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"one-dock/app/config"
	"one-dock/app/logger"
	"one-dock/app/middleware"
	"one-dock/app/response"
	"one-dock/pkgs/console"
	"one-dock/public"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/static"
	"gorm.io/gorm"
)

type Server struct {
	App *fiber.App
	cfg *config.Cfg
	log *logger.LogFactor
	db  *gorm.DB
}

// NewServer 创建服务实例
func NewServer(cfg *config.Cfg, logFactor *logger.LogFactor, db *gorm.DB) *Server {
	server := &Server{}
	server.cfg = cfg
	server.log = logFactor
	server.db = db
	app := fiber.New(fiber.Config{
		BodyLimit:    1024 * 1024 * 10, // 10MB
		ErrorHandler: response.ErrorHandler(server.log),
	})

	// 静态文件服务
	app.Get("/*", static.New("/", static.Config{
		FS:     public.HomeFS,
		Browse: false,
	}))

	// 全局中间件
	app.Use(
		recover.New(),
		middleware.FaviconMiddleware(),
		middleware.HeaderMiddleware(),
		middleware.LoggerMiddleware(server.log),
		middleware.CorsMiddleware(),
		middleware.GetUserMiddleware(cfg),
	)

	server.App = app

	app.Hooks().OnPostStartupMessage(func(sm *fiber.PostStartupMessageData) error {
		console.Infof("服务启动成功: http://localhost:%s\n", server.cfg.App.Port)
		return nil
	})

	return server

}

func (s *Server) Start() {
	// 注册路由
	RegisterRoutes(s.App, s.db, s.cfg)

	if err := s.App.Listen(fmt.Sprintf(":%s", s.cfg.App.Port), fiber.ListenConfig{
		DisableStartupMessage: true,
	}); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}

}

func (s *Server) Restart(cfg *config.Cfg, logFactor *logger.LogFactor, db *gorm.DB) {
	console.Info("重启服务中...")
	s.cfg = cfg
	s.log = logFactor
	s.db = db
	if s.App != nil {
		s.App.Shutdown()
	}
	go s.Start()
}

func (s *Server) WaitForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	console.Info("准备关闭服务...")
	// s.App.Shutdown()
	console.Info("服务已关闭")
}
