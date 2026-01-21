package main

import (
	"database/sql"
	"log"

	"github.com/iamdasun/go-gin-events-api/internal/database"
	"github.com/iamdasun/go-gin-events-api/internal/env"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models := database.NewModels(db)
	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "some-secret-1234567"),
		models:    models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
