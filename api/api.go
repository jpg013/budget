package api

import (

	// Importing go sql driver

	"fmt"
	"os"
	"time"

	"github.com/jpg013/budget"
	"github.com/jpg013/budget/config"
	"github.com/jpg013/budget/csv"
)

func makePostgresDSN(cfg config.Configuration) string {
	pass := os.Getenv("POSTGRES_PASSWORD")

	// Yea this is for dev only
	if pass == "" {
		pass = "dev"
	}

	return fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=disable host=%s",
		cfg.Postgres.Username,
		pass,
		cfg.Postgres.Database,
		cfg.Postgres.Port,
		cfg.Postgres.Host,
	)
}

func Start(cfg config.Configuration) error {
	db, err := budget.CreateDB(cfg)

	if err != nil {
		return err
	}

	err = csv.Migrate(db)

	if err != nil {
		return err
	}

	pipeline, err := budget.NewActivityPipeline(db)

	if err != nil {
		return err
	}

	err = pipeline.Start()

	if err != nil {
		return err
	}

	time.Sleep(1000 * time.Second)
	return nil
}
