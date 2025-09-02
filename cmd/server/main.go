package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/router"
	"github.com/Sanoy24/lyrics-rest-api/internal/config"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"github.com/Sanoy24/lyrics-rest-api/pkg/database"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	cfg, err := config.LoadConfig()
	logger := setupLogger(cfg)

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
		if err := database.SetupTsVector(db); err != nil {
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

	router := router.SetupRouter(db, logger, cfg)

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
		log.Fatal("Server forced to shutdown", err)
	}

	log.Println("Server exited gracefully.")
}

func setupLogger(cfg *config.Config) *zap.Logger {
	// Define log file path
	logDir := "logs"
	logFile := filepath.Join(logDir, "app.log")

	// Create logs directory if it doesn't exist
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatal("Failed to create log directory:", err)
	}

	// Configure file sink with rotation
	fileSink := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    10, // MB
		MaxBackups: 5,
		MaxAge:     30, // days
		Compress:   true,
	}

	// Base encoder config
	baseEncoderCfg := zap.NewProductionEncoderConfig()
	baseEncoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	baseEncoderCfg.EncodeCaller = zapcore.ShortCallerEncoder

	// Console encoder config
	consoleEncoderCfg := baseEncoderCfg
	if cfg.Server.Env == "development" {
		consoleEncoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder // with colors
	} else {
		consoleEncoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder // no colors
	}
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderCfg)

	// File encoder config (ALWAYS plain JSON, no colors)
	fileEncoderCfg := baseEncoderCfg
	fileEncoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	fileEncoder := zapcore.NewJSONEncoder(fileEncoderCfg)

	// Set up core with multiple outputs
	var core zapcore.Core
	if cfg.Server.Env == "development" {
		core = zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
			zapcore.NewCore(fileEncoder, zapcore.AddSync(fileSink), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
			zapcore.NewCore(fileEncoder, zapcore.AddSync(fileSink), zapcore.InfoLevel),
		)
	}

	// Create logger
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger
}
