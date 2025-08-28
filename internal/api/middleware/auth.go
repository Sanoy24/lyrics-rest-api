package middleware

import (
	"net/http"
	"slices"
	"strings"

	"github.com/Sanoy24/lyrics-rest-api/internal/config"
	"github.com/Sanoy24/lyrics-rest-api/pkg/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
		zap.L().Info("Authenticated user", zap.String("email", claims.Email))
		ctx.Set("user_id", claims.UserID)
		ctx.Set("email", claims.Email)
		ctx.Set("role", claims.Role)
		ctx.Set("permissions", claims.Permission)
		ctx.Next()
	}
}

func RequirePermission(required string) gin.HandlerFunc {
	return func(c *gin.Context) {
		val, ok := c.Get("permissions")
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No permissions found", "data": val})
			return
		}

		if slices.Contains(val.([]string), required) {
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
	}
}
