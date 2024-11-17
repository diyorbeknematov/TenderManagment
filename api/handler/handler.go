package handler

import (
	"log/slog"
	"tender/service"
	"tender/storage"
)

type Handler struct{
	Service service.Service
	Log *slog.Logger
	Storage storage.Storage
}

func NewHandler(service service.Service, logger *slog.Logger)*Handler{
	return &Handler{
		Service: service,
		Log: logger,
	}
}