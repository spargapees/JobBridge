package main

import (
	"database/sql"
	server "job_bridge/vacancies_service"
	"job_bridge/vacancies_service/internal/controllers"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	envPath = ".env"
)

func main() {
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file from %s: %v", envPath, err)
	}

	postgresURI := os.Getenv("DATABASE_URL")
	log.Print(postgresURI)

	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		log.Panic(err)
	}

	err = db.Ping()

	if err != nil {
		db.Close()
		log.Panic(err)
	}

	handler := controllers.NewController()

	srv := new(server.HTTPServer)
	if err := srv.Run_server("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("error with run server: %s", err.Error())
	}
}
