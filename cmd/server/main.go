package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/handlers"
	"github.com/Sanoy24/lyrics-rest-api/internal/api/middleware"
	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/user"
	"github.com/Sanoy24/lyrics-rest-api/internal/api/services"
	"github.com/Sanoy24/lyrics-rest-api/internal/config"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"github.com/Sanoy24/lyrics-rest-api/pkg/database"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig()
	logger := setupLogger()

	runMigrations := flag.Bool("migrate", false, "Run database migrations and exit")
	runSeeder := flag.Bool("seed", false, "Run database seeder and exit")
	flag.Parse()

	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	dbPort, err := strconv.Atoi(cfg.Database.DBPort)
	if err != nil {
		log.Fatalf("Invalid DB_PORT in config: %v", err)
	}

	err = database.InitDB(&models.PostgresParam{
		DB_HOST:     cfg.Database.DBHost,
		DB_USER:     cfg.Database.DBUser,
		DB_PASSWORD: cfg.Database.DBPassword,
		DB_NAME:     cfg.Database.DBName,
		DB_PORT:     dbPort,
	})
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer database.CloseDB()
	db := database.GetDB()

	if *runMigrations {
		log.Println("Running database migrations...")
		if err := database.Migrate(db); err != nil {
			log.Fatal("Failed to run migrations:", err)
		}
		log.Println("Migrations completed successfully.")
		return
	}
	// Conditionally run seeder
	if *runSeeder {
		log.Println("Starting database seeding...")
		if err := database.Seed(db); err != nil {
			log.Fatalf("Database seeding failed: %v", err)
		}

		return
	}

	router := setupRouter(db, logger, cfg)

	server := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server:", err)
		}
	}()

	gracefulShutdown(server)

}

func gracefulShutdown(server *http.Server) {
	// Implement graceful shutdown logic here
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited gracefully.")
}

func setupRouter(db *gorm.DB, logger *zap.Logger, cfg *config.Config) *gin.Engine {
	router := gin.Default()
	userRepo := user.NewUserRepo(db, logger)
	authService := services.NewAuthService(userRepo, cfg.JWT.Secret, cfg.JWT.ExpireIn, logger)
	authHandler := handlers.NewAuthHandler(authService, logger)
	userService := services.NewUserService(userRepo, logger)
	userHandler := handlers.NewUserHandler(userService, logger)
	// healthCheck := handlers.NewHealthHandler()
	// router.GET("/health", healthCheck.HealthCheck)
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
	protected.Use(middleware.AuthMiddleware(cfg), middleware.RequirePermission("song:create"))
	{
		user := protected.Group("/users")
		{
			user.GET("/me", userHandler.GetCurrentUser)
			user.PUT("/me", nil)
		}
	}

	return router
}

func setupLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Failed to initialize logger:", err)
	}

	return logger

}
