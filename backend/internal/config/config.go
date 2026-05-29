package config

import "os"

type Config struct {
	Env            string
	Port           string
	DBUrl          string
	JWTSecret      string
	AppEnv         string
	DBMigrationUrl string
}

func Load() *Config {
	return &Config{
		Env:            getEnv("APP_ENV", "development"),
		Port:           getEnv("API_PORT", "8080"),
		DBUrl:          getEnv("DATABASE_URL", ""),
		JWTSecret:      getEnv("JWT_SECRET", ""),
		AppEnv:         getEnv("APP_ENV", "development"),
		DBMigrationUrl: getEnv("DATABASE_MIGRATION_URL", ""),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
