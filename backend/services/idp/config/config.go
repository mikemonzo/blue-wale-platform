package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Port int
}

type DBConfig struct {
	User     string
	Password string
	DBName   string
	SSLMode  string
	Host     string
	Port     int
}

func NewConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnvAsInt("PORT", 8080),
		},
		DB: DBConfig{
			User:     getEnv("DB_USER", "bwadmin"),
			Password: getEnv("DB_PASSWORD", "bwpassword"),
			DBName:   getEnv("DB_NAME", "idp_db"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
			Host:     getEnv("DB_HOST", "postgres"),
			Port:     getEnvAsInt("DB_PORT", 5432),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return fallback
}
