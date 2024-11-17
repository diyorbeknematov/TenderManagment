package service

import (
	"log/slog"
	"tender/storage"
)

type Service struct {
	Storage storage.Storage
	Log     *slog.Logger
}

func NewService(storage storage.Storage, logger *slog.Logger) Service {
	return Service{
		Storage: storage,
		Log:     logger,
	}
}
