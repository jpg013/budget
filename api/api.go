package api

import (

	// Importing go sql driver

	"fmt"
	"os"

	"github.com/jpg013/budget/config"
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
	fmt.Println("Hello there")

	return nil
	// db, err := budget.CreateDB(cfg)

	// if err != nil {
	// 	return err
	// }

	// processor, err := budget.NewActivityFileProcessor(db)

	// if err != nil {
	// 	return err
	// }

	// err = processor.Start()

	// if err != nil {
	// 	return err
	// }

	// time.Sleep(1000 * time.Second)
	// return nil
}
