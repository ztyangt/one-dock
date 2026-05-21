package server

import (
	"io/fs"
	"log"

	"one-dock/app/config"
	configuration "one-dock/core/features/config"
	"one-dock/core/features/storage"
	"one-dock/core/features/system"
	"one-dock/core/features/user"
	"one-dock/pkgs/console"
	"one-dock/public"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB, cfg *config.Cfg) {
	// ==========================================
	// 1. API 路由 (必须在最前面)
	// ==========================================
	apiRouter := app.Group("/api")

	system.Setup(apiRouter, cfg)
	configuration.Setup(apiRouter, db, cfg)
	user.Setup(apiRouter, db, cfg)
	storage.Setup(apiRouter, db, cfg)

	// ==========================================
	// 2. 独立静态资源目录 (/sources)
	// ==========================================
	sourceFS, err := fs.Sub(public.SourcesFS, "sources")
	if err != nil {
		log.Fatalf("无法剥离 sources 目录前缀: %v", err)
	}

	app.Get("/sources/*", static.New("", static.Config{
		FS: sourceFS,
	}))
	// 拦截器：如果 /sources/xxx 没找到文件，直接返回 404，防止漏给最底下的 SPA index.html
	app.All("/sources/*", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	// ==========================================
	// 3. SPA 单页应用资源 (/js, /css, /img 等)
	// ==========================================
	distFS, err := fs.Sub(public.DistFS, "dist")
	if err != nil {
		log.Fatalf("无法剥离 dist 目录前缀: %v", err)
	}

	// 将 index.html 预加载到内存中 (性能优化)
	indexHTML, err := fs.ReadFile(distFS, "index.html")
	if err != nil {
		log.Fatalf("无法读取 index.html: %v", err)
	}

	// 静态文件代理
	// Fiber v3 使用 middleware/static 来处理静态文件
	// 如果匹配不到文件(比如前端路由 /user)，它会自动调用 c.Next() 传递给下一个路由
	app.Get("/*", static.New("", static.Config{
		FS: distFS,
	}))

	// ==========================================
	// 4. SPA 路由兜底 (Catch-all)
	// ==========================================
	// 当上述静态文件中间件找不到对应的文件时，就会来到这里
	// 无论用户请求 /about 还是 /users/1，统一返回 index.html
	app.Get("/*", func(c fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
		return c.Send(indexHTML) // c.Send 直接支持 []byte，性能最佳
	})

	console.Info("路由挂载成功！")
}
