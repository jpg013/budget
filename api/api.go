package api

import (

	// Importing go sql driver

	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jpg013/budget/config"
	"github.com/jpg013/budget/fileload"
	_ "github.com/lib/pq"
)

func makePostgresDSN(cfg config.Configuration) string {
	pass := os.Getenv("POSTGRES_PASSWORD")

	// Yea this is for dev only
	if pass == "" {
		pass = "dev"
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Username, pass, cfg.Postgres.Database)

	return psqlInfo
}

func Start(cfg config.Configuration) error {
	connStr := makePostgresDSN(cfg)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	fileload.NewManager(db)

	return nil
}
