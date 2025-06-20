package config

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type Config interface {
	GetConnectionString() string
	Migration()
}

func (c *DatabaseConfig) GetConnectionString() string {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", "error", err)
		panic("Error loading .env file")
	}
	c.Host = os.Getenv("DB_HOST")
	if c.Host == "" {
		panic("DB_HOST is not set")
	}
	c.Port = os.Getenv("DB_PORT")
	if c.Port == "" {
		panic("DB_PORT is not set")
	}
	c.User = os.Getenv("DB_USER")
	if c.User == "" {
		panic("DB_USER is not set")
	}
	c.Password = os.Getenv("DB_PASSWORD")
	if c.Password == "" {
		panic("DB_PASSWORD is not set")
	}
	c.DBName = os.Getenv("DB_NAME")
	if c.DBName == "" {
		panic("DB_NAME is not set")
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.DBName)
}

func (c *DatabaseConfig) Migration(ctx context.Context) {
	db, err := sql.Open("postgres", c.GetConnectionString())
	if err != nil {
		panic(err)
	}
	dir, err := os.Getwd()
	if err != nil {
		slog.Error("Error getting working directory", "error", err)
		panic(err)
	}
	files, err := os.ReadDir(dir + "/models")
	if err != nil {
		slog.Error("Error reading models directory", "error", err)
		// panic(err)
	}
	for _, file := range files {
		content, err := os.ReadFile(dir + "/models/" + file.Name())
		if err != nil {
			slog.Error("Error reading model file", "error", err)
			panic(err)
		}
		_, err = db.ExecContext(ctx, string(content))
		if err != nil {
			slog.Error("Error executing migration", "error", err)
			panic(err)
		}
	}
	defer db.Close()
}
