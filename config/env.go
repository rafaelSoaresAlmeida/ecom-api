package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DbUser     string
	DbPassword string
	DbAddress  string
	DbName     string
	JwtExpirationInSeconds int64
	JwtSecret string
}

var Envs = initConfig()

func initConfig() Config {

	godotenv.Load()


	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DbUser:     getEnv("DB_USER", "root"),
		DbPassword: getEnv("DB_PASSWORD", "password"),
		DbAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DbName:     getEnv("DB_NAME", "ecom"),
		JwtSecret:              getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
		JwtExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600 * 24 *7),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)

		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}