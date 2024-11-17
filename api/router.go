package api

import (
	"log/slog"
	_ "tender/api/docs"
	"tender/api/handler"
	"tender/api/middleware"
	"tender/service"
	"tender/storage"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Dependencies struct {
	Logger         *slog.Logger
	Enforcer       *casbin.Enforcer
	RateLimiter    middleware.RateLimiter
	ServiceManager service.Service
	Storage        storage.Storage
}

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
func Router(deps *Dependencies) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h := handler.NewHandler(deps.ServiceManager, deps.Logger, deps.Storage)

	router.POST("/register", h.RegistrationHandler)
	router.POST("/login", h.LoginHandler)

	tender := router.Group("/tenders")
	tender.Use(middleware.AuthMiddleware(deps.Logger))
	{
		tender.POST("", h.CreateTender)
		tender.GET("", h.GetAllTenders)
		tender.PUT("/:id", h.UpdateTender)
		tender.DELETE("/:id", h.DeleteTender)
		tender.GET("/:id/my/bids", h.GetTenderBids)
		tender.POST("/status_change/:id/bids", h.SubmitBit)
		tender.POST("/:id/award/:bid_id", h.AwardTender)
		tender.POST("/:id/bids", h.CreateBid)
		tender.GET("/:id/bids", h.GetBidsOfTender)
		tender.GET("/all", h.GetTendersByFilters)
	}

	router.GET("/ws/notifications", h.WebSocketNotifications)

	user := router.Group("/users")
	{
		user.GET("/:id/bids", h.GetMyBidHistory)
		user.GET("/:id/tenders", h.GetMyTenderHistory)
	}
	return router
}
