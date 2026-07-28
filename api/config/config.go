package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	AppName string
	AppHost string
	AppPort string
)

var (
	DBEngine string
	DBName   string
	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
)

var (
	SMTPHost string
	SMTPPort string
	SMTPUser string
	SMTPPass string
	MailFrom string
)

var (
	BasicAuthUser string
	BasicAuthPass string
)

var (
	AccessTokenSecret          string
	RefreshTokenSecret         string
	AccessTokenExpiresSeconds  int
	RefreshTokenExpiresSeconds int

	CookieAccessSecure    bool
	CookieRefreshSecure   bool
	CookieAccessHttpOnly  bool
	CookieRefreshHttpOnly bool
)

var (
	LogLevel       string
	FrontendOrigin string
)

var (
	CodeExecutorHost string
)

func init() {
	envFile, err := findEnvFile(".env")
	if err != nil {
		log.Panic("Failed to find environment variables:", err)
	}
	err = godotenv.Load(envFile)
	if err != nil {
		log.Panic("Failed to load environment variables:", err)
	}

	AppName = getEnv("APP_NAME")
	AppHost = getEnv("APP_HOST", "localhost")
	AppPort = getEnv("APP_PORT", "8000")

	DBEngine = getEnv("DB_ENGINE")
	DBName = getEnv("DB_NAME")
	DBHost = getEnv("DB_HOST")
	DBPort = getEnv("DB_PORT")
	DBUser = getEnv("DB_USER")
	DBPass = getEnv("DB_PASSWORD")

	SMTPHost = getEnv("SMTP_HOST")
	SMTPPort = getEnv("SMTP_PORT")
	SMTPUser = getEnv("SMTP_USER")
	SMTPPass = getEnv("SMTP_PASSWORD")
	MailFrom = getEnv("MAIL_FROM")

	BasicAuthUser = getEnv("BASIC_AUTH_USER")
	BasicAuthPass = getEnv("BASIC_AUTH_PASSWORD")

	AccessTokenSecret = getEnv("ACCESS_TOKEN_SECRET", "secret")
	AccessTokenExpiresSeconds, err = strconv.Atoi(getEnv("ACCESS_TOKEN_EXPIRES_SECONDS", "900"))
	if err != nil {
		log.Fatalf("unable to convert ACCESS_TOKEN_EXPIRES_SECONDS from environment to integer: %v", err)
	}

	RefreshTokenSecret = getEnv("REFRESH_TOKEN_SECRET", "secret")
	RefreshTokenExpiresSeconds, err = strconv.Atoi(getEnv("REFRESH_TOKEN_EXPIRES_SECONDS", "86400"))
	if err != nil {
		log.Fatalf("unable to convert REFRESH_TOKEN_EXPIRES_SECONDS from environment to integer: %v", err)
	}

	CookieAccessSecure, err = strconv.ParseBool(getEnv("COOKIE_ACCESS_SECURE", "true"))
	if err != nil {
		log.Fatalf("unable to convert COOKIE_ACCESS_SECURE from environment to boolean: %v", err)
	}

	CookieRefreshSecure, err = strconv.ParseBool(getEnv("COOKIE_REFRESH_SECURE", "true"))
	if err != nil {
		log.Fatalf("unable to convert COOKIE_REFRESH_SECURE from environment to boolean: %v", err)
	}

	CookieAccessHttpOnly, err = strconv.ParseBool(getEnv("COOKIE_ACCESS_HTTPONLY", "true"))
	if err != nil {
		log.Fatalf("unable to convert COOKIE_ACCESS_HTTPONLY from environment to boolean: %v", err)
	}

	CookieRefreshHttpOnly, err = strconv.ParseBool(getEnv("COOKIE_REFRESH_HTTPONLY", "true"))
	if err != nil {
		log.Fatalf("unable to convert COOKIE_REFRESH_HTTPONLY from environment to boolean: %v", err)
	}
	LogLevel = getEnv("LOG_LEVEL", "INFO")
	FrontendOrigin = getEnv("FRONTEND_ORIGIN", "http://localhost:3000")

	CodeExecutorHost = getEnv("CODE_EXECUTOR_HOST")
}

func getEnv(key string, defaultValue ...string) string {
	value := os.Getenv(key)
	if value == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return value
}

func findEnvFile(fileName string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		candidate := filepath.Join(dir, "config", "env", fileName)
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		} else if !os.IsNotExist(err) {
			return "", err
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("config/env/%s not found", fileName)
}
