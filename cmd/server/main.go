package main

import (
	"log"

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
	err = database.InitDB(&models.PostgresParam{
		DB_HOST:     cfg.Database.DBHost,
		DB_USER:     cfg.Database.DBUser,
		DB_PASSWORD: cfg.Database.DBPassword,
		DB_NAME:     cfg.Database.DBName,
		DB_PORT:     cfg.Database.DBPort,
	})
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer database.CLoseDB()

	db := database.GetDB()

	router := gin.Default()
	router.Run(":8000")

}
