package main

import (
	"log"
	"strconv"

	"github.com/Sanoy24/lyrics-rest-api/internal/config"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"github.com/Sanoy24/lyrics-rest-api/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {

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

	if err = database.Seed(db); err != nil {
		log.Fatal("Failed to seed database:", err)
	}

	router := gin.Default()
	// TODO: Add routes here

	router.Run(":" + cfg.Server.Port) // Use port from config

}
