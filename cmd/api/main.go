package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jpg013/budget/api"
	"github.com/jpg013/budget/config"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	cfg, err := config.NewParser().Parse("./cmd/api/config.json")

	checkErr(err)

	cfg.Application.Version = os.Getenv("S3_BUCKET")

	// Set application version
	checkErr(api.Start(cfg))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
