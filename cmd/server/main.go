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
	"github.com/Sanoy24/lyrics-rest-api/internal/config"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"github.com/Sanoy24/lyrics-rest-api/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {

	runMigrations := flag.Bool("migrate", false, "Run database migrations and exit")
	runSeeder := flag.Bool("seed", false, "Run database seeder and exit")
	flag.Parse()

	cfg, err := config.LoadConfig()

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

	router := setupRouter()

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

func setupRouter() *gin.Engine {
	router := gin.Default()
	healthCheck := handlers.NewHealthHandler()
	router.GET("/health", healthCheck.HealthCheck)
	router.Group("/api/v1")
	{
		// v1.GET("/health", func(c *gin.Context) {
		// 	c.JSON(http.StatusOK, gin.H{"status": "ok"})
		// })
	}

	return router
}
