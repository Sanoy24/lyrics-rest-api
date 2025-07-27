package middleware

import (
	"net/http"
	"strings"

	"github.com/Sanoy24/lyrics-rest-api/internal/config"
	"github.com/Sanoy24/lyrics-rest-api/pkg/util"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, util.Response{
				Status:  false,
				Message: "Unauthorized",
			})
			ctx.Abort()
			return
		}
		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, util.Response{
				Status:  false,
				Message: "Invalid Authorization Header",
			})
			ctx.Abort()
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := util.ValidateToken(token, cfg.JWT.Secret)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, util.Response{
				Status:  false,
				Message: "Invalid Token or expired token",
			})
			ctx.Abort()
			return
		}
		ctx.Set("user_id", claims.UserID)
		ctx.Set("username", claims.Email)
		ctx.Set("role", claims.Role)
		ctx.Next()
	}
}
