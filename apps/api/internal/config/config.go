package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port             int
	AppEnv           string
	DatabaseURL      string
	FrontendURL      string
	SessionSecret    string
	SessionMaxAge    time.Duration
	GoogleClientID   string
	GoogleSecret     string
	GoogleRedirect   string
	YouTubeAPIKey    string
}

func Load() *Config {
	return &Config{
		Port:          getEnvInt("PORT", 8080),
		AppEnv:        getEnv("APP_ENV", "development"),
		DatabaseURL:   getEnv("DATABASE_URL", "postgres://tubedex:tubedex@localhost:5432/tubedex?sslmode=disable"),
		FrontendURL:   getEnv("FRONTEND_URL", "http://localhost:5173"),
		SessionSecret: getEnv("SESSION_SECRET", "dev-secret-change-in-production"),
		SessionMaxAge: 30 * 24 * time.Hour,
		GoogleClientID:   getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleSecret:     getEnv("GOOGLE_CLIENT_SECRET", ""),
		GoogleRedirect:   getEnv("GOOGLE_REDIRECT_URL", "http://localhost:8080/api/v1/auth/callback"),
		YouTubeAPIKey:    getEnv("YOUTUBE_API_KEY", ""),
	}
}

func (c *Config) IsDevelopment() bool {
	return c.AppEnv == "development"
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if val := os.Getenv(key); val != "" {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return fallback
}
