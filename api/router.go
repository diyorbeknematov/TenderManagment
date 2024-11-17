package api

import (
	"log/slog"
	_ "tender/api/docs"
	"tender/api/handler"
	"tender/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
func Router(service service.Service, logger *slog.Logger) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h := handler.NewHandler(service, logger)

	tender := router.Group("/tenders")
	{
		tender.POST("", h.CreateTender)
		tender.GET("", h.GetAllTenders)
		tender.PUT("/:id", h.UpdateTender)
		tender.DELETE("/:id", h.DeleteTender)
		tender.GET("/:id/bids", h.GetTenderBids)
		tender.POST("/:id/bids", h.SubmitBit)
		tender.POST("/:id/award/:bid_id", h.AwardTender)
		tender.POST("/:id/bids", h.CreateBid)
		tender.GET("/:id/bids", h.GetBidsOfTender)
		tender.GET("/all", h.GetTendersByFilters)
	}

	user := router.Group("/users")
	{
		user.GET("/:id/bids", h.GetMyBidHistory)
	}
	return router
}
