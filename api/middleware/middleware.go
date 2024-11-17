package middleware

import (
	"fmt"
	"log/slog"
	"tender/api/token"
	"tender/model"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			logger.Error("No Authorization header")
			ctx.JSON(model.ErrUnauthorized.Code, model.ErrUnauthorized)
			ctx.Abort()
			return
		}

		// JWT tokenni tekshirish va tasdiqlash
		claims, err := token.ExtractAccessClaims(tokenString)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to extract: %v", err))
			ctx.JSON(model.ErrUnauthorized.Code, model.ErrUnauthorized)
			ctx.Abort()
			return
		}

		// Foydalanuvchi ma'lumotlarini context ga qo'shish
		ctx.Set("UserID", claims.ID)
		ctx.Set("Username", claims.Username)
		ctx.Set("UserRole", claims.Role)

		ctx.Next()
	}
}

func AuthorizeMiddleware(logger *slog.Logger, enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole := ctx.GetString("UserRole")
		if userRole == "" {
			ctx.JSON(model.ErrUnauthorized.Code, model.ErrUnauthorized)
			ctx.Abort()
			return
		}

		ok, err := enforcer.Enforce(userRole, ctx.FullPath(), ctx.Request.Method)
		if err != nil {
			ctx.JSON(model.ErrInternalServerError.Code, model.ErrInternalServerError)
			ctx.Abort()
			return
		}

		if !ok {
			ctx.JSON(model.ErrForbidden.Code, model.ErrForbidden)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func LogMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger.Info("Request received",
			slog.String("method", ctx.Request.Method),
			slog.String("path", ctx.Request.URL.Path),
		)

		ctx.Next()

		logger.Info("Response sent",
			slog.Int("status", ctx.Writer.Status()),
		)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}