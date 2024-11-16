package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	DB_HOST            string
	DB_PORT            string
	DB_USER            string
	DB_NAME            string
	DB_PASSWORD        string
	API_PORT           string
	ACCESS_SECRET_KEY  string
	REFRESH_SECRET_KEY string
}

func LoadConfig() Config {
	if err := godotenv.Load("./../../.env"); err != nil {
		log.Println("error loading .env file or not found", err)
	}

	config := Config{}

	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToString(coalesce("DB_PORT", "5432"))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "education_center"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "secret"))
	config.API_PORT = cast.ToString(coalesce("API_PORT", "8080"))
	config.ACCESS_SECRET_KEY = cast.ToString(coalesce("ACCESS_SECRET_KEY", "secret"))
	config.REFRESH_SECRET_KEY = cast.ToString(coalesce("REFRESH_SECRET_KEY", "secret"))
	return config
}

func coalesce(key string, defValue interface{}) interface{} {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defValue
}
