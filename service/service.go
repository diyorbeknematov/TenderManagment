package service

import (
	"log/slog"
	"tender/pkg/websocket"
	"tender/storage"
)

type Service struct {
	Storage          storage.Storage
	webSocketManager *websocket.Manager
	Log              *slog.Logger
}

func NewService(storage storage.Storage, logger *slog.Logger, wsManager *websocket.Manager) Service {
	return Service{
		Storage:          storage,
		webSocketManager: wsManager,
		Log:              logger,
	}
}
