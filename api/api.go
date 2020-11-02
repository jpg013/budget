package api

import (
	"strconv"

	"github.com/jpg013/budget"
	"github.com/jpg013/budget/models"

	// Importing go sql driver

	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jpg013/budget/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func ParseActivityFile(file string) []*models.Activity {
	// path := fmt.Sprintf("%s-%d-YearToDateSummary.csv", source, year)
	f, err := os.Open(file)

	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse file as CSV", err)
	}

	// Remove the first record (these are the column names)
	records = records[1:]
	rowNum := 1
	activityData := make([]*models.Activity, len(records))

	for _, r := range records {
		s := strings.Split(r[0], "/")
		s1 := fmt.Sprintf("%s-%s-%s", s[2], s[0], s[1])
		td, err := time.Parse("2006-01-02", s1)

		if err != nil {
			panic(err)
		}

		s = strings.Split(r[1], "/")
		s1 = fmt.Sprintf("%s-%s-%s", s[2], s[0], s[1])
		pd, err := time.Parse("2006-01-02", s1)

		if err != nil {
			panic(err)
		}

		amount, err := strconv.ParseFloat(r[3], 32)

		if err != nil {
			panic(err)
		}

		code := fmt.Sprintf("%s-%d", file)

		activityData[rowNum-1] = &models.Activity{
			TransactionDate: td,
			PostedDate:      pd,
			Description:     r[2],
			Amount:          float32(amount),
			Category:        r[4],
			Code:            code,
			SourceID:        1,
		}

		rowNum++
	}

	return activityData
}

func Start(cfg config.Configuration) error {
	dsn := makePostgresDSN(cfg)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	activityProcessor := budget.NewActivityProcessor(db)

	err = activityProcessor.Start()

	if err != nil {
		return err
	}

	time.Sleep(10 * time.Minute)
	// path := fmt.Sprintf("%s-%d-YearToDateSummary.csv", "Discover", 2020)
	// activities := ParseActivityFile("Discover", 2020)
	// for _, a := range activities {
	// 	result := db.Create(a)
	// 	fmt.Println(result.Error)
	// 	fmt.Println(result.RowsAffected)
	// 	fmt.Println(a.ID)
	// }
	// var activity models.Activity
	// db.First(&activity, "code = ?", "2020-01-01:1")
	// fmt.Println(activity.ActivityID)
	// activitySource := models.ActivitySource{Name: "Discover"}
	// result := db.Create(&activitySource) // pass pointer of data to Create
	// fmt.Println(result.Error)
	// fmt.Println(result.RowsAffected)
	// fmt.Println(activitySource.ID)
	// if err != nil {
	// 	panic(err)
	// }
	// logger := logging.NewLogger().WithTransports(
	// 	logging.NewStdOutTransport(logging.StdOutTransportConfig{Level: logging.InfoLevel}),
	// )
	// e := server.New(cfg, logger)
	// transport.NewHTTP(user.Initialize(db, cfg), e.Group("/user"))
	// server.Start(e, cfg)

	return nil
}
