package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Darlley/fc-gateway-golang/blob/develop/internal/repository"
	"github.com/Darlley/fc-gateway-golang/blob/develop/internal/service"
	"github.com/Darlley/fc-gateway-golang/blob/develop/internal/web/server"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "fullcycle-gateway"),
		getEnv("DB_SSLMODE", "disable"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)

	port := getEnv("HTTP_PORT", "8000")
	svr := server.NewServer(accountService, port)
	svr.ConfigureRoutes()
	svr.Start()

	if err := svr.Start(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
