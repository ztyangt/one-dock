package handler

import "one-dock/core/features/storage/services"

type Handler struct {
	Upload UploadHandler
}

func NewHandler(svc services.Service) Handler {
	return Handler{Upload: newUploadHandler(svc)}
}
