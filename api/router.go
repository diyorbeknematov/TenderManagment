package api

import (
	"tender/api/handler"

	"github.com/gin-gonic/gin"
)

// @title 						TENDER MANAGMENT API
// @version 					0.1
// @description 				This is a sample API.
// @schemes 					http https
// @BasePath 					/
// @consumes 					application/json
// @produces 					application/json
// @securityDefinitions.apiKey 	Bearer
// @in 							header
// @name 						Authorization
// @swagger:meta
func Router(h *handler.Handler) *gin.Engine {
	router := gin.Default()

	return router
}
