package config

import "os"

type Config struct {
	Env       string
	Port      string
	DBUrl     string
	JWTSecret string
}

func Load() *Config {
	return &Config{
		Env:       getEnv("APP_ENV", "development"),
		Port:      getEnv("API_PORT", "8080"),
		DBUrl:     getEnv("DATABASE_URL", ""),
		JWTSecret: getEnv("JWT_SECRET", ""),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
