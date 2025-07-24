package config

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    DBHost   string
    DBPort   string
    DBUser   string
    DBPass   string
    DBName   string
    SSLMode  string
    SecretKey string
}

func Load() *Config {
    if err := godotenv.Load(); err != nil {
        log.Println(".env não encontrado, usando variáveis de ambiente do sistema")
    }

    return &Config{
        DBHost:    os.Getenv("DB_HOST"),
        DBPort:    os.Getenv("DB_PORT"),
        DBUser:    os.Getenv("DB_USER"),
        DBPass:    os.Getenv("DB_PASS"),
        DBName:    os.Getenv("DB_NAME"),
        SSLMode:   os.Getenv("DB_SSLMODE"),
        SecretKey: os.Getenv("SECRET_KEY"),
    }
}

func (c *Config) DSN() string {
    return fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        c.DBHost, c.DBUser, c.DBPass, c.DBName, c.DBPort, c.SSLMode,
    )
}
