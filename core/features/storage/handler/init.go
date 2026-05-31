package handler

import "one-dock/core/features/storage/services"

type Handler struct {
	Upload UploadHandler
	Manage ManageHandler
}

func NewHandler(svc services.Service) Handler {
	return Handler{Upload: newUploadHandler(svc), Manage: newManageHandler(svc)}
}
