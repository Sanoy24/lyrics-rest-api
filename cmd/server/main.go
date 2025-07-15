package main

import (
	"github.com/Sanoy24/lyrics-rest-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

type HealthCheckResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func main() {

	router := gin.Default()
	healthHandler := handlers.NewHealthHandler()

	router.GET("/health", healthHandler.HealthCheck)

	router.Run(":8000")

}
