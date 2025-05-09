package db

import (
	"context"
	"log"
	"notrust-server/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func ConnectDB() {
	dsn := "postgres://" +
		config.GetEnv("DB_USER", "notrust") + ":" +
		config.GetEnv("DB_PASSWORD", "supersecret") + "@" +
		config.GetEnv("DB_HOST", "localhost") + ":" +
		config.GetEnv("DB_PORT", "5432") + "/" +
		config.GetEnv("DB_NAME", "notrustdb")

	var err error
	Pool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}

	log.Println("Conectado a PostgreSQL")
}

