package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load() // 忽略文件不存在错误

	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "db_pg"),
		DBPassword: getEnv("DB_PASSWORD", "admin123"),
		DBName:     getEnv("DB_NAME", "admin"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}, nil
}

func getEnv(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}
