package storage

import (
	"log"
	"one-dock/app/config"
	"one-dock/core/features/storage/handler"
	"one-dock/core/features/storage/models"
	"one-dock/core/features/storage/repository"
	"one-dock/core/features/storage/services"
	"one-dock/core/middleware"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func Setup(routerGroup fiber.Router, db *gorm.DB, cfg *config.Cfg) {

	// 初始化存储模型
	err := models.AutoMigrateStorageModel(db)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	svc := services.NewService(repo)
	h := handler.NewHandler(svc)

	router := routerGroup.Group("/storage")

	{
		manageRouter := router.Group("", middleware.AdminMiddleware(cfg))
		manageRouter.Get("/list", h.Manage.List)
		manageRouter.Get("/stats", h.Manage.Stats)
		manageRouter.Get("/download/:id", h.Manage.Download)
		manageRouter.Post("/folder", h.Manage.CreateFolder)
		manageRouter.Put("/rename", h.Manage.Rename)
		manageRouter.Delete("", h.Manage.Delete)

		chunkUploadRouter := router.Group("/chunk", middleware.AdminMiddleware(cfg))
		chunkUploadRouter.Post("/init", h.Upload.InitChunkUpload)
		chunkUploadRouter.Post("/upload", h.Upload.UploadChunk)
		chunkUploadRouter.Post("/merge", h.Upload.MergeChunk)
	}

}
