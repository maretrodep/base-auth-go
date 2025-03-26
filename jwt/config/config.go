package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port string
}

type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

type AuthConfig struct {
	JWTSecret string
}

type Config struct {
	Server ServerConfig
	DB     DBConfig
	Auth   AuthConfig
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found.")
	}

	cfg := &Config{
		Server: ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
		},
		DB: DBConfig{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
		},
		Auth: AuthConfig{
			JWTSecret: os.Getenv("JWT_SECRET"),
		},
	}
	return cfg, nil
}
