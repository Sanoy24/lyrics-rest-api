package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port string
	Env  string
}

type DatabaseConfig struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
}

type JWTConfig struct {
	Secret   string
	ExpireIn time.Duration
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
	}

	expiresIn, _ := time.ParseDuration(getEnv("JWT_EXPIRES_IN", "30m"))
	fmt.Println("EXPIRES IN", expiresIn)

	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Env:  getEnv("ENV", "development"),
		}, Database: DatabaseConfig{
			DBHost:     getEnv("DB_HOST", "localhost"),
			DBUser:     getEnv("DB_USER", "postgres"),
			DBPassword: getEnv("DB_PASSWORD", "postgres"),
			DBName:     getEnv("DB_NAME", "lyrics"),
			DBPort:     getEnv("DB_PORT", "5432"),
		}, JWT: JWTConfig{
			Secret:   getEnv("JWT_SECRET", "secret"),
			ExpireIn: expiresIn,
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
