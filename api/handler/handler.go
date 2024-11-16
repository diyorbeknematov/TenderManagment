package handler

import (
	"log/slog"
	"tender/service"
)

type Handler struct{
	Service service.Service
	Log *slog.Logger
}

func NewHandler(service service.Service, logger *slog.Logger)*Handler{
	return &Handler{
		Service: service,
		Log: logger,
	}
}