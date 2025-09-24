package router

import (
	_ "github.com/Sanoy24/lyrics-rest-api/docs"
	"github.com/Sanoy24/lyrics-rest-api/internal/api/handlers"
	"github.com/Sanoy24/lyrics-rest-api/internal/api/middleware"
	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/annotation"
	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/artist"
	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/song"
	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/user"
	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/vote"
	"github.com/Sanoy24/lyrics-rest-api/internal/api/services"
	"github.com/Sanoy24/lyrics-rest-api/internal/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, logger *zap.Logger, cfg *config.Config) *gin.Engine {
	// Default router
	router := gin.Default()
	// Intialize repo
	userRepo := user.NewUserRepo(db, logger)
	artistRepo := artist.NewArtistRepo(db, logger)
	songRepo := song.NewSongRepo(db, logger)
	annotationRepo := annotation.NewAnnotationRepo(db, logger)
	voteRepo := vote.NewVoteRepo(db, logger)

	// Initialize services
	authService := services.NewAuthService(userRepo, cfg.JWT.Secret, cfg.JWT.ExpireIn, logger)
	userService := services.NewUserService(userRepo, logger)
	artistService := services.NewArtistService(artistRepo, logger)
	songService := services.NewSongService(songRepo, artistRepo, logger)
	annotationService := services.NewAnnotationService(annotationRepo, songRepo, logger)
	voteService := services.NewVoteService(voteRepo, songRepo, annotationRepo, logger)
	uploadService := services.NewUploadService(cfg)

	// Initialize Handler
	authHandler := handlers.NewAuthHandler(authService, logger)
	userHandler := handlers.NewUserHandler(userService, logger)
	artistHandler := handlers.NewArtistHandler(artistService, logger)
	songHandler := handlers.NewSongHandler(songService, logger)
	annotationHandler := handlers.NewAnnotationHandler(annotationService, logger)
	voteHandler := handlers.NewVoteHandler(voteService, logger)
	uploadHander := handlers.NewUploadHandler(uploadService)

	healthCheck := handlers.NewHealthHandler(logger)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/health", healthCheck.HealthCheck)

	// Middlewares
	router.Use(middleware.LoggerMiddleware(logger))
	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		auth.Use(middleware.RateLimiterMiddleware(5, 10))
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
			user.POST("/upload", uploadHander.UploadImageHandler)
		}
		artist := protected.Group("/artists")
		{
			artist.POST("", artistHandler.CreateArtist)
			artist.GET("", artistHandler.GetAllArtists)
			artist.GET("/:id", artistHandler.GetArtistByID)
			artist.PUT("/:id", artistHandler.UpdateArtist)
			artist.DELETE("/:id", artistHandler.DeleteArtist)
			// get all artists")
			// get single artist
			// update artist
		}
		song := protected.Group("/song")
		{
			song.POST("", songHandler.CreateSong)
			song.GET("", songHandler.GetAllSongs)
			song.GET("/:id", songHandler.GetSongById)
			song.GET("/slug/:slug", songHandler.GetSongBySlug)
			song.PUT("/:id", songHandler.UpdateSong)
			song.DELETE("/:id", songHandler.DeleteSong)
			song.GET("/search", songHandler.SearchSongs)
		}

		annotation := protected.Group("/annotation")
		{
			annotation.POST("/:song_id", annotationHandler.CreateAnnotation)
			annotation.GET("/:song_id", annotationHandler.GetAnnotationsBySongID)
			annotation.PUT("/songs/:song_id/annotations/:annotation_id", annotationHandler.UpdateAnnotation)
			annotation.DELETE("/:id", annotationHandler.DeleteAnnotation)
		}

		vote := protected.Group("/vote")
		{
			vote.POST("", voteHandler.CastVote)
			vote.POST("/:id", voteHandler.UpdateVoteScore)
		}

	}

	return router
}
