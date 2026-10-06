package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Port           string
	Environment    string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBSSLMode      string
	RedisHost      string
	RedisPort      string
	JWTSecret      string
	StoragePublic  string
	StoragePrivate string
}

func LoadConfig() *Config {
	return &Config{
		Port:           getEnv("PORT", "8080"),
		Environment:    getEnv("NODE_ENV", "production"),
		DBHost:         getEnv("DB_HOST", "127.0.0.1"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "meldir_user"),
		DBPassword:     getEnv("DB_PASSWORD", "PasswordKuatAnda2026!"),
		DBName:         getEnv("DB_NAME", "meldir_db"),
		DBSSLMode:      getEnv("DB_SSLMODE", getEnv("SSL_MODE", "disable")),
		RedisHost:      getEnv("REDIS_HOST", "127.0.0.1"),
		RedisPort:      getEnv("REDIS_PORT", "6379"),
		JWTSecret:      getEnv("JWT_SECRET", "meldir-production-secure-jwt-secret-key-2026"),
		StoragePublic:  getEnv("STORAGE_PUBLIC", "/home/meldir.id/storage/public"),
		StoragePrivate: getEnv("STORAGE_PRIVATE", "/home/meldir.id/storage/private"),
	}
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBName, c.DBSSLMode,
	)
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		val = strings.TrimSpace(val)
		val = strings.Trim(val, "\"'`")
		return val
	}
	return defaultVal
}
