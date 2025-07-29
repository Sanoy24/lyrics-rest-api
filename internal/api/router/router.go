package router

import (
	"github.com/Sanoy24/lyrics-rest-api/internal/api/handlers"
	"github.com/Sanoy24/lyrics-rest-api/internal/api/middleware"
	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/artist"
	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/user"
	"github.com/Sanoy24/lyrics-rest-api/internal/api/services"
	"github.com/Sanoy24/lyrics-rest-api/internal/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, logger *zap.Logger, cfg *config.Config) *gin.Engine {
	// Default router
	router := gin.Default()
	// Intialize repo
	userRepo := user.NewUserRepo(db, logger)
	artistRepo := artist.NewArtistRepo(db, logger)

	// Initialize services
	authService := services.NewAuthService(userRepo, cfg.JWT.Secret, cfg.JWT.ExpireIn, logger)
	userService := services.NewUserService(userRepo, logger)
	artistService := services.NewArtistService(artistRepo, logger)

	// Initialize Handler
	authHandler := handlers.NewAuthHandler(authService, logger)
	userHandler := handlers.NewUserHandler(userService, logger)
	artistHandler := handlers.NewArtistHandler(artistService, logger)

	healthCheck := handlers.NewHealthHandler(logger)

	router.GET("/health", healthCheck.HealthCheck)

	// Middlewares
	router.Use(middleware.LoggerMiddleware(logger))
	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}
	}

	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware(cfg))
	{
		user := protected.Group("/users")
		{
			user.GET("/me", userHandler.GetCurrentUser)
			user.GET("/:id", userHandler.GetPublicUser)
			user.PUT("/me", userHandler.UpdateUser)
		}
		artist := protected.Group("/artists")
		{
			artist.POST("", artistHandler.CreateArtist)
		}
	}

	return router
}
