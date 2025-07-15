package database

import (
	"fmt"

	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(dbParam *models.PostgresParam) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbParam.DB_HOST, dbParam.DB_USER, dbParam.DB_PASSWORD, dbParam.DB_NAME, dbParam.DB_PORT,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

}
