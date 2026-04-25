package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/prabowoteguh/belajar-vibe-code/pkg/logger"
)

type Config struct {
	AppPort string
	AppEnv  string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int

	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int

	JiraHost  string
	JiraUser  string
	JiraToken string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		logger.Info("No .env file found, using system environment variables")
	}

	return &Config{
		AppPort: getEnv("APP_PORT", "8080"),
		AppEnv:  getEnv("APP_ENV", "development"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "1433"),
		DBUser:     getEnv("DB_USER", "sa"),
		DBPassword: getEnv("DB_PASSWORD", "Password123"),
		DBName:     getEnv("DB_NAME", "master"),

		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvAsInt("REDIS_DB", 0),

		ReadTimeout:  getEnvAsInt("READ_TIMEOUT", 15),
		WriteTimeout: getEnvAsInt("WRITE_TIMEOUT", 15),
		IdleTimeout:  getEnvAsInt("IDLE_TIMEOUT", 60),

		JiraHost:  getEnv("JIRA_HOST", "https://jira.bri.co.id"),
		JiraUser:  getEnv("JIRA_USER", "00345834"),
		JiraToken: getEnv("JIRA_TOKEN", "P@ssw0rdBrilian"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(name string, defaultValue int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
