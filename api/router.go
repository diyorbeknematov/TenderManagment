package api

import (
	"tender/api/handler"

	"github.com/gin-gonic/gin"
)

func Router(h *handler.Handler) *gin.Engine {
	router := gin.Default()

	return router
}
